package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Ltecond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
	Ct string      `json:"category"`
}

func (t Ltecond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Ltecond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Ltecond) Cond2QueryObj(cate string) bson.M {
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		v := make(map[string]interface{})
		v["$lte"] = cond.Vy
		return bson.M{cond.Ky: v}
	} else {
		return bson.M{}
	}
}

func (cond Ltecond) Cond2UpdateObj(cate string) bson.M {
	return bson.M{}
}

func (cond Ltecond) IsQueryCondi() bool {
	return true
}

func (cond Ltecond) IsUpdateCondi() bool {
	return false
}
