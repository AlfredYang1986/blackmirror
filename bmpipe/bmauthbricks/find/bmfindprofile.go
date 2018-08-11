package authfind

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	//"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type tBMFindProfileBrick struct {
	bk *bmpipe.BMBrick
}

func FindProfileBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	conf := bmconf.GetBMBrickConf("tBMFindProfileBrick")

	pfb := &tBMFindProfileBrick{
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

func (b *tBMFindProfileBrick) Exec(f func(error)) error {
	var tmp profile.BMProfile
	err := tmp.FindOne(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *tBMFindProfileBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.bk.Req = &req
	return nil
}

func (b *tBMFindProfileBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMFindProfileBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMFindProfileBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuthProp)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *tBMFindProfileBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval profile.BMProfile = b.BrickInstance().Pr.(profile.BMProfile)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
