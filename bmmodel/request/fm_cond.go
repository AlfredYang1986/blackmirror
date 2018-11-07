package request

import (
	"gopkg.in/mgo.v2/bson"
)

type FmCond struct {
	Id   string      `json:"id"    bson:"_id"`
	Take int         `json:"take"  bson:"take"`
	Page int         `json:"page"  bson:"page"`
	Ky   string      `json:"key"`
	Vy   interface{} `json:"val"`
	Ct   string      `json:"category"`
}

func (t FmCond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t FmCond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond FmCond) Cond2QueryObj(cate string) bson.M {
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp {
		return bson.M{cond.Ky: cond.Vy}
	} else {
		return bson.M{}
	}
}

func (cond FmCond) Cond2UpdateObj() bson.M {
	return bson.M{}
}

func (cond FmCond) IsQueryCondi() bool {
	return true
}

func (cond FmCond) IsUpdateCondi() bool {
	return false
}
