package authser

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmser/profileser"
	"github.com/gorilla/mux"
	"sync"
)

var r *mux.Router
var o sync.Once

func GetRouter() *mux.Router {
	o.Do(func() {
		r = mux.NewRouter()
		r.HandleFunc("/auth/push", PushAuth)
		r.HandleFunc("/auth/phone/push", PushPhone)
		r.HandleFunc("/auth/wechat/push", PushWechat)
		r.HandleFunc("/auth/rs/push", PushAuthRS)
		r.HandleFunc("/auth/profile/push", PushProfile)

		r.HandleFunc("/find/phone/2/rs", PhoneToAuthRS)
		r.HandleFunc("/find/rs/2/auth", AuthRS2Auth)

		//r.HandleFunc("/auth/rs/update", UpdateAuthRS)
		r.HandleFunc("/auth/phone/update", UpdateAuthPhone)

		r.HandleFunc("/login/phone", LoginWithPhone)

		r.HandleFunc("/profile/company/push", profileser.PushCompany)
	})

	return r
}
