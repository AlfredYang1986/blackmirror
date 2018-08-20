package request

import (
	"gopkg.in/mgo.v2/bson"
)

type FMUCond struct {
	Id   string `json:"id"    bson:"_id"`
	Take int    `json:"take"  bson:"take"`
	Page int    `json:"page"  bson:"page"`
}

func (t FMUCond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t FMUCond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond FMUCond) Cond2QueryObj(cate string) bson.M {
	return bson.M{}
}

func (cond FMUCond) Cond2UpdateObj() bson.M {
	return bson.M{}
}

func (cond FMUCond) IsQueryCondi() bool {
	return true
}

func (cond FMUCond) IsUpdateCondi() bool {
	return false
}
