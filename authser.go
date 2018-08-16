package main //authser

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/find"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/update"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	//"github.com/alfredyang1986/blackmirror/bmser/authser"
	"net/http"
	"github.com/alfredyang1986/blackmirror/bmmodel/contact"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmcontactbricks/push"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmlocationbricks/push"
	"github.com/alfredyang1986/blackmirror/bmmodel/location"
	"github.com/alfredyang1986/blackmirror/bmmodel/order"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmorderbricks/push"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmcontactbricks/find"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmorderbricks/find"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()

	/*------------------------------------------------
	 * model object
	 *------------------------------------------------*/
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
	fac.RegisterModel("Contact", &contact.Contact{})
	fac.RegisterModel("BMContactProp", &contact.BMContactProp{})
	fac.RegisterModel("Location", &location.Location{})
	fac.RegisterModel("Order", &order.Order{})

	/*------------------------------------------------
	 * auth find bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthPhoneFindBrick", &authfind.BMAuthPhoneFindBrick{})
	fac.RegisterModel("BMAuthRS2AuthBrick", &authfind.BMAuthRS2AuthBrick{})
	fac.RegisterModel("BMPhone2AuthRSBrick", &authfind.BMPhone2AuthRSBrick{})
	fac.RegisterModel("BMContactFindBrick", &contactfind.BMContactFindBrick{})
	fac.RegisterModel("BMOrderFindBrick", &orderfind.BMOrderFindBrick{})
	fac.RegisterModel("BMOrderFindMultiBrick", &orderfind.BMOrderFindMultiBrick{})

	/*------------------------------------------------
	 * auth push bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMPhonePushBrick", &authpush.BMPhonePushBrick{})
	fac.RegisterModel("BMWechatPushBrick", &authpush.BMWechatPushBrick{})
	fac.RegisterModel("BMProfilePushBrick", &authpush.BMProfilePushBrick{})
	fac.RegisterModel("BMAuthRSPushBrick", &authpush.BMAuthRSPushBrick{})
	fac.RegisterModel("BMAuthPushBrick", &authpush.BMAuthPushBrick{})

	fac.RegisterModel("BMLocationPushBrick", &locationpush.BMLocationPushBrick{})
	fac.RegisterModel("BMOrderPushBrick", &orderpush.BMOrderPushBrick{})
	fac.RegisterModel("BMContactRSPushBrick", &contactpush.BMContactRSPushBrick{})
	fac.RegisterModel("BMContactPushBrick", &contactpush.BMContactPushBrick{})

	/*------------------------------------------------
	 * auth update bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("BMAuthPhoneUpdateBrick", &authupdate.BMAuthPhoneUpdateBrick{})
	fac.RegisterModel("BMAuthWechatUpdateBrick", &authupdate.BMAuthWechatUpdateBrick{})

	r := bmrouter.BindRouter()
	http.ListenAndServe(":8080", r)
}
