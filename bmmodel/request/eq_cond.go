package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Eqcond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
	Ct string      `json:"category"`
}

func (t Eqcond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Eqcond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Eqcond) Cond2QueryObj(cate string) bson.M {
	//TODO:当传递eq_cond并包含category时,以下逻辑会有一些问题.
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		if cond.Ky == "id" {
			v := bson.ObjectIdHex(cond.Vy.(string))
			return bson.M{"_id": v}
		}
		return bson.M{cond.Ky: cond.Vy}
	} else {
		return bson.M{}
	}
}

func (cond Eqcond) Cond2UpdateObj() bson.M {
	return bson.M{}
}

func (cond Eqcond) IsQueryCondi() bool {
	return true
}

func (cond Eqcond) IsUpdateCondi() bool {
	return false
}
