package bmmodel

import (
	"errors"
	"reflect"
	"strings"
)

type NoPtr struct {
}

const (
	BMJson  string = "json"
	BMMongo string = "mongo"
)

func AttrWithName(ptr interface{}, attr string, tagN string) (interface{}, error) {
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a.reflect.struct.field
		fieldValue := v.Field(i)
		tag := fieldInfo.Tag // a.reflect.tag

		var name string
		if tagN == BMJson {
			name = tag.Get(BMJson)
		} else if tagN == BMMongo {
			name = tag.Get(BMMongo)
		} else {
			name = strings.ToLower(fieldInfo.Name)
		}

		if name == attr {
			return attrValue(fieldValue)
		}
	}

	return NoPtr{}, nil
}

func attrValue(v reflect.Value) (interface{}, error) {
	switch v.Kind() {
	case reflect.Invalid:
		return nil, nil
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
			tmp, _ := attrValue(v.Index(i))
			rst = append(rst, tmp)
		}
		return rst, nil
	}

	return NoPtr{}, errors.New("not implement")
}
