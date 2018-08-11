package authpush

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type tBMAuthRSPushBrick struct {
	bk *bmpipe.BMBrick
}

func AuthRelationshipPushBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	conf := bmconf.GetBMBrickConf("tBMAuthRSPushBrick")

	arsb := &tBMAuthRSPushBrick{
		bk: &bmpipe.BMBrick{
			Host:   conf.Host,
			Port:   conf.Port,
			Router: conf.Router, //"/auth/rs/push",
			Next:   n,
			Pr:     nil,
			Req:    nil,
			Err:    0,
		},
	}
	return arsb
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMAuthRSPushBrick) Exec(f func(error)) error {
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
		qr.Profile_id = tmp.Profile.Id
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
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMAuthRSPushBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMAuthRSPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *tBMAuthRSPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = b.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
