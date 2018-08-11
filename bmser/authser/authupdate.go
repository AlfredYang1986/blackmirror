package authser

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/update"
	"github.com/alfredyang1986/blackmirror/bmser"
	"net/http"
)

func UpdateAuthPhone(w http.ResponseWriter, r *http.Request) {
	bks := authupdate.AuthPhoneUpdate(nil)
	bmser.InvokeSkeleton(w, r, bks, nil)
}

func UpdateWechatPhone(w http.ResponseWriter, r *http.Request) {
	bks := authupdate.AuthWechatUpdate(nil)
	bmser.InvokeSkeleton(w, r, bks, nil)
}
