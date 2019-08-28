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
	//t["accountlogin"] = []string{"BMAccountFindBrick", "BmAccountGenerateToken"}

	t["phonelogin"] = []string{"BMAuthPhoneFindBrick"}
	t["phone2auth"] = []string{"BMPhone2AuthRSBrick", "BMAuthRS2AuthBrick", "BMAuthGenerateToken"}
	t["insertauth"] = []string{"BMPhonePushBrick", "BMWechatPushBrick", "BMAuthRSPushBrick", "BMAuthPushBrick", "BMAuthGenerateToken"}

	t["updatephone"] = []string{"BMAuthPhoneUpdateBrick"}
	t["updatewechat"] = []string{"BMAuthWechatUpdateBrick"}

	t["accountbindbrand"] = []string{"BmAccountBindBrand"}
	t["pushbrand"] = []string{"BmBrandPushBrick", "BmBrandPushProp", "BmBrandBindProp"}
	t["findbrand"] = []string{"BmBrandFindBrick"}
	t["findbrandmulti"] = []string{"BmBrandFindMultiBrick"}

	t["insertattendee"] = []string{"BMAttendeePushGuardian", "BMAttendeePushGuardianRS", "BMAttendeePushBrick"}
	t["findattendee"] = []string{"BMAttendeeFindBrick"}
	t["findattendeemulti"] = []string{"BMAttendeeFindMulti"}
	t["updateattendee"] = []string{"BmAttendeeUpdateBrick"}
	t["updateguardian"] = []string{"BmGuardianUpdateBrick"}

	t["pushteacher"] = []string{"BmTeacherPushBrick"}
	t["findteacher"] = []string{"BmTeacherFindBrick"}
	t["findteachermulti"] = []string{"BmTeacherFindMultiBrick"}

	t["pushroom"] = []string{"BmRoomPushBrick"}
	t["findroom"] = []string{"BmRoomFindBrick"}
	t["deleteroom"] = []string{"BmRoomDeleteBrick"}
	t["findroommulti"] = []string{"BmRoomFindMultiBrick"}

	t["pushsessioninfo"] = []string{"BmSessionInfoPushBrick", "BmSessionCatPushBrick", "BmSessionImgPushBrick", "BmSessionPushProp"}
	t["findsessioninfo"] = []string{"BmFindSessionInfoBrick"}
	t["findsessioninfomulti"] = []string{"BmFindSessionInfoMultiBrick"}
	t["updatesessioninfo"] = []string{"BmSessionInfoUpdateBrick"}
	t["updatecategory"] = []string{"BmCategoryUpdateBrick"}

	t["pushyard"] = []string{"BmYardPushBrick", "BmTagImgYardPushBrick", "BmYardPushCertificationBrick", "BmBindYardPropBrick"}
	t["findyard"] = []string{"BmYardFindBrick"}
	t["findyardmulti"] = []string{"BmYardFindMulti"}
	t["pushtagimg"] = []string{"BmTagImgPushBrick"}
	t["tagimgbindyard"] = []string{"BmTagImgBindYard"}

	t["pushreservable"] = []string{"BmReservablePushBrick", "BmBindReservableProp", "BmReservablePushSession", "BmSessionCatPushBrick", "BmSessionImgPushBrick", "BmSessionPushProp"}
	t["findreservable"] = []string{"BmReservableFindBrick"}
	t["deletereservable"] = []string{"BmReservableDeleteBrick"}
	t["findreservablemulti"] = []string{"BmReservableFindMulti"}

	t["pushapplyee"] = []string{"BmApplyeePushBrick", "BmApplyeeGenerateToken"}
	t["findapplyee"] = []string{"BmApplyeeFindBrick"}
	t["applyeegeneratetoken"] = []string{"BmApplyeeGenerateToken"}

	t["pushapply"] = []string{"BmApplyPushBrick", "BmApplyPushKids", "BmApplyPushProp"}
	t["findapply"] = []string{"BmApplyFindBrick"}
	t["findapplies"] = []string{"BmAppliesFindBrick"}
	t["findapplydetail"] = []string{"BmApplyDetailFindBrick"}
	t["findapplydetailmulti"] = []string{"BmApplyDetailFindMultiBrick"}
	t["findcount"] = []string{"BmCountFindBrick"}

	t["pushsessionable"] = []string{"BmSessionablePushBrick", "BmSessionablePushProp"}
	t["findsessionable"] = []string{"BmSessionableFindBrick"}
	t["deletesessionable"] = []string{"BmSessionableDeleteBrick"}
	t["findsessionablemulti"] = []string{"BmSessionableFindMulti"}
	t["findsessionablemultibyyard"] = []string{"BmSessionableFindMultiByYard"}

	t["pushcourseunit"] = []string{"BmCourseUnitPushBrick", "BmCourseUnitPushProp"}
	t["findcourseunit"] = []string{"BmCourseUnitFindBrick"}
	t["findcourseunitmulti"] = []string{"BmCourseUnitFindMulti"}

	t["findwechatinfo"] = []string{"BmWeChatInfoFindBrick"}

	//TODO：max 待提出
	t["maxregister"] = []string{"PHAuthProfilePush", "PHAuthCompanyPush", "PHAuthProfileRSPush",
		"PHAuthRSPushBrick", "PHAuthPushBrick", "PHAuthGenerateToken"}
	t["maxlogin"] = []string{"PHAuthFindProfileBrick", "PHProfile2AuthProp", "PHAuthProp2AuthBrick", "PHAuthGenerateToken"}
	t["maxjobgenerate"] = []string{"PHMaxJobGenerateBrick"}
	t["maxjobdelete"] = []string{"PHMaxJobDeleteBrick"}
	t["maxjobpush"] = []string{"PHMaxJobPushBrick"}
	t["maxjobpushpanel"] = []string{"PHMaxJobPushPanelBrick"}
	t["maxjobsend"] = []string{"PHMaxJobSendBrick"}
	t["samplecheckselecter"] = []string{"PHSampleCheckSelecterForwardBrick"}
	t["marketselecter"] = []string{"PhSelecterForMarketBrick"}
	t["samplecheckbody"] = []string{"PHSampleCheckBodyForwardBrick"}
	t["resultcheck"] = []string{"PHResultCheckForwardBrick"}
	t["exportmaxresult"] = []string{"PHExportMaxResultForwardBrick"}
	t["pushconfig"] = []string{"PhActionReadBrick", "PhActionPushConfigBrick"}

	k = []string{
		"getpublickey", "generatersakey", "insertaccount", "accountlogin", "phonelogin", "phone2auth", "insertauth", "maxregister", "maxlogin", "findteacherprimary",
		"pushapplyee", "findapplyee", "applyeegeneratetoken", "findwechatinfo",
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
