package authfind

import (
	"fmt"
	//"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	//"github.com/alfredyang1986/blackmirror/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"reflect"
)

type BMAuthPhoneFindBrick struct {
	bk *bmpipe.BMBrick
}

func PhoneFindBrick(n bmpipe.BMBrickFace) bmpipe.BMBrickFace {
	/* conf := bmconf.GetBMBrickConf("BMAuthPhoneFindBrick")*/

	//pfb := &BMAuthPhoneFindBrick{
	//bk: &bmpipe.BMBrick{
	//Host:   conf.Host,
	//Port:   conf.Port,
	//Router: conf.Router, //"/auth/phone/find",
	//Next:   n,
	//Pr:     nil,
	//Req:    nil,
	//Err:    0,
	//},
	/*}*/
	return nil
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAuthPhoneFindBrick) Exec(f func(error)) error {
	var tmp auth.BMPhone
	err := tmp.FindOne(*b.bk.Req)
	b.bk.Pr = tmp
	if f != nil {
		f(err)
	}
	return err
}

func (b *BMAuthPhoneFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	//b.bk.Req = &req
	return nil
}

func (b *BMAuthPhoneFindBrick) Done() error {
	bmpipe.NextBrickRemote(b)
	return nil
}

func (b *BMAuthPhoneFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAuthPhoneFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	if reflect.ValueOf(pr).Type().Name() == "BMPhone" {
		tmp := pr.(auth.BMPhone)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	} else {
		tmp := pr.(auth.BMAuth)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	}
}

func (b *BMAuthPhoneFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BMAuth = b.BrickInstance().Pr.(auth.BMAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

/*------------------------------------------------
 * brick inner error interface
 *------------------------------------------------*/

type BMAuthPhoneFindBrickExtends struct {
	bks bmpipe.BMBrickFace
}

func (tbf BMAuthPhoneFindBrickExtends) InnerErrorHandle(err error) {
	if err != nil && err.Error() == "not found" {
		fmt.Println("not found")
		/*tmp := authpush.PhonePushBrick(nil)*/
		//reval := auth.BMAuth{}
		//reval.Phone = auth.BMPhone{}
		//reval.Phone.Phone = tbf.bks.BrickInstance().Req.CondiQueryVal("phone", "BMPhone").(string)
		//tbf.bks.BrickInstance().Pr = reval
		//t*/bf.bks.BrickInstance().Next = tmp
	} else {
		fmt.Println("found")
		//tmp := authfind.Phone2AuthRSBrick(nil)
		//tmp := Phone2AuthRSBrick(nil)
		//tbf.bks.BrickInstance().Next = tmp
	}
}
