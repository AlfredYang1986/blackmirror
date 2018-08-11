package authupdate

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	//"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type tBMAuthPhoneUpdateBrick struct {
	bk *bmpipe.BMBrick
}

func AuthPhoneUpdate(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	conf := bmconf.GetBMBrickConf("tBMAuthPhoneUpdateBrick")

	pfb := &tBMAuthPhoneUpdateBrick{
		bk: &bmpipe.BMBrick{
			Host:   conf.Host,
			Port:   conf.Port,
			Router: conf.Router, //"/find/rs/2/auth",
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

func (b *tBMAuthPhoneUpdateBrick) Exec(f func(error)) error {
	tmp := auth.BMPhone{}
	tmp.UpdateBMObject(*b.bk.Req)
	b.bk.Pr = tmp
	return nil
}

func (b *tBMAuthPhoneUpdateBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.bk.Req = &req
	return nil
}

func (b *tBMAuthPhoneUpdateBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMAuthPhoneUpdateBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMAuthPhoneUpdateBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMPhone)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *tBMAuthPhoneUpdateBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMPhone = b.BrickInstance().Pr.(auth.BMPhone)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
