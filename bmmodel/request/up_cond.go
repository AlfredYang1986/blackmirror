package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Upcond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
}

func (t Upcond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Upcond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Upcond) Cond2QueryObj(cat string) bson.M {
	return bson.M{}
}

func (cond Upcond) Cond2UpdateObj() bson.M {
	return bson.M{cond.Ky: cond.Vy}
}

func (cond Upcond) IsQueryCondi() bool {
	return false
}

func (cond Upcond) IsUpdateCondi() bool {
	return true
}
