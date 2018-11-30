package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Gtcond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
	Ct string      `json:"category"`
}

func (t Gtcond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Gtcond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Gtcond) Cond2QueryObj(cate string) bson.M {
	//TODO:当传递eq_cond并包含category时,以下逻辑会有一些问题.
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		v := make(map[string]interface{})
		v["$gt"] = cond.Vy
		return bson.M{cond.Ky: v}
	} else {
		return bson.M{}
	}
}

func (cond Gtcond) Cond2UpdateObj(cate string) bson.M {
	return bson.M{}
}

func (cond Gtcond) IsQueryCondi() bool {
	return true
}

func (cond Gtcond) IsUpdateCondi() bool {
	return false
}
