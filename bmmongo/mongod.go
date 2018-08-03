package bmmongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"log"
	"errors"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"reflect"
	"strings"
)

func InsertBMObject(ptr interface{}) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	v := reflect.ValueOf(ptr).Elem()

	rst := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a.reflect.struct.field
		fieldValue := v.Field(i)
		tag := fieldInfo.Tag // a.reflect.tag

		var name string
		if tag.Get(bmmodel.BMMongo) != "" {
			name = tag.Get(bmmodel.BMMongo)
		} else {
			name = strings.ToLower(fieldInfo.Name)
		}

		if name == "_id" {
			oid := bson.NewObjectId()
			rst[name] = oid
			continue
		} else if name == "found" {
			continue
		} else if name == "locations" {
			continue
		} else {
			tmp, _ := AttrValue(fieldValue)
			rst[name] = tmp
		}

	}
	fmt.Println(rst)

	cn := v.Type().Name()
	c := session.DB("test").C(cn)
	c.Insert(rst)

	return err
}

func AttrValue(v reflect.Value) (interface{}, error) {
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
			tmp, _ := AttrValue(v.Index(i))
			rst = append(rst, tmp)
		}
		return rst, nil
	case reflect.Map:
		rst := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			kv := v.MapIndex(key)
			tmp, _ := AttrValue(kv)
			rst[key.String()] = tmp
		}
		return rst, nil
	}

	return 0, errors.New("not implement")
}
