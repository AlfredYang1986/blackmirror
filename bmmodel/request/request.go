package request

type Request struct {
	Id   string      `json:"id"`
	res  string      `json:"res"`
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
