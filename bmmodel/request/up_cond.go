package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Upcond struct {
	Id string      `json:"id"`
	Ky string      `json:"key"`
	Vy interface{} `json:"val"`
	Ct string      `json:"category"`
}

func (t Upcond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Upcond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Upcond) Cond2QueryObj(cat string) bson.M {
	return bson.M{}
}

func (cond Upcond) Cond2UpdateObj(cate string) bson.M {
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

func (cond Upcond) IsQueryCondi() bool {
	return false
}

func (cond Upcond) IsUpdateCondi() bool {
	return true
}
