package jsonapi

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmate"
	//"github.com/alfredyang1986/blackmirror/bmmodel/brand"
	"github.com/alfredyang1986/blackmirror/bmmodel/relationships"
)

func map2Object(m map[string]interface{}) (interface{}, error) {

	rt := m[ROOT].(map[string]interface{})
	tdt := rt[DATA]
	inc := rt[INCLUDED].([]interface{})

	remapIncluded(inc)

	if bmmate.IsSeq(tdt) {
		var result []interface{}
		ldt := tdt.([]interface{})
		for _, ndt := range ldt {
			dt := ndt.(map[string]interface{})
			rst := dt[dt[TYPE].(string)].(relationships.Relationships)
			rs := dt[RELATIONSHIPS].(map[string]interface{})
			reval, _ := queryRelationships(rs, inc, rst)
			fmt.Println(reval)
			result = append(result, reval)
		}
		return result, nil
	} else if bmmate.IsMap(tdt) {
		dt := tdt.(map[string]interface{})
		rst := dt[dt[TYPE].(string)].(relationships.Relationships)
		rs := dt[RELATIONSHIPS].(map[string]interface{})
		reval, _ := queryRelationships(rs, inc, rst)
		fmt.Println(reval)
		return reval, nil
	}

	return m, errors.New("something wrong")
}

func queryRelationships(rs map[string]interface{}, inc []interface{}, m relationships.Relationships) (interface{}, error) {
	var rst interface{}

	for k, v := range rs {
		tmp := v.(map[string]interface{})
		vd := tmp[DATA]
		if bmmate.IsMap(vd) {
			vdm := vd.(map[string]interface{})
			vid := vdm[ID].(string)
			vtype := vdm[TYPE].(string)
			incv, _ := qRIObj(vid, vtype, inc)
			//fmt.Println(incv)
			rst = m.SetConnect(k, incv)
		} else if bmmate.IsSeq(vd) {
			vdl := vd.([]interface{})
			var ritem []interface{}
			for _, item := range vdl {
				vdm := item.(map[string]interface{})
				vid := vdm[ID].(string)
				vtype := vdm[TYPE].(string)
				incv, _ := qRIObj(vid, vtype, inc)
				//fmt.Println(incv)
				ritem = append(ritem, incv)
			}
			rst = m.SetConnect(k, ritem)
		}
	}
	return rst, nil
}

func qRIObj(vid string, vtype string, inc []interface{}) (interface{}, error) {
	for _, item := range inc {
		incm := item.(map[string]interface{})
		//fmt.Println(incm)
		incmid := incm[ID].(string)
		incmtype := incm[TYPE].(string)
		if vid == incmid && vtype == incmtype {
			return incm[incmtype], nil
		}
	}

	return 0, errors.New("not included")
}
