package authser

import (
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
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

		r.HandleFunc("/find/phone/2/rs", PhoneToAuthRS)
		r.HandleFunc("/find/rs/2/auth", AuthRS2Auth)

		r.HandleFunc("/login/phone", LoginWithPhone)
	})

	return r
}

func AuthSkeleton(w http.ResponseWriter, r *http.Request, bks bmpipe.BMBrickFace) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	sjson := string(body)

	rst, _ := jsonapi.FromJsonAPI(sjson)
	t := rst.(auth.BMAuth)

	w.Header().Add("Content-Type", "application/json")

	bks.Prepare(t)
	bks.Exec(nil)
	bks.Done()

	ec := bks.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = bks.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
