package bmmate

import (
	"reflect"
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
