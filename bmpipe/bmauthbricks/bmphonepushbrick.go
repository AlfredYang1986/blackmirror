package bmauthbricks

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe"
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
				Host:   "localhost",
				Port:   8080,
				Router: "/auth/phone/push",
				Next:   n,
				Pr:     nil,
				Req:    nil,
				Err:    0,
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

func (b *tBMPhonePushBrick) Done() error {
	bmpipe.HttpPost(b)
	return nil
}

func (b *tBMPhonePushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}
