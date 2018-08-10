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

type tBMWechatPushBrick struct {
	bk *bmpipe.BMBrick
}

func WechatPushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	conf := bmconf.GetBMBrickConf("tBMWechatPushBrick")

	wpb := &tBMWechatPushBrick{
		bk: &bmpipe.BMBrick{
			Host:   conf.Host,
			Port:   conf.Port,
			Router: conf.Router, //"/auth/wechat/push",
			Next:   n,
			Pr:     nil,
			Req:    nil,
			Err:    0,
		},
	}
	return wpb
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMWechatPushBrick) Exec(f func(error)) error {
	var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
	aw := tmp.Wechat
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
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMWechatPushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMWechatPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *tBMWechatPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = b.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
