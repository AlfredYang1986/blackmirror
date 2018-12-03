package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Incond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
	Ct string      `json:"category"`
}

func (t Incond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Incond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Incond) Cond2QueryObj(cate string) bson.M {
	//TODO:当传递eq_cond并包含category时,以下逻辑会有一些问题.
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		v := make(map[string]interface{})
		if cond.Ky == "id" {
			ids := []bson.ObjectId{}
			for _,id := range cond.Vy.([]string) {
				ids = append(ids, bson.ObjectIdHex(id))
			}
			v["$in"] = ids
			return bson.M{"_id": v}
		}
		v["$in"] = cond.Vy
		return bson.M{cond.Ky: v}
	} else {
		return bson.M{}
	}
}

func (cond Incond) Cond2UpdateObj(cate string) bson.M {
	return bson.M{}
}

func (cond Incond) IsQueryCondi() bool {
	return true
}

func (cond Incond) IsUpdateCondi() bool {
	return false
}
