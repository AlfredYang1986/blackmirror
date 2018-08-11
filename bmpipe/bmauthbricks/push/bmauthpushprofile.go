package authpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type tBMProfilePushBrick struct {
	bk *bmpipe.BMBrick
}

func ProfilePushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	conf := bmconf.GetBMBrickConf("tBMProfilePushBrick")

	apb := &tBMProfilePushBrick{
		bk: &bmpipe.BMBrick{
			Host:   conf.Host,
			Port:   conf.Port,
			Router: conf.Router, //"/auth/push",
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

func (b *tBMProfilePushBrick) Exec(f func(error)) error {
	var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)

	phone := tmp.Phone
	wechat := tmp.Wechat
	profile := tmp.Profile

	if wechat.Valid() {
		profile.ScreenName = wechat.Name
		profile.ScreenPhoto = wechat.Photo
	} else if phone.Valid() {
		profile.ScreenName = phone.Phone
		profile.ScreenPhoto = ""
	}
	tmp.Profile = profile
	err := profile.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *tBMProfilePushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	b.bk.Pr = req
	return nil
}

func (b *tBMProfilePushBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMProfilePushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMProfilePushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *tBMProfilePushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = b.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
