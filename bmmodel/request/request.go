package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Request struct {
	Id     string      `json:"id"`
	Res    string      `json:"res"`
	Cond   []Condition `json:"conditions" jsonapi:"relationships"`
	EqCond []EqCond    `json:"EqCond" jsonapi:"relationships"`
	FmCond []FmCond    `json:"FmCond" jsonapi:"relationships"`
	UpCond []UpCond    `json:"UpCond" jsonapi:"relationships"`
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
	case "EqCond":
		var rst []EqCond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(EqCond))
		}
		req.EqCond = rst
	case "FmCond":
		var rst []FmCond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(FmCond))
		}
		req.FmCond = rst
	case "UpCond":
		var rst []UpCond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(UpCond))
		}
		req.UpCond = rst
	}
	return req
}

func (req Request) QueryConnect(tag string) interface{} {
	switch tag {
	case "conditions":
		return req.Cond
	case "EqCond":
		return req.EqCond
	case "FmCond":
		return req.FmCond
	case "UpCond":
		return req.UpCond
	}
	return req

}

func (req Request) Cond2QueryObj(cat string) bson.M {
	rst := make(map[string]interface{})

	var conds_tmp []EqCond
	var conds_all []EqCond
	for _, cond := range req.Cond {
		conds_tmp = append(conds_tmp, cond.(EqCond))
	}
	conds_all = append(conds_tmp, req.EqCond...)

	for _, cond := range conds_all {
		if cond.IsQueryCondi() {
			tmp := cond.Cond2QueryObj(cat)
			for k, v := range tmp {
				rst[k] = v
			}
		}
	}
	return rst
}

func (req Request) Cond2UpdateObj() bson.M {
	rst := make(map[string]interface{})

	var conds_tmp []UpCond
	var conds_all []UpCond
	for _, cond := range req.Cond {
		conds_tmp = append(conds_tmp, cond.(UpCond))
	}
	conds_all = append(conds_tmp, req.UpCond...)

	for _, cond := range conds_all {
		if cond.IsUpdateCondi() {
			for k, v := range cond.Cond2UpdateObj() {
				rst[k] = v
			}
		}
	}
	return rst
}

func (req Request) CondiQueryVal(ky string, cat string) interface{} {
	for _, cond := range req.EqCond {
		if cond.IsQueryCondi() {
			for k, v := range cond.Cond2QueryObj(cat) {
				if k == ky {
					return v
				}
			}
		}
	}
	return nil
}
