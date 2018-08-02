package test

import (
//"reflect"
)

type Test struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func (t Test) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Test) QueryConnect(tag string) interface{} {
	return nil
}
