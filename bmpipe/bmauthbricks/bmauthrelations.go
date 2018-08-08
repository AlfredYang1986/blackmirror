package bmauthbricks

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type tBMAuthRSPushBrick struct {
	bk *bmpipe.BMBrick
}

var arsb *tBMAuthRSPushBrick
var arsbo sync.Once

func AuthRelationshipPushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	arsbo.Do(func() {
		arsb = &tBMAuthRSPushBrick{
			bk: &bmpipe.BMBrick{
				Host:   "localhost",
				Port:   8080,
				Router: "/auth/rs/push",
				Next:   n,
				Pr:     nil,
				Req:    nil,
				Err:    0,
			},
		}
	})
	return arsb

}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMAuthRSPushBrick) Exec() error {
	var tmp auth.BMAuth = b.bk.Pr.(auth.BMAuth)
	eq := request.EQCond{}
	eq.Ky = "auth_id"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "BMAuthProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	var qr auth.BMAuthProp
	err := qr.FindOne(c.(request.Request))
	if err != nil && err.Error() == "not found" {
		//panic(err)
		qr.Id_ = bson.NewObjectId()
		qr.Id = qr.Id_.Hex()
		qr.Auth_id = tmp.Id
		qr.Phone_id = tmp.Phone.Id
		qr.Wechat_id = tmp.Wechat.Id
		qr.InsertBMObject()
	}
	fmt.Println(qr)
	//tmp.InsertBMObject()
	return nil
}

func (b *tBMAuthRSPushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuth)
	b.bk.Pr = req
	return nil
}

func (b *tBMAuthRSPushBrick) Done() error {
	bmpipe.HttpPost(b)
	return nil
}

func (b *tBMAuthRSPushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}
