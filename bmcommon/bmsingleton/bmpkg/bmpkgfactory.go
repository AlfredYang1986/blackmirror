package bmpkg

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmpipe"
)

var t map[string][]string = make(map[string][]string)
var k []string
var oc sync.Once

func initEPipeline() {
	//TODO：ddsaas 待提出
	t["phonelogin"] = []string{"BMAuthPhoneFindBrick"}
	t["phone2auth"] = []string{"BMPhone2AuthRSBrick", "BMAuthRS2AuthBrick", "BMAuthGenerateToken"}
	t["insertauth"] = []string{"BMPhonePushBrick", "BMWechatPushBrick",
		"BMProfilePushBrick", "BMAuthCompanyPushBrick", "BMProfileCompanyRSPushBrick", "BMAuthRSPushBrick", "BMAuthPushBrick", "BMAuthGenerateToken"}

	t["updatephone"] = []string{"BMAuthPhoneUpdateBrick"}
	t["updatewechat"] = []string{"BMAuthWechatUpdateBrick"}

	t["pushbrand"] = []string{"BMBrandPushBrick", "BMBrandPushLocationBrick", "BMBrandLocationRSPush", "BMBrandCompanyRSPush"}
	t["pushcourse"] = []string{"BMCoursePushBrick"}
	t["pushstudent"] = []string{"BMStudentPushBrick", "BMStudentRSPushBrick"}
	t["pushteacher"] = []string{"BMTeacherPushBrick"}
	t["pushlocation"] = []string{"BMLocationPushBrick"}
	t["pushclass"] = []string{"BMClassPushBrick"}
	t["pushactivity"] = []string{"BMActivityPushBrick", "BMActivityBrandRSPush"}

	t["findstudent"] = []string{"BMStudentFindBrick", "BMStudent2StudentRSBrick", "BMStudentRS2StudentBrick"}
	t["findstudents"] = []string{"BMStudentFindMultiBrick"}

	t["pushcontact"] = []string{"BMContactPushBrick", "BMOrderPushBrick", "BMContactRSPushBrick"}
	t["findcontact"] = []string{"BMContactFindBrick"}
	t["findorder"] = []string{"BMOrderFindBrick"}
	t["deleteorder"] = []string{"BMOrderDeleteBrick"}
	t["findordermulti"] = []string{"BMOrderFindMultiBrick"}

	//TODO：max 待提出
	t["maxregister"] = []string{"PHAuthProfilePush", "PHAuthCompanyPush", "PHAuthProfileRSPush",
		"PHAuthRSPushBrick", "PHAuthPushBrick", "PHAuthGenerateToken"}
	t["maxlogin"] = []string{"PHAuthFindProfileBrick", "PHProfile2AuthProp", "PHAuthProp2AuthBrick", "PHAuthGenerateToken"}
	t["maxjobgenerate"] = []string{"PHMaxJobGenerateBrick"}
	t["maxjobdelete"] = []string{"PHMaxJobDeleteBrick"}
	t["maxjobpush"] = []string{"PHMaxJobPushBrick"}
	t["maxjobsend"] = []string{"PHMaxJobSendBrick"}
	t["samplecheckselecter"] = []string{"PHSampleCheckSelecterForwardBrick"}
	t["samplecheckbody"] = []string{"PHSampleCheckBodyForwardBrick"}
	t["resultcheck"] = []string{"PHResultCheckForwardBrick"}
	t["exportmaxresult"] = []string{"PHExportMaxResultForwardBrick"}

	k = []string{
		"phonelogin", "phone2auth", "insertauth", "maxregister", "maxlogin",
	}
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
	oc.Do(initEPipeline)
	for _, itm := range k {
		if itm == pkg {
			return false
		}
	}
	return cur == 0
}
