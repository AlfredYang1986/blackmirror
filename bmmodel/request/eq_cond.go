package request

import (
	"gopkg.in/mgo.v2/bson"
)

type EQCond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
}

func (t EQCond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t EQCond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond EQCond) Cond2QueryObj() bson.M {
	return bson.M{cond.Ky: cond.Vy}
}
