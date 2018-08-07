package authser

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io/ioutil"
	"log"
	"net/http"
)

func PushAuth(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	sjson := string(body)
	fmt.Println(sjson)

	rst, _ := jsonapi.FromJsonAPI(sjson)
	println(rst)

	t := rst.(auth.BMAuth)
	//fmt.Println(t)

	w.Header().Add("Content-Type", "application/json")

	tmp :=
		bmauthbricks.AuthPushBrick(
			bmauthbricks.PhonePushBrick(
				bmauthbricks.WechatPushBrick(nil),
			))

	tmp.Prepare(t)
	tmp.Exec()
	tmp.Done(w)

	//err = t.InsertBMObject()
	//fmt.Fprintf(w, "Welcome to my website!")
}
