package test

import (
	//"reflect"
	"gopkg.in/mgo.v2/bson"
)

type Test struct {
	Id    string        `json:"id"`
	Id_   bson.ObjectId `bson:"_id"`
	Title string        `json:"title" bson:"title"`
}

func (t Test) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Test) QueryConnect(tag string) interface{} {
	return nil
}
