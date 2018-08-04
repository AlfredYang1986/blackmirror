package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Request struct {
	Id   string      `json:"id"`
	Res  string      `json:"res"`
	Cond []Condition `json:"conditions" jsonapi:"relationships"`
}

func (req Request) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	default:
		panic("not implement")
	case "conditions":
		var rst []Condition
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Condition))
		}
		req.Cond = rst
	}
	return req
}

func (req Request) QueryConnect(tag string) interface{} {
	switch tag {
	case "conditions":
		return req.Cond
	}
	return req

}

func (req Request) Cond2QueryObj() bson.M {
	rst := make(map[string]interface{})
	for _, cond := range req.Cond {
		for k, v := range cond.Cond2QueryObj() {
			rst[k] = v
		}
	}
	return rst
}
