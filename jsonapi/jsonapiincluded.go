package jsonapi

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/relationships"
)

func remapIncluded(inc []interface{}) ([]interface{}, error) {

	var result []interface{}
	//increl := 0
	for _, item := range inc {
		itm := item.(map[string]interface{})
		rst := itm[itm[TYPE].(string)].(relationships.Relationships)
		for k, _ := range itm {
			if k == ATTRIBUTES {
				panic("attrtibutes not parse corectly")
			} else if k == RELATIONSHIPS {
				//increl++
				rs := itm[RELATIONSHIPS].(map[string]interface{})
				reval, _ := queryRelationships(rs, inc, rst)
				itm[itm[TYPE].(string)] = reval
			}
		}
		result = append(result, itm)
	}

	return result, nil
}
