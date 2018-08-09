package main //authser

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmser/authser"
	"net/http"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()
	fac.RegisterModel("BMAuth", &auth.BMAuth{})
	fac.RegisterModel("BMPhone", &auth.BMPhone{})
	fac.RegisterModel("BMWechat", &auth.BMWechat{})
	fac.RegisterModel("BMAuthProp", &auth.BMAuthProp{})
	fac.RegisterModel("BMErrorNode", &bmerror.BMErrorNode{})
	fac.RegisterModel("request", &request.Request{})
	fac.RegisterModel("eq_condi", &request.EQCond{})

	r := authser.GetRouter()
	http.ListenAndServe(":8080", r)
}
