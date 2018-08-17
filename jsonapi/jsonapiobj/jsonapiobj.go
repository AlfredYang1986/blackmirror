package jsonapiobj

import (
	"errors"
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmmate"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"reflect"
	"strings"
)

type JsResult struct {
	Obj map[string]interface{}
	Inc []map[string]interface{}
}

func (o *JsResult) FromObject(ptr interface{}) error {

	if bmmate.IsSeq(ptr) {
		ppr := ptr.([]interface{})
		var rst []interface{}
		for _, itm := range ppr {
			v := reflect.ValueOf(itm).Elem()
			tmp, _ := o.struct2jsonAcc(v)
			rst = append(rst, tmp)
		}
		o.Obj = map[string]interface{}{"data": rst, "included": o.Inc}
	} else {
		v := reflect.ValueOf(ptr).Elem()
		tmp, _ := o.struct2jsonAcc(v)
		o.Obj = map[string]interface{}{"data": tmp, "included": o.Inc}
	}

	return nil
}

func (o *JsResult) struct2jsonAcc(v reflect.Value) (interface{}, error) {
	var rsl []string
	var atr []string
	attr := make(map[string]interface{})
	rships := make(map[string]interface{})
	result := make(map[string]interface{})

	rst := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a.reflect.struct.field
		fieldValue := v.Field(i)
		tag := fieldInfo.Tag // a.reflect.tag

		var name string
		if tag.Get(bmmodel.BMJson) != "" {
			name = tag.Get(bmmodel.BMJson)
		} else {
			name = strings.ToLower(fieldInfo.Name)
		}

		if tag.Get(bmmodel.BMJsonAPI) == "relationships" {
			rsl = append(rsl, name)
		} else {
			atr = append(atr, name)
		}

		if reval, err := o.value2jsonAcc(fieldValue); err == nil {
			rst[name] = reval
		}
	}

	for _, ky := range atr {
		if ky != "id" && ky != "id_" {
			attr[ky] = rst[ky]
		}
	}

	for _, ky := range rsl {
		tmp := make(map[string]interface{})
		val := rst[ky]
		if bmmate.IsMap(val) {
			tmp["data"] = val
		} else if bmmate.IsSeq(val) {
			var rt []interface{}
			for _, tt := range val.([]interface{}) {
				rt = append(rt, tt)
			}
			tmp["data"] = rt
		}
		rships[ky] = tmp
	}

	result["id"] = rst["id"]
	result["type"] = v.Type().Name()
	result["attributes"] = attr

	if len(rships) > 0 {
		rships, _ = o.remapRS2Included(rships)
		result["relationships"] = rships
	}

	return result, nil
}

func (o *JsResult) value2jsonAcc(v reflect.Value) (interface{}, error) {

	switch v.Kind() {
	default:
		return nil, errors.New("not implement")
	case reflect.Invalid:
		return nil, errors.New("invalid")
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	case reflect.String:
		return v.String(), nil
	case reflect.Array, reflect.Slice:
		var rst []interface{}
		for i := 0; i < v.Len(); i++ {
			tmp, _ := o.value2jsonAcc(v.Index(i))
			rst = append(rst, tmp)
		}
		return rst, nil
	case reflect.Map:
		rst := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			kv := v.MapIndex(key)
			tmp, err := o.value2jsonAcc(kv)
			if err != nil {
				panic(err)
			}
			rst[key.String()] = tmp
		}
		return rst, nil
	case reflect.Struct, reflect.Interface:
		return o.struct2jsonAcc(v)
	}
}

func (o *JsResult) remapRS2Included(rships map[string]interface{}) (map[string]interface{}, error) {
	rst := make(map[string]interface{})
	for k, v := range rships {
		vm := v.(map[string]interface{})
		vmdat := vm["data"] //.(map[string]interface{})

		if bmmate.IsSeq(vmdat) {
			vml := vmdat.([]interface{})
			var rev []map[string]interface{}
			for _, itm := range vml {
				tmp, _ := o.mapRS2IncludedAcc(itm.(map[string]interface{}))
				rev = append(rev, tmp)
			}
			dt := make(map[string]interface{})
			dt["data"] = rev
			rst[k] = dt
			//rst[k] = rev

		} else if bmmate.IsMap(vmdat) {
			tmp, _ := o.mapRS2IncludedAcc(vmdat.(map[string]interface{}))
			dt := make(map[string]interface{})
			dt["data"] = tmp
			rst[k] = dt
		}

	}

	return rst, nil
}

func (o *JsResult) mapRS2IncludedAcc(vmm map[string]interface{}) (map[string]interface{}, error) {
	vid := vmm["id"].(string)
	vtype := vmm["type"].(string)

	rst := make(map[string]interface{})

	bExist := false
	for _, itm := range o.Inc {
		incmid := itm["id"]
		incmtp := itm["type"]

		if vid == incmid && vtype == incmtp {
			bExist = true
		}
	}

	if !bExist {
		o.Inc = append(o.Inc, vmm)
	}

	rst["id"] = vid
	rst["type"] = vtype

	return rst, nil
}
