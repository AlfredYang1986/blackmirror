package authser

import (
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
	})

	return r
}
