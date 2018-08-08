package main //authser

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmser/authser"
	"net/http"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()
	fac.RegisterModel("BMAuth", &auth.BMAuth{})
	fac.RegisterModel("BMPhone", &auth.BMPhone{})
	fac.RegisterModel("BMWechat", &auth.BMWechat{})
	fac.RegisterModel("BMErrorNode", &bmerror.BMErrorNode{})

	r := authser.GetRouter()
	http.ListenAndServe(":8080", r)
}
