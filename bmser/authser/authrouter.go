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
		r.HandleFunc("/push", PushAuth)
	})

	return r
}
