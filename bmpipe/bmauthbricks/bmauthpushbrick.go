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

type tBMAuthPushBrick struct {
	bk *bmpipe.BMBrick
}

var apb *tBMAuthPushBrick
var apbo sync.Once

func AuthPushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	apbo.Do(func() {
		apb = &tBMAuthPushBrick{
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
	return apb
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMAuthPushBrick) Exec() error {
	var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
	tmp.InsertBMObject()
	return nil
}

func (b *tBMAuthPushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	b.bk.Pr = req
	return nil
}

func (b *tBMAuthPushBrick) Done(w http.ResponseWriter) error {
	fmt.Println(b.bk.Pr)
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
		//fmt.Fprintf(w, "Welcome to my website!")
	}

	return nil
}

func (b *tBMAuthPushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}
