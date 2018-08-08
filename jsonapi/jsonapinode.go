package jsonapi

import (
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/blackmirror/adt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
	"reflect"
	"strings"
)

const (
	incr = 1
	decr
)

type DDStm struct {
	ddsk *adt.Stack // NOTE: stack for the json api pasre
	ct   string     // NOTE: current stack machine
	doc  *json.Decoder

	rst interface{} // NOTE: stm jsonapi return value
}

func STMInstance(sk *adt.Stack, pdoc *json.Decoder) DDStm {
	return DDStm{
		ddsk: sk,
		doc:  pdoc}
}

func (s *DDStm) EnterStatusWithTag(tag string) {
	s.ct = tag
	s.ddsk.PushElement(s)
}

func (s *DDStm) LeaveStatus() (interface{}, error) {
	//fmt.Println(s)
	return s.ddsk.PopElement()
}

func (s *DDStm) DetailDecoder() (interface{}, error) {

	cur := s.ct
	rst := make(map[string]interface{})
	odd := 0

	for {
		t, err := s.doc.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		strType := fmt.Sprintf("%T", t)
		strValue := fmt.Sprintf("%v", t)
		//fmt.Printf("%s : %s ==> %s\n", s.ct, strType, strValue)

		if IsMainResult(s, cur) && strValue == ATTRIBUTES {
			//rst[rst["type"].(string)], _ = s.mainResultParse(rst)
			rst["ronaldo"], _ = s.mainResultParse(rst)
			odd++
			//break
		} else if IsLeftObjDelim(strType, strValue) {
			ma := STMInstance(s.ddsk, s.doc)
			ma.EnterStatusWithTag(cur)
			rst[cur], _ = ma.DetailDecoder()
		} else if IsRightObjDelim(strType, strValue) {
			s.ddsk.PopElement()
			break
		} else if IsLeftArrayDelim(strType, strValue) {
			ma := STMInstance(s.ddsk, s.doc)
			ma.EnterStatusWithTag(cur)
			rst[cur], _ = ma.DetailDecoderList()

		} else if IsRightArrayDelim(strType, strValue) {
			s.ddsk.PopElement()
			break

		} else {
			if odd%2 == 1 && cur != "{" && cur != "[" { // NOTE: indicate key value pair
				rst[cur] = strValue
			}
		}

		odd++
		cur = strValue
	}

	ronaldo := rst["ronaldo"]
	if ronaldo != nil {
		rst[rst["type"].(string)], _ = s.map2Instance(rst["id"].(string), rst["type"].(string), ronaldo.(map[string]interface{}))
	}

	return rst, nil
}

func (s *DDStm) DetailDecoderList() ([]interface{}, error) {

	cur := s.ct
	//rst := make(map[string]interface{})
	var rst []interface{}
	//odd := 0

	for {
		t, err := s.doc.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		strType := fmt.Sprintf("%T", t)
		strValue := fmt.Sprintf("%v", t)
		//fmt.Printf("%s : %s ==> %s\n", s.ct, strType, strValue)

		/*if IsMainResult(s, cur) && strValue == ATTRIBUTES {*/
		//rst[rst["type"].(string)], _ = s.mainResultParse(rst)
		/*}*/

		if IsLeftObjDelim(strType, strValue) {
			ma := STMInstance(s.ddsk, s.doc)
			ma.EnterStatusWithTag(cur)
			//rst[cur], _ = ma.DetailDecoder()
			t, _ = ma.DetailDecoder()
			rst = append(rst, t)
			//rst = make(map[string]interface{})
		} else if IsRightObjDelim(strType, strValue) {
			s.ddsk.PopElement()
			break
		} else if IsLeftArrayDelim(strType, strValue) {
			ma := STMInstance(s.ddsk, s.doc)
			ma.EnterStatusWithTag(cur)
			t, _ = ma.DetailDecoderList()
			rst = append(rst, t)

		} else if IsRightArrayDelim(strType, strValue) {
			s.ddsk.PopElement()
			break

		} else {
			rst = append(rst, strValue)
		}

		//odd++
		cur = strValue
	}

	return rst, nil

}

func (s *DDStm) mainResultParse(rst map[string]interface{}) (map[string]interface{}, error) {
	var reval map[string]interface{}
	err := s.doc.Decode(&reval)
	return reval, err
}

func (s *DDStm) map2Instance(id string, tp string, m map[string]interface{}) (interface{}, error) {

	var nid string
	var oid bson.ObjectId
	if bson.IsObjectIdHex(id) {
		nid = id
		oid = bson.ObjectIdHex(id)
	} else {
		oid = bson.NewObjectId()
		nid = oid.Hex()
	}

	fac := bmsingleton.GetFactoryInstance()
	v, _ := fac.ReflectValue(tp)

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

		if name == "id" {
			fieldValue.SetString(nid)
		} else if name == "id_" {
			fieldValue.Set(reflect.ValueOf(oid))
		} else if name == "type" {
			fieldValue.SetString(tp)
		} else if m[name] != nil {
			vp := reflect.ValueOf(m[name]) //.Elem()
			switch fieldValue.Type().Kind() {
			default:
				fieldValue.Set(vp)
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				var f float64 = m[name].(float64)
				tmp := int64(f)
				fieldValue.SetInt(tmp)
			case reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				var f float64 = m[name].(float64)
				tmp := int64(f)
				fieldValue.SetInt(tmp)
			case reflect.Float32, reflect.Float64:
				fieldValue.SetFloat(vp.Float())
			}
		}
	}
	tmp := v.Interface()

	return tmp, nil
}
