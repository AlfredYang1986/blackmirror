package request

import (
	"gopkg.in/mgo.v2/bson"
)

type UpCond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
}

func (t UpCond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t UpCond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond UpCond) Cond2QueryObj(cat string) bson.M {
	return bson.M{}
}

func (cond UpCond) Cond2UpdateObj() bson.M {
	return bson.M{cond.Ky: cond.Vy}
}

func (cond UpCond) IsQueryCondi() bool {
	return false
}

func (cond UpCond) IsUpdateCondi() bool {
	return true
}
