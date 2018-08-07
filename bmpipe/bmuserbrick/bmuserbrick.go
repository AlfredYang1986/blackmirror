package bmuserbrick

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/brand"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"sync"
)

type BMUserBrick struct {
	bk *bmpipe.BMBrick
}

var ub *BMUserBrick
var o sync.Once

func UserBricks() *BMUserBrick {
	o.Do(func() {
		ub = &BMUserBrick{
			bk: &bmpipe.BMBrick{
				Host: "localhost",
				Port: 8080,
				Next: nil,
				Pr:   nil,
				Req:  nil,
				Err:  0,
			},
		}
	})
	return ub
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (b *BMUserBrick) Exec() error {
	tmp := brand.Brand{}
	tmp.FindOne(*b.bk.Req)
	fmt.Println(tmp)
	b.bk.Pr = &tmp
	return nil
}

func (b *BMUserBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.bk.Req = &req
	return nil
}

func (b *BMUserBrick) Done(w http.ResponseWriter) error {
	fmt.Println(b.bk.Pr)
	if b.bk.Err != 0 {
		// TODO: 错误处理
	} else {
		jsonapi.ToJsonAPI(b.bk.Pr, w)
	}

	return nil
}

func (b *BMUserBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}
