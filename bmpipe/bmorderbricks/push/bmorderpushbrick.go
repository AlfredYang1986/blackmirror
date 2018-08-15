package orderpush

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

type BMOrderPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMOrderPushBrick) Exec() error {
	con := b.bk.Pr.(contact.Contact)
	//var tmp order.Order = con.Order
	//tmp.InsertBMObject()
	for _,tmp := range con.Orders {
		tmp.InsertBMObject()
	}
	return nil
}

func (b *BMOrderPushBrick) Prepare(pr interface{}) error {
	b.BrickInstance().Pr = pr
	return nil
}

func (b *BMOrderPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMOrderPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMOrderPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(contact.Contact)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMOrderPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval contact.Contact = b.BrickInstance().Pr.(contact.Contact)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
