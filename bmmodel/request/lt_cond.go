package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Ltcond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
	Ct string      `json:"category"`
}

func (t Ltcond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Ltcond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Ltcond) Cond2QueryObj(cate string) bson.M {
	//TODO:当传递eq_cond并包含category时,以下逻辑会有一些问题.
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		v := make(map[string]interface{})
		v["$lt"] = cond.Vy
		return bson.M{cond.Ky: v}
	} else {
		return bson.M{}
	}
}

func (cond Ltcond) Cond2UpdateObj(cate string) bson.M {
	return bson.M{}
}

func (cond Ltcond) IsQueryCondi() bool {
	return true
}

func (cond Ltcond) IsUpdateCondi() bool {
	return false
}
