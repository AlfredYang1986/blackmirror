package contactpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmmodel/contact"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type BMContactPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMContactPushBrick) Exec() error {
	var tmp contact.Contact = b.bk.Pr.(contact.Contact)
	tmp.InsertBMObject()
	b.bk.Pr = tmp
	return nil
}

func (b *BMContactPushBrick) Prepare(pr interface{}) error {
	req := pr.(contact.Contact)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *BMContactPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMContactPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMContactPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(contact.Contact)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMContactPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval contact.Contact = b.BrickInstance().Pr.(contact.Contact)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

