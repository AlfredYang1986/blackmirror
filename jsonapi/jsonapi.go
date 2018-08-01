package jsonapi

import (
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/blackmirror/adt"
	"io"
	"log"
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
