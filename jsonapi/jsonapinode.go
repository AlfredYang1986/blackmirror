package jsonapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alfredyang1986/blackmirror/adt"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/brand"
	"github.com/alfredyang1986/blackmirror/bmmodel/location"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
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
			rst[rst["type"].(string)], _ = s.mainResultParse(rst)
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

func (s *DDStm) mainResultParse(rst map[string]interface{}) (interface{}, error) {

	nid := rst["id"].(string)
	ntype := rst["type"].(string)
	var err error
	var reval interface{}

	switch ntype {
	default:
		err = errors.New("not implement")
		return reval, err
	case "brand":
		var itm brand.Brand
		s.doc.Decode(&itm)
		if bson.IsObjectIdHex(nid) {
			itm.Id = nid
			itm.Id_ = bson.ObjectIdHex(nid)
		} else {
			itm.Id_ = bson.NewObjectId()
			itm.Id = itm.Id_.Hex()
		}
		//bd.Id = nid
		reval = itm
	case "location":
		var itm location.Location
		s.doc.Decode(&itm)
		if bson.IsObjectIdHex(nid) {
			itm.Id = nid
			itm.Id_ = bson.ObjectIdHex(nid)
		} else {
			itm.Id_ = bson.NewObjectId()
			itm.Id = itm.Id_.Hex()
		}
		//loc.Id = nid
		reval = itm
	case "request":
		var req request.Request
		s.doc.Decode(&req)
		req.Id = nid
		reval = req
	case "eq_cond":
		var eq request.EQCond
		s.doc.Decode(&eq)
		eq.Id = nid
		reval = eq
	case "auth":
		var itm auth.BMAuth
		s.doc.Decode(&itm)
		if bson.IsObjectIdHex(nid) {
			itm.Id = nid
			itm.Id_ = bson.ObjectIdHex(nid)
		} else {
			itm.Id_ = bson.NewObjectId()
			itm.Id = itm.Id_.Hex()
		}
		reval = itm
	case "phone":
		var itm auth.BMPhone
		s.doc.Decode(&itm)
		if bson.IsObjectIdHex(nid) {
			itm.Id = nid
			itm.Id_ = bson.ObjectIdHex(nid)
		} else {
			itm.Id_ = bson.NewObjectId()
			itm.Id = itm.Id_.Hex()
		}
		reval = itm
	case "wechat":
		var itm auth.BMWechat
		s.doc.Decode(&itm)
		if bson.IsObjectIdHex(nid) {
			itm.Id = nid
			itm.Id_ = bson.ObjectIdHex(nid)
		} else {
			itm.Id_ = bson.NewObjectId()
			itm.Id = itm.Id_.Hex()
		}
		reval = itm
	}

	return reval, nil
}
