package authfind

import (
	"fmt"
	//"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type tBMPhone2AuthRSBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMPhone2AuthRSBrick) Exec() error {
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
	b.BrickInstance().Pr = req
	//b.bk.Pr = req
	return nil
}

func (b *tBMPhone2AuthRSBrick) Done(pkg string, idx int64, e error) error {
	//bmpipe.NextBrickRemote(b)
	bmrouter.NextBrickRemote(pkg, idx, b)
	return nil
}

func (b *tBMPhone2AuthRSBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *tBMPhone2AuthRSBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuthProp)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *tBMPhone2AuthRSBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = b.BrickInstance().Pr.(auth.BMAuth)
		//var reval auth.BMAuthProp = bks.BrickInstance().Pr.(auth.BMAuthProp)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
