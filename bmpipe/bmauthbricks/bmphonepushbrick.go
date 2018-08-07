package bmauthbricks

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	//"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"sync"
)

type tBMPhonePushBrick struct {
	bk *bmpipe.BMBrick
}

var appb *tBMPhonePushBrick
var appbo sync.Once

func PhonePushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	appbo.Do(func() {
		appb = &tBMPhonePushBrick{
			bk: &bmpipe.BMBrick{
				Host: "localhost",
				Port: 8080,
				Next: n,
				Pr:   nil,
				Req:  nil,
				Err:  0,
			},
		}
	})
	return appb
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMPhonePushBrick) Exec() error {
	var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
	ap := tmp.Phone
	fmt.Println(ap)
	if ap.Id != "" && ap.Id_.Valid() {
		if ap.IsPhoneRegisted() {
			b.bk.Err = -1
		} else {
			ap.InsertBMObject()
		}
	}
	return nil
}

func (b *tBMPhonePushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	b.bk.Pr = req
	return nil
}

func (b *tBMPhonePushBrick) Done(w http.ResponseWriter) error {
	fmt.Println(b.bk.Pr)
	fmt.Println(123456)
	fmt.Println(b.bk.Err)
	if b.bk.Err != 0 {
		bmerror.ErrInstance().ErrorReval(b.bk.Err, w)
	} else {
		if b.bk.Next == nil {
			var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
			jsonapi.ToJsonAPI(&tmp, w)
		} else {
			nxt := b.bk.Next
			nxt.Prepare(b.bk.Pr)
			nxt.Exec()
			nxt.Done(w)
		}
	}

	return nil
}

func (b *tBMPhonePushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}
