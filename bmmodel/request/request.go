package request

import (
	"gopkg.in/mgo.v2/bson"
)

type Request struct {
	Id       string      `json:"id"`
	Res      string      `json:"res"`
	Cond     []Condition `json:"conditions" jsonapi:"relationships"`
	Eqcond   []Eqcond    `json:"Eqcond" jsonapi:"relationships"`
	Necond   []Necond    `json:"Necond" jsonapi:"relationships"`
	Gtcond   []Gtcond    `json:"Gtcond" jsonapi:"relationships"`
	Gtecond  []Gtecond   `json:"Gtecond" jsonapi:"relationships"`
	Ltcond   []Ltcond    `json:"Ltcond" jsonapi:"relationships"`
	Ltecond  []Ltecond   `json:"Ltecond" jsonapi:"relationships"`
	Incond   []Incond    `json:"Incond" jsonapi:"relationships"`
	Nincond  []Nincond   `json:"Nincond" jsonapi:"relationships"`
	Fmcond   Fmcond      `json:"Fmcond" jsonapi:"relationships"`
	Upcond   []Upcond    `json:"Upcond" jsonapi:"relationships"`
	Sortcond []Sortcond  `json:"Sortcond" jsonapi:"relationships"`
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
	case "Necond":
		var rst []Necond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Necond))
		}
		req.Necond = rst
	case "Gtcond":
		var rst []Gtcond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Gtcond))
		}
		req.Gtcond = rst
	case "Gtecond":
		var rst []Gtecond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Gtecond))
		}
		req.Gtecond = rst
	case "Ltcond":
		var rst []Ltcond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Ltcond))
		}
		req.Ltcond = rst
	case "Ltecond":
		var rst []Ltecond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Ltecond))
		}
		req.Ltecond = rst
	case "Incond":
		var rst []Incond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Incond))
		}
		req.Incond = rst
	case "Nincond":
		var rst []Nincond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Nincond))
		}
		req.Nincond = rst
	case "Fmcond":
		req.Fmcond = v.(Fmcond)
	case "Upcond":
		var rst []Upcond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Upcond))
		}
		req.Upcond = rst
	case "Sortcond":
		var rst []Sortcond
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(Sortcond))
		}
		req.Sortcond = rst
	}
	return req
}

func (req Request) QueryConnect(tag string) interface{} {
	switch tag {
	case "conditions":
		return req.Cond
	case "Eqcond":
		return req.Eqcond
	case "Necond":
		return req.Necond
	case "Gtcond":
		return req.Gtcond
	case "Gtecond":
		return req.Gtecond
	case "Ltcond":
		return req.Ltcond
	case "Ltecond":
		return req.Ltecond
	case "Incond":
		return req.Incond
	case "Nincond":
		return req.Nincond
	case "Fmcond":
		return req.Fmcond
	case "Upcond":
		return req.Upcond
	case "Sortcond":
		return req.Sortcond
	}
	return req

}

func (req Request) Cond2QueryObj(cat string) bson.M {
	rst := make(map[string]interface{})

	for _, cond := range req.Gtcond {
		if cond.IsQueryCondi() {
			tmp := cond.Cond2QueryObj(cat)
			for k, v := range tmp {
				rst[k] = v
			}
		}
	}

	for _, cond := range req.Gtecond {
		if cond.IsQueryCondi() {
			tmp := cond.Cond2QueryObj(cat)
			for k, v := range tmp {
				rst[k] = v
			}
		}
	}

	for _, cond := range req.Ltcond {
		if cond.IsQueryCondi() {
			tmp := cond.Cond2QueryObj(cat)
			for k, v := range tmp {
				rst[k] = v
			}
		}
	}

	for _, cond := range req.Ltecond {
		if cond.IsQueryCondi() {
			tmp := cond.Cond2QueryObj(cat)
			for k, v := range tmp {
				rst[k] = v
			}
		}
	}

	for _, cond := range req.Cond {
		if cond.IsQueryCondi() {
			tmp := cond.Cond2QueryObj(cat)
			for k, v := range tmp {
				rst[k] = v
			}
		}
	}

	for _, cond := range req.Incond {
		if cond.IsQueryCondi() {
			tmp := cond.Cond2QueryObj(cat)
			for k, v := range tmp {
				rst[k] = v
			}
		}
	}

	for _, cond := range req.Nincond {
		if cond.IsQueryCondi() {
			tmp := cond.Cond2QueryObj(cat)
			for k, v := range tmp {
				rst[k] = v
			}
		}
	}

	for _, cond := range req.Eqcond {
		if cond.IsQueryCondi() {
			tmp := cond.Cond2QueryObj(cat)
			for k, v := range tmp {
				rst[k] = v
			}
		}
	}

	for _, cond := range req.Necond {
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

func (req Request) Cond2SortObj(cate string) []string {
	var result []string
	for _, cond := range req.Sortcond {
		tmp := cond.Cond2SortObj(cate)
		if tmp != "" {
			result = append(result, tmp)
		}
	}
	//如果Sortcond有参数，优先对其进行排序，最后再按create_time倒序排序
	//记得手动对需要排序的实体创建CreateTime
	result = append(result, "-create_time")
	return result
}
