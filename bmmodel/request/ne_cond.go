package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Necond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
	Ct string      `json:"category"`
}

func (t Necond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Necond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Necond) Cond2QueryObj(cate string) bson.M {
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		v := make(map[string]interface{})
		if cond.Ky == "id" {
			v["$ne"] = bson.ObjectIdHex(cond.Vy.(string))
			return bson.M{"_id": v}
		}
		v["$ne"] = cond.Vy
		return bson.M{cond.Ky: v}
	} else {
		return bson.M{}
	}
}

func (cond Necond) Cond2UpdateObj(cate string) bson.M {
	return bson.M{}
}

func (cond Necond) IsQueryCondi() bool {
	return true
}

func (cond Necond) IsUpdateCondi() bool {
	return false
}
