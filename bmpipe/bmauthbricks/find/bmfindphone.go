package authfind

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"reflect"
)

type tBMAuthPhoneFindBrick struct {
	bk *bmpipe.BMBrick
}

func PhoneFindBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	conf := bmconf.GetBMBrickConf("tBMAuthPhoneFindBrick")

	pfb := &tBMAuthPhoneFindBrick{
		bk: &bmpipe.BMBrick{
			Host:   conf.Host,
			Port:   conf.Port,
			Router: conf.Router, //"/auth/phone/find",
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
	if f != nil {
		f(err)
	}
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
	if reflect.ValueOf(pr).Type().Name() == "BMPhone" {
		tmp := pr.(auth.BMPhone)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	} else {
		tmp := pr.(auth.BMAuth)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	}
}

func (b *tBMAuthPhoneFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = b.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
