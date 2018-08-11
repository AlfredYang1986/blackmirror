package request

import (
	"gopkg.in/mgo.v2/bson"
)

type EQCond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
	Ct string      `json:"category"`
}

func (t EQCond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t EQCond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond EQCond) Cond2QueryObj(cate string) bson.M {
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		return bson.M{cond.Ky: cond.Vy}
	} else {
		return bson.M{}
	}
}

func (cond EQCond) Cond2UpdateObj() bson.M {
	return bson.M{}
}

func (cond EQCond) IsQueryCondi() bool {
	return true
}

func (cond EQCond) IsUpdateCondi() bool {
	return false
}
