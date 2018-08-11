package authser

import (
	"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/blackmirror/bmser"
	"net/http"
)

func PushAuth(w http.ResponseWriter, r *http.Request) {
	tmp := authpush.AuthPushBrick(nil)
	bmser.InvokeSkeleton(w, r, tmp, nil)
}

func PushPhone(w http.ResponseWriter, r *http.Request) {
	//tmp := authpush.PhonePushBrick(authpush.AuthRelationshipPushBrick(nil))
	tmp := authpush.PhonePushBrick(authpush.WechatPushBrick(nil))
	bmser.InvokeSkeleton(w, r, tmp, nil)
}

func PushWechat(w http.ResponseWriter, r *http.Request) {
	tmp := authpush.WechatPushBrick(authpush.ProfilePushBrick(nil))
	bmser.InvokeSkeleton(w, r, tmp, nil)
}

func PushProfile(w http.ResponseWriter, r *http.Request) {
	tmp := authpush.ProfilePushBrick(authpush.AuthRelationshipPushBrick(nil))
	bmser.InvokeSkeleton(w, r, tmp, nil)
}

func PushAuthRS(w http.ResponseWriter, r *http.Request) {
	tmp := authpush.AuthRelationshipPushBrick(authpush.AuthPushBrick(nil))
	bmser.InvokeSkeleton(w, r, tmp, nil)
}
