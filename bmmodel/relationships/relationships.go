package relationships

import (
//"reflect"
)

type Relationships interface {
	//SetConnect(tag string, v interface{}, tp reflect.Type) interface{}
	SetConnect(tag string, v interface{}) interface{}
	QueryConnect(tag string) interface{}
}
