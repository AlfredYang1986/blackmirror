package jsonapiobj

import (
	"errors"
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"reflect"
	"strings"
)

type JsObject struct {
	tag string
	obj map[string]interface{}
}

func FromObject(ptr interface{}) (interface{}, error) {
	v := reflect.ValueOf(ptr).Elem()
	return struct2jsonAcc(v)
}

func struct2jsonAcc(v reflect.Value) (interface{}, error) {
	rst := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a.reflect.struct.field
		fieldValue := v.Field(i)
		tag := fieldInfo.Tag // a.reflect.tag

		var name string
		if tag.Get(bmmodel.BMJson) != "" {
			name = tag.Get(bmmodel.BMJson)
		} else {
			name = strings.ToLower(fieldInfo.Name)
		}

		if reval, err := value2jsonAcc(fieldValue); err == nil {
			rst[name] = reval
		}
	}

	return rst, nil
}

func value2jsonAcc(v reflect.Value) (interface{}, error) {

	switch v.Kind() {
	default:
		return nil, errors.New("not implement")
	case reflect.Invalid:
		return nil, errors.New("invalid")
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), nil
	case reflect.String:
		return v.String(), nil
	case reflect.Array, reflect.Slice:
		var rst []interface{}
		for i := 0; i < v.Len(); i++ {
			tmp, _ := value2jsonAcc(v.Index(i))
			rst = append(rst, tmp)
		}
		return rst, nil
	case reflect.Map:
		rst := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			kv := v.MapIndex(key)
			tmp, err := value2jsonAcc(kv)
			if err != nil {
				panic(err)
			}
			rst[key.String()] = tmp
		}
		return rst, nil
	case reflect.Struct, reflect.Interface:
		return struct2jsonAcc(v)
	}
}
