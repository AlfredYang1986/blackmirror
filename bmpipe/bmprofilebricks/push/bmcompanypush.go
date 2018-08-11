package profilepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type tBMCompanyPushBrick struct {
	bk *bmpipe.BMBrick
}

func CompanyPushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	conf := bmconf.GetBMBrickConf("tBMCompanyPushBrick")

	apb := &tBMCompanyPushBrick{
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

func (b *tBMCompanyPushBrick) Exec(f func(error)) error {
	var tmp profile.BMCompany = b.bk.Pr.(profile.BMCompany)
	tmp.InsertBMObject()
	b.bk.Pr = tmp
	return nil
}

func (b *tBMCompanyPushBrick) Prepare(pr interface{}) error {
	req := pr.(profile.BMCompany)
	b.bk.Pr = req
	return nil
}

func (b *tBMCompanyPushBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMCompanyPushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMCompanyPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(profile.BMCompany)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *tBMCompanyPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval profile.BMCompany = b.BrickInstance().Pr.(profile.BMCompany)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
