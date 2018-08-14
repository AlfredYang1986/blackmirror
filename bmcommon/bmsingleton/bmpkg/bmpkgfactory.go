package bmpkg

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"sync"
)

var t map[string][]string = make(map[string][]string)
var oc sync.Once

func initEPipeline() {
	t["phonelogin"] = []string{"BMAuthPhoneFindBrick"}
	t["phone2auth"] = []string{"tBMPhone2AuthRSBrick", "tBMAuthRS2AuthBrick"}
	//t["insertauth"] = []string{"tBMPhonePushBrick", "tBMWechatPushBrick",
	//"tBMProfilePushBrick", "tBMAuthRSPushBrick", "tBMAuthPushBrick"}
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
