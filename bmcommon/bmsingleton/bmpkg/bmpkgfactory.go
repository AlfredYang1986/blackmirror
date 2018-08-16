package bmpkg

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"sync"
)

const (
	NoNeedAuthIdx = 3
)

var t map[string][]string = make(map[string][]string)
var oc sync.Once

func initEPipeline() {
	t["phonelogin"] = []string{"BMAuthPhoneFindBrick"}
	t["phone2auth"] = []string{"BMPhone2AuthRSBrick", "BMAuthRS2AuthBrick"}
	t["insertauth"] = []string{"BMPhonePushBrick", "BMWechatPushBrick",
		"BMProfilePushBrick", "BMAuthRSPushBrick", "BMAuthPushBrick"}

	t["updatephone"] = []string{"BMAuthPhoneUpdateBrick"}
	t["updatewechat"] = []string{"BMAuthWechatUpdateBrick"}
}

func GetPkgLen(pkg string) (int, error) {
	oc.Do(initEPipeline)

	tmp := t[pkg]
	var err error
	if tmp == nil {
		err = errors.New("query resource router error")
	}

	return len(tmp), err
}

func GetCurBrick(pkg string, idx int64) (bmpipe.BMBrickFace, error) {

	oc.Do(initEPipeline)

	tmp := t[pkg]
	var err error
	if tmp == nil {
		err = errors.New("query resource router error")
	}

	reval := tmp[idx]
	fmt.Println(reval)
	if reval == "" {
		err = errors.New("query resource router error")
	}

	face, err := bmconf.GetBMBrick(reval)
	return face, err
}

func IsNeedAuth(pkg string, cur int64) bool {
	tmp := GetNoNeedAuthSlice()

	for _, itm := range tmp {
		if itm == pkg {
			return cur == 0
		}
	}

	return true
}

func GetNoNeedAuthSlice() []string {

	oc.Do(initEPipeline)
	var reval []string
	idx := 0
	for k, _ := range t {

		if idx == NoNeedAuthIdx {
			break
		} else {
			reval = append(reval, k)
		}

		idx++
	}

	fmt.Println(reval)
	return reval
}
