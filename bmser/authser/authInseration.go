package authser

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io/ioutil"
	"log"
	"net/http"
)

func authPushSkeleton(w http.ResponseWriter, r *http.Request, bks bmpipe.BMBrickFace) {
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
	bks.Exec()
	bks.Done()

	ec := bks.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = bks.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func PushAuth(w http.ResponseWriter, r *http.Request) {
	tmp :=
		bmauthbricks.PhonePushBrick(
			bmauthbricks.WechatPushBrick(
				bmauthbricks.AuthRelationshipPushBrick(nil)))
		//bmauthbricks.AuthPushBrick(nil))))
	authPushSkeleton(w, r, tmp)
}

func PushPhone(w http.ResponseWriter, r *http.Request) {
	tmp := bmauthbricks.PhonePushBrick(nil)
	authPushSkeleton(w, r, tmp)
}

func PushWechat(w http.ResponseWriter, r *http.Request) {
	tmp := bmauthbricks.WechatPushBrick(nil)
	authPushSkeleton(w, r, tmp)
}

func PushAuthRS(w http.ResponseWriter, r *http.Request) {
	tmp := bmauthbricks.AuthRelationshipPushBrick(nil)
	authPushSkeleton(w, r, tmp)
}
