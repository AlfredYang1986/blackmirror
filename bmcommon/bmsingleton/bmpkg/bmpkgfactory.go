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
	t["generatersakey"] = []string{"BMRsaKeyGenerateBrick"}
	t["getpublickey"] = []string{"BMGetPublicKeyBrick"}
	t["insertaccount"] = []string{"BMAccountPushBrick"}
	t["accountlogin"] = []string{"BMAccountFindBrick"}
	t["phonelogin"] = []string{"BMAuthPhoneFindBrick"}
	t["phone2auth"] = []string{"BMPhone2AuthRSBrick", "BMAuthRS2AuthBrick", "BMAuthGenerateToken"}
	t["insertauth"] = []string{"BMPhonePushBrick", "BMWechatPushBrick",
		/*"BMProfilePushBrick", "BMAuthCompanyPushBrick", "BMProfileCompanyRSPushBrick", */"BMAuthRSPushBrick", "BMAuthPushBrick", "BMAuthGenerateToken"}

	t["updatephone"] = []string{"BMAuthPhoneUpdateBrick"}
	t["updatewechat"] = []string{"BMAuthWechatUpdateBrick"}

	t["insertattendee"] = []string{"BMAttendeePushPerson", "BMAttendeePushBrick", "BMAttendeePushGuardian", "BMAttendeePushPersonRS", "BMAttendeePushGuardianRS"}
	t["findattendee"] = []string{"BMAttendeeFindBrick", "BMAttendeeRS2Attendee"}
	t["findattendeemulti"] = []string{"BMAttendeeFindMulti"}
	t["updateattendee"] = []string{"BMAttendeeUpdate"}

	t["pushbrand"] = []string{"BMBrandPushBrick", "BMBrandPushLocationBrick", "BMBrandLocationRSPush", "BMBrandCompanyRSPush"}
	t["pushteacher"] = []string{"BmTeacherPushBrick", "BmTeacherPersonPushBrick", "BmTeacherPushPersonRS"}
	t["findteacher"] = []string{"BmTeacherFindBrick", "BmTeacherRS2Teacher"}
	t["findteacherprimary"] = []string{"BmPersonFindBrick", "BmPersonTeacherRS"}
	t["findteachermulti"] = []string{"BmTeacherFindMultiBrick", "BmTeacherMultiRS"}

	t["pushsessioninfo"] = []string{"BmSessionInfoPushBrick", "BmSessionCatPushBrick", "BmBindSessionCatPushBrick"}

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
		"getpublickey", "generatersakey", "insertaccount", "accountlogin", "phonelogin", "phone2auth", "insertauth", "maxregister", "maxlogin", "findteacherprimary",
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
