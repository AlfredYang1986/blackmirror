package bmmate

import (
	"fmt"
	"reflect"
	"strings"
)

func IsSeq(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Array ||
		reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsStruct(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}

func IsMap(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func PushRelationships2Model(ptr interface{}, tag string, res interface{}) {
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a.reflect.struct.fieldv
		fieldName := strings.ToLower(fieldInfo.Name)
		fmt.Println(fieldName)

		if fieldName == "Relationships" {
			fmt.Printf("haaha")
		}

		//fieldValue := v.Field(i)
		//tag := fieldInfo.Tag // a.reflect.tag

		//if name == attr {
		//return attrValue(fieldValue)
		//}
	}
}
