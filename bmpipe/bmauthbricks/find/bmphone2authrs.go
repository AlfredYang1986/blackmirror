package authfind

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
)

type tBMPhone2AuthRSBrick struct {
	bk *bmpipe.BMBrick
}

func Phone2AuthRSBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	pfb := &tBMPhone2AuthRSBrick{
		bk: &bmpipe.BMBrick{
			Host:   "localhost",
			Port:   8080,
			Router: "/find/phone/2/rs",
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

func (b *tBMPhone2AuthRSBrick) Exec(f func(error)) error {
	var tmp auth.BMPhone = b.bk.Pr.(auth.BMPhone)
	eq := request.EQCond{}
	eq.Ky = "phone_id"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "BMAuthProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	var reval auth.BMAuthProp
	err := reval.FindOne(c.(request.Request))
	b.bk.Pr = reval
	return err
}

func (b *tBMPhone2AuthRSBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMPhone)
	b.bk.Pr = req
	return nil
}

func (b *tBMPhone2AuthRSBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMPhone2AuthRSBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMPhone2AuthRSBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuthProp)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}
