package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Sortcond struct {
	Id string `json:"id"`
	Ky string `json:"key"`
	Vy int    `json:"val"`
	Ct string `json:"category"`
}

func (t Sortcond) SetConnect(tag string, v interface{}) interface{} {
	return t
}

func (t Sortcond) QueryConnect(tag string) interface{} {
	return nil
}

func (cond Sortcond) Cond2QueryObj(cate string) bson.M {
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		return bson.M{cond.Ky: cond.Vy}
	} else {
		return bson.M{}
	}
}

func (cond Sortcond) Cond2UpdateObj(cate string) bson.M {
	return bson.M{}
}

func (cond Sortcond) IsQueryCondi() bool {
	return true
}

func (cond Sortcond) IsUpdateCondi() bool {
	return false
}

func (cond Sortcond) Cond2SortObj(cate string) string {
	tmp := len(cond.Ct) > 0 && cond.Ct == cate
	if tmp || len(cond.Ct) == 0 {
		switch cond.Vy {
		case 1:
			return cond.Ky
		case -1:
			return "-" + cond.Ky
		default:
			return ""
		}
	} else {
		return ""
	}
}
