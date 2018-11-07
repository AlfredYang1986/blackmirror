package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Fmcond struct {
	Id   string      `json:"id"    bson:"_id"`
	Take int         `json:"take"  bson:"take"`
	Page int         `json:"page"  bson:"page"`
	Ky   string      `json:"key"`
	Vy   interface{} `json:"val"`
	Ct   string      `json:"category"`
}

func (t Fmcond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Fmcond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Fmcond) Cond2QueryObj(cate string) bson.M {
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp {
		return bson.M{cond.Ky: cond.Vy}
	} else {
		return bson.M{}
	}
}

func (cond Fmcond) Cond2UpdateObj() bson.M {
	return bson.M{}
}

func (cond Fmcond) IsQueryCondi() bool {
	return true
}

func (cond Fmcond) IsUpdateCondi() bool {
	return false
}
