package authfind

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"io"
)

type tBMAuthRS2AuthBrick struct {
	bk *bmpipe.BMBrick
}

func AuthRS2AuthBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	pfb := &tBMAuthRS2AuthBrick{
		bk: &bmpipe.BMBrick{
			Host:   "localhost",
			Port:   8080,
			Router: "/find/rs/2/auth",
			Next:   n,
			Pr:     nil,
			Req:    nil,
			Err:    0,
		},
	}
	return pfb

}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *tBMAuthRS2AuthBrick) Exec(f func(error)) error {
	prop := b.bk.Pr.(auth.BMAuthProp)
	reval, err := findAuth(prop)
	phone, err := findPhone(prop)
	wechat, err := findWechat(prop)
	reval.Phone = phone
	reval.Wechat = wechat
	b.bk.Pr = reval
	return err
}

func (b *tBMAuthRS2AuthBrick) Prepare(pr interface{}) error {
	req := pr.(auth.BMAuthProp)
	b.bk.Pr = req
	return nil
}

func (b *tBMAuthRS2AuthBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *tBMAuthRS2AuthBrick) BrickInstance() *bmpipe.BMBrick {
	return b.bk
}

func (b *tBMAuthRS2AuthBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuthProp)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

/*------------------------------------------------
 * brick inner function
 *------------------------------------------------*/

func findPhone(prop auth.BMAuthProp) (auth.BMPhone, error) {
	eq := request.EQCond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.Phone_id)
	req := request.Request{}
	req.Res = "BMPhone"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	reval := auth.BMPhone{}
	err := reval.FindOne(c.(request.Request))

	return reval, err
}

func findWechat(prop auth.BMAuthProp) (auth.BMWechat, error) {
	eq := request.EQCond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.Wechat_id)
	req := request.Request{}
	req.Res = "BMWechat"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	reval := auth.BMWechat{}
	err := reval.FindOne(c.(request.Request))

	return reval, err
}

func findAuth(prop auth.BMAuthProp) (auth.BMAuth, error) {
	eq := request.EQCond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.Auth_id)
	req := request.Request{}
	req.Res = "BMAuth"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	reval := auth.BMAuth{}
	err := reval.FindOne(c.(request.Request))

	return reval, err

}
