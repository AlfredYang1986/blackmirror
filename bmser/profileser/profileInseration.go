package profileser

import (
	"github.com/alfredyang1986/blackmirror/bmpipe/bmprofilebricks/push"
	"github.com/alfredyang1986/blackmirror/bmser"
	"net/http"
)

func PushCompany(w http.ResponseWriter, r *http.Request) {
	tmp := profilepush.CompanyPushBrick(nil)
	bmser.InvokeSkeleton(w, r, tmp, nil)
}

func EnrollCompany(w http.ResponseWriter, r *http.Request) {

}
