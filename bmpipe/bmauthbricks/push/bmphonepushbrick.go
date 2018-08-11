package authpush

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type tBMPhonePushBrick struct {
	bk *bmpipe.BMBrick
}

func PhonePushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	conf := bmconf.GetBMBrickConf("tBMPhonePushBrick")

	appb := &tBMPhonePushBrick{
		bk: &bmpipe.BMBrick{
			Host:   conf.Host,
			Port:   conf.Port,
			Router: conf.Router, //"/auth/phone/push",
			Next:   n,
			Pr:     nil,
			Req:    nil,
			Err:    0,
		},
	}
	return appb
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMPhonePushBrick) Exec(f func(error)) error {
	var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
	ap := tmp.Phone
	if ap.Id != "" && ap.Id_.Valid() {
		if ap.Valid() && ap.IsPhoneRegisted() {
			b.bk.Err = -1
		} else {
			ap.InsertBMObject()
		}
	}
	return nil
}

func (b *tBMPhonePushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	b.bk.Pr = req
	return nil
}

func (b *tBMPhonePushBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMPhonePushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMPhonePushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *tBMPhonePushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = b.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
