package jsonapi

import (
	"encoding/json"
	"fmt"
	"blackmirror/adt"
	//"blackmirror/bmmodel/brand"
	"blackmirror/jsonapi/jsonapiobj"
	"io"
	"log"
	//"os"
	"strings"
)

func FromJsonAPI(jsonStream string) (interface{}, error) {
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	sk := adt.StackInstance()
	cur := "root"
	rst := make(map[string]interface{})
	odd := 0
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		strType := fmt.Sprintf("%T", t)
		strValue := fmt.Sprintf("%v", t)

		if IsLeftObjDelim(strType, strValue) {
			ma := STMInstance(&sk, dec)
			ma.EnterStatusWithTag(cur)
			rst[cur], _ = ma.DetailDecoder()
		} else if IsRightObjDelim(strType, strValue) {
			sk.PopElement()
			break
		} else if IsLeftArrayDelim(strType, strValue) {
			ma := STMInstance(&sk, dec)
			ma.EnterStatusWithTag(cur)
			rst[cur], _ = ma.DetailDecoderList()

		} else if IsRightArrayDelim(strType, strValue) {
			sk.PopElement()
			break

		} else {
			if odd%2 == 1 && cur != "{" && cur != "[" { // NOTE: indicate key value pair
				rst[cur] = strValue
			}
		}

		odd++
		cur = strValue
	}

	//return rst, nil
	return map2Object(rst)
}

func ToJsonAPI(bm interface{}, w io.Writer) error {
	jso := jsonapiobj.JsResult{}
	err := jso.FromObject(bm)
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return err
}

func ToJsonString(bm interface{}) (string, error) {
	jso := jsonapiobj.JsResult{}
	err := jso.FromObject(bm)
	out, _ := json.Marshal(jso.Obj)
	return string(out), err
}

func ToJsonAPIForError(bm interface{}, w io.Writer) error {
	//jso := jsonapiobj.JsResult{}
	enc := json.NewEncoder(w)
	enc.Encode(bm)
	return nil
}
