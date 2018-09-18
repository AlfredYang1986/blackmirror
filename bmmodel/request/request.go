package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Request struct {
	Id     string      `json:"id"`
	Res    string      `json:"res"`
	Cond   []Condition `json:"conditions" jsonapi:"relationships"`
	EqCond []EQCond    `json:"eqcond" jsonapi:"relationships"`
	FmCond []FMUCond   `json:"fmcond" jsonapi:"relationships"`
	UpCond []UPCond    `json:"upcond" jsonapi:"relationships"`
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
	case "eqcond":
		var rst []EQCond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(EQCond))
		}
		req.EqCond = rst
	case "fmcond":
		var rst []FMUCond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(FMUCond))
		}
		req.FmCond = rst
	case "upcond":
		var rst []UPCond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(UPCond))
		}
		req.UpCond = rst
	}
	return req
}

func (req Request) QueryConnect(tag string) interface{} {
	switch tag {
	case "conditions":
		return req.Cond
	case "eqcond":
		return req.EqCond
	case "fmcond":
		return req.FmCond
	case "upcond":
		return req.UpCond
	}
	return req

}

func (req Request) Cond2QueryObj(cat string) bson.M {
	rst := make(map[string]interface{})

	var conds_tmp []EQCond
	var conds_all []EQCond
	for _, cond := range req.Cond {
		conds_tmp = append(conds_tmp, cond.(EQCond))
	}
	conds_all = append(conds_tmp, req.EqCond...)

	for _, cond := range conds_all {
		if cond.IsQueryCondi() {
			for k, v := range cond.Cond2QueryObj(cat) {
				rst[k] = v
			}
		}
	}
	return rst
}

func (req Request) Cond2UpdateObj() bson.M {
	rst := make(map[string]interface{})

	var conds_tmp []UPCond
	var conds_all []UPCond
	for _, cond := range req.Cond {
		conds_tmp = append(conds_tmp, cond.(UPCond))
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
