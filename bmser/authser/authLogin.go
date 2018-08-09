package authser

import (
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	//"github.com/alfredyang1986/blackmirror/bmpipe"
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/find"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io/ioutil"
	"log"
	"net/http"
)

func LoginWithPhone(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	sjson := string(body)

	rst, _ := jsonapi.FromJsonAPI(sjson)
	t := rst.(request.Request)

	w.Header().Add("Content-Type", "application/json")

	bks := authfind.PhoneFindBrick(nil)
	bks.Prepare(t)

	bks.Exec(func(err error) {
		if err != nil && err.Error() == "not found" {
			tmp := authpush.PhonePushBrick(nil)
			reval := auth.BMAuth{}
			reval.Phone = auth.BMPhone{}
			reval.Phone.Phone = t.Cond[0].(request.EQCond).Vy.(string)
			bks.BrickInstance().Pr = reval
			bks.BrickInstance().Next = tmp
		} else {
			tmp := authfind.Phone2AuthRSBrick(nil)
			bks.BrickInstance().Next = tmp
		}
	})
	bks.Done()

	ec := bks.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = bks.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func PhoneToAuthRS(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	sjson := string(body)

	rst, _ := jsonapi.FromJsonAPI(sjson)
	t := rst.(auth.BMPhone)

	w.Header().Add("Content-Type", "application/json")

	bks := authfind.Phone2AuthRSBrick(authfind.AuthRS2AuthBrick(nil))
	bks.Prepare(t)
	bks.Exec(nil)
	bks.Done()

	ec := bks.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = bks.BrickInstance().Pr.(auth.BMAuth)
		//var reval auth.BMAuthProp = bks.BrickInstance().Pr.(auth.BMAuthProp)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func AuthRS2Auth(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	sjson := string(body)

	rst, _ := jsonapi.FromJsonAPI(sjson)
	t := rst.(auth.BMAuthProp)

	w.Header().Add("Content-Type", "application/json")

	bks := authfind.AuthRS2AuthBrick(nil)
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
