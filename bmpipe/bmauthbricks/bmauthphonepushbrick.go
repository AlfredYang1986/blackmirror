package bmauthbricks

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	//"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	//"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"sync"
)

type tBMAuthPhonePushBrick struct {
	bk *bmpipe.BMBrick
}

var ub *tBMAuthPhonePushBrick
var o sync.Once

func AuthPhonePushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	o.Do(func() {
		ub = &tBMAuthPhonePushBrick{
			bk: &bmpipe.BMBrick{
				Host: "localhost",
				Port: 8080,
				Next: n,
				Pr:   nil,
				Req:  nil,
				Err:  0,
			},
		}
	})
	return ub
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMAuthPhonePushBrick) Exec() error {
	var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
	fmt.Println(tmp)
	fmt.Println("fuck my boobs")
	//tmp.FindOne(*b.bk.Req)
	//b.bk.Pr = &tmp
	return nil
}

func (b *tBMAuthPhonePushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	b.bk.Pr = req
	return nil
}

func (b *tBMAuthPhonePushBrick) Done(w http.ResponseWriter) error {
	fmt.Println(b.bk.Pr)
	if b.bk.Err != 0 {
		bmerror.ErrInstance().ErrorReval(b.bk.Err, w)
	} else {
		//jsonapi.ToJsonAPI(b.bk.Pr, w)
		fmt.Fprintf(w, "Welcome to my website!")
	}

	return nil
}

func (b *tBMAuthPhonePushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}
