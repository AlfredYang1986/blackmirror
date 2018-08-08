package bmauthbricks

import (
	//"bytes"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	//"github.com/alfredyang1986/blackmirror/jsonapi"
	//"io/ioutil"
	//"log"
	//"net/http"
	//"strings"
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
				Host:   "localhost",
				Port:   8080,
				Router: "/auth/push",
				Next:   n,
				Pr:     nil,
				Req:    nil,
				Err:    0,
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

func (b *tBMAuthPushBrick) Done() error {
	fmt.Println(b.bk.Pr)
	if b.bk.Err != 0 {
		//bmerror.ErrInstance().ErrorReval(b.bk.Err, w)
	} else {
		if b.bk.Next == nil {
			//var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
			//jsonapi.ToJsonAPI(&tmp, w)
		} else {

			// NOTE: Local query
			/* nxt := b.bk.Next*/
			//nxt.Prepare(b.bk.Pr)
			//nxt.Exec()
			//nxt.Done()
			//b.bk.Err = b.bk.Next.BrickInstance().Err
			//b.bk.Pr = b.bk.Next.BrickInstance().Pr

			// NOTE: Remote query
			bmpipe.HttpPost(b)
		}
		//fmt.Fprintf(w, "Welcome to my website!")
	}

	return nil
}

func (b *tBMAuthPushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}
