package locationpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmmodel/location"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"github.com/alfredyang1986/blackmirror/bmmodel/contact"
)

type BMLocationPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMLocationPushBrick) Exec() error {
	con := b.bk.Pr.(contact.Contact)
	var tmp location.Location = con.Location //b.bk.Pr.(location.Location)
	tmp.InsertBMObject()
	return nil
}

func (b *BMLocationPushBrick) Prepare(pr interface{}) error {
	b.BrickInstance().Pr = pr
	return nil
}

func (b *BMLocationPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMLocationPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMLocationPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	//tmp := pr.(location.Location)
	tmp := pr.(contact.Contact)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMLocationPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		//var reval location.Location = b.BrickInstance().Pr.(location.Location)
		var reval contact.Contact = b.BrickInstance().Pr.(contact.Contact)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
