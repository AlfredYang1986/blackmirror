package main //authser

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/find"
	//"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	//"github.com/alfredyang1986/blackmirror/bmser/authser"
	"net/http"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()
	fac.RegisterModel("BMAuth", &auth.BMAuth{})
	fac.RegisterModel("BMPhone", &auth.BMPhone{})
	fac.RegisterModel("BMWechat", &auth.BMWechat{})
	fac.RegisterModel("BMAuthProp", &auth.BMAuthProp{})
	fac.RegisterModel("BMProfile", &profile.BMProfile{})
	fac.RegisterModel("BMCompany", &profile.BMCompany{})
	fac.RegisterModel("BMErrorNode", &bmerror.BMErrorNode{})
	fac.RegisterModel("request", &request.Request{})
	fac.RegisterModel("eq_condi", &request.EQCond{})
	fac.RegisterModel("up_condi", &request.UPCond{})

	fac.RegisterModel("BMAuthPhoneFindBrick", &authfind.BMAuthPhoneFindBrick{})
	//fac.RegisterModel("BMAuthPhoneFindBrickExtends", &authfind.BMAuthPhoneFindBrickExtends{})
	/* t["phonelogin"] = []string{"tBMAuthPhoneFindBrick:tBMAuthPhoneFindBrickExtends"}*/
	//t["phone2auth"] = []string{"tBMPhone2AuthRSBrick", "tBMAuthRS2AuthBrick"}
	//t["insertauth"] = []string{"tBMPhonePushBrick", "tBMWechatPushBrick",
	//"tBMProfilePushBrick", "tBMAuthRSPushBrick", "tBMAuthPushBrick"}

	//r := authser.GetRouter()
	r := bmrouter.BindRouter()
	http.ListenAndServe(":8080", r)
}
