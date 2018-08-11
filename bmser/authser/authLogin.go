package authser

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	//"github.com/alfredyang1986/blackmirror/bmmodel/request"
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/find"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/blackmirror/bmser"
	"net/http"
)

func LoginWithPhone(w http.ResponseWriter, r *http.Request) {
	bks := authfind.PhoneFindBrick(nil)
	bmser.InvokeSkeleton(w, r, bks, func(err error) {
		if err != nil && err.Error() == "not found" {
			tmp := authpush.PhonePushBrick(nil)
			reval := auth.BMAuth{}
			reval.Phone = auth.BMPhone{}
			reval.Phone.Phone = bks.BrickInstance().Req.CondiQueryVal("phone", "BMPhone").(string)
			bks.BrickInstance().Pr = reval
			bks.BrickInstance().Next = tmp
		} else {
			tmp := authfind.Phone2AuthRSBrick(nil)
			bks.BrickInstance().Next = tmp
		}
	})
}

func PhoneToAuthRS(w http.ResponseWriter, r *http.Request) {
	bks := authfind.Phone2AuthRSBrick(authfind.AuthRS2AuthBrick(nil))
	bmser.InvokeSkeleton(w, r, bks, nil)
}

func AuthRS2Auth(w http.ResponseWriter, r *http.Request) {
	bks := authfind.AuthRS2AuthBrick(nil)
	bmser.InvokeSkeleton(w, r, bks, nil)
}
