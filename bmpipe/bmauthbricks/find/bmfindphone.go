package authfind

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
)

type tBMAuthPhoneFindBrick struct {
	bk *bmpipe.BMBrick
}

func PhoneFindBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	pfb := &tBMAuthPhoneFindBrick{
		bk: &bmpipe.BMBrick{
			Host:   "localhost",
			Port:   8080,
			Router: "/auth/phone/find",
			Next:   n,
			Pr:     nil,
			Req:    nil,
			Err:    0,
		},
	}
	return pfb
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMAuthPhoneFindBrick) Exec(f func(error)) error {
	var tmp auth.BMPhone
	err := tmp.FindOne(*b.bk.Req)
	b.bk.Pr = tmp
	f(err)
	return err
}

func (b *tBMAuthPhoneFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.bk.Req = &req
	return nil
}

func (b *tBMAuthPhoneFindBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMAuthPhoneFindBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMAuthPhoneFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMPhone)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}
