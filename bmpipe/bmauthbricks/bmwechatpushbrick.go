package bmauthbricks

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"sync"
)

type tBMWechatPushBrick struct {
	bk *bmpipe.BMBrick
}

var wpb *tBMWechatPushBrick
var wpbo sync.Once

func WechatPushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	wpbo.Do(func() {
		wpb = &tBMWechatPushBrick{
			bk: &bmpipe.BMBrick{
				Host:   "localhost",
				Port:   8080,
				Router: "/auth/wechat/push",
				Next:   n,
				Pr:     nil,
				Req:    nil,
				Err:    0,
			},
		}
	})
	return wpb
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMWechatPushBrick) Exec() error {
	var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
	aw := tmp.Wechat
	fmt.Println(aw)
	if aw.Id != "" && aw.Id_.Valid() {
		if aw.IsWechatRegisted() {
			b.bk.Err = -2
		} else {
			aw.InsertBMObject()
		}
	}
	return nil
}

func (b *tBMWechatPushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	b.bk.Pr = req
	return nil
}

func (b *tBMWechatPushBrick) Done() error {
	fmt.Println(b.bk.Pr)
	if b.bk.Err != 0 {
		//bmerror.ErrInstance().ErrorReval(b.bk.Err, w)
	} else {
		if b.bk.Next == nil {
			//var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
			//jsonapi.ToJsonAPI(&tmp, w)
		} else {
			nxt := b.bk.Next
			nxt.Prepare(b.bk.Pr)
			nxt.Exec()
			nxt.Done()
			b.bk.Err = b.bk.Next.BrickInstance().Err
			b.bk.Pr = b.bk.Next.BrickInstance().Pr
		}
	}

	return nil
}

func (b *tBMWechatPushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}
