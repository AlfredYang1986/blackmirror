package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Nincond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
	Ct string      `json:"category"`
}

func (t Nincond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Nincond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Nincond) Cond2QueryObj(cate string) bson.M {
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		v := make(map[string]interface{})
		if cond.Ky == "id" {
			ids := []bson.ObjectId{}
			for _,id := range cond.Vy.([]string) {
				ids = append(ids, bson.ObjectIdHex(id))
			}
			v["$nin"] = ids
			return bson.M{"_id": v}
		}
		v["$nin"] = cond.Vy
		return bson.M{cond.Ky: v}
	} else {
		return bson.M{}
	}
}

func (cond Nincond) Cond2UpdateObj(cate string) bson.M {
	return bson.M{}
}

func (cond Nincond) IsQueryCondi() bool {
	return true
}

func (cond Nincond) IsUpdateCondi() bool {
	return false
}
