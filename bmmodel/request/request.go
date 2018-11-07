package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Request struct {
	Id     string      `json:"id"`
	Res    string      `json:"res"`
	Cond   []Condition `json:"conditions" jsonapi:"relationships"`
	Eqcond []Eqcond    `json:"Eqcond" jsonapi:"relationships"`
	Fmcond Fmcond    `json:"Fmcond" jsonapi:"relationships"`
	Upcond []Upcond    `json:"Upcond" jsonapi:"relationships"`
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
	case "Eqcond":
		var rst []Eqcond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Eqcond))
		}
		req.Eqcond = rst
	case "Fmcond":
		req.Fmcond = v.(Fmcond)
	case "Upcond":
		var rst []Upcond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Upcond))
		}
		req.Upcond = rst
	}
	return req
}

func (req Request) QueryConnect(tag string) interface{} {
	switch tag {
	case "conditions":
		return req.Cond
	case "Eqcond":
		return req.Eqcond
	case "Fmcond":
		return req.Fmcond
	case "Upcond":
		return req.Upcond
	}
	return req

}

func (req Request) Cond2QueryObj(cat string) bson.M {
	rst := make(map[string]interface{})

	var conds_tmp []Eqcond
	var conds_all []Eqcond
	for _, cond := range req.Cond {
		conds_tmp = append(conds_tmp, cond.(Eqcond))
	}
	conds_all = append(conds_tmp, req.Eqcond...)

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

func (req Request) Cond2UpdateObj(cat string) bson.M {
	rst := make(map[string]interface{})

	var conds_tmp []Upcond
	var conds_all []Upcond
	for _, cond := range req.Cond {
		conds_tmp = append(conds_tmp, cond.(Upcond))
	}
	conds_all = append(conds_tmp, req.Upcond...)

	for _, cond := range conds_all {
		if cond.IsUpdateCondi() {
			for k, v := range cond.Cond2UpdateObj(cat) {
				rst[k] = v
			}
		}
	}
	return rst
}

func (req Request) CondiQueryVal(ky string, cat string) interface{} {
	for _, cond := range req.Eqcond {
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
