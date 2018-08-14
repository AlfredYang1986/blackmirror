package profilepush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type tBMCompanyPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMCompanyPushBrick) Exec() error {
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

func (b *tBMCompanyPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *tBMCompanyPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
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
