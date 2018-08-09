package authpush

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
)

type tBMAuthPushBrick struct {
	bk *bmpipe.BMBrick
}

func AuthPushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	apb := &tBMAuthPushBrick{
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
	return apb
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMAuthPushBrick) Exec(f func(error)) error {
	var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
	tmp.InsertBMObject()
	b.bk.Pr = tmp
	return nil
}

func (b *tBMAuthPushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	b.bk.Pr = req
	return nil
}

func (b *tBMAuthPushBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMAuthPushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMAuthPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}
