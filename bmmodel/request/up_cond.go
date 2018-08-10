package request

import (
	"gopkg.in/mgo.v2/bson"
)

type UPCond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
}

func (t UPCond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t UPCond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond UPCond) Cond2QueryObj() bson.M {
	return bson.M{}
}

func (cond UPCond) Cond2UpdateObj() bson.M {
	return bson.M{cond.Ky: cond.Vy}
}

func (cond UPCond) IsQueryCondi() bool {
	return false
}

func (cond UPCond) IsUpdateCondi() bool {
	return true
}
