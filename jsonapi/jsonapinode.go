package jsonapi

import (
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/blackmirror/adt"
	"io"
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
	fmt.Println(s)
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
			panic("some thing error with decode")
		}

		strType := fmt.Sprintf("%T", t)
		strValue := fmt.Sprintf("%v", t)
		//fmt.Printf("%s : %s ==> %s\n", s.ct, strType, strValue)

		if IsLeftObjDelim(strType, strValue) {
			ma := STMInstance(s.ddsk, s.doc)
			ma.EnterStatusWithTag(cur)
			rst[cur], _ = ma.DetailDecoder()
		} else if IsRightObjDelim(strType, strValue) {
			s.ddsk.PopElement()
			break
		} else if IsLeftArrayDelim(strType, strValue) {

		} else if IsRightArrayDelim(strType, strValue) {
			break

		} else {
			if odd%2 == 1 {
				rst[cur] = strValue
			}
		}

		odd++
		cur = strValue
	}

	return rst, nil
}
