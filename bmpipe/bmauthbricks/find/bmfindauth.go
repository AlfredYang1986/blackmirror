package authfind

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
)

type tBMAuthFindBrick struct {
	bk *bmpipe.BMBrick
}

func AuthFindBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	afb := &tBMAuthFindBrick{
		bk: &bmpipe.BMBrick{
			Host:   "localhost",
			Port:   8080,
			Router: "/auth/find",
			Next:   n,
			Pr:     nil,
			Req:    nil,
			Err:    0,
		},
	}
	return afb
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMAuthFindBrick) Exec(f func(error)) error {
	var tmp auth.BMAuth
	tmp.FindOne(*b.bk.Req)
	return nil
}

func (b *tBMAuthFindBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	b.bk.Pr = req
	return nil
}

func (b *tBMAuthFindBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMAuthFindBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMAuthFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}
