package bmpipe

import (
	"bytes"
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	//"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	//"errors"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type BMBrick struct {
	Host   string
	Port   int
	Router string

	Next BMBrickFace

	Req  *request.Request
	Name string
	Pr   interface{}

	Err int

	face BMBrickFace

	/* PrepareFunc func(ptr interface{}) error*/
	//ExecFunc    func(func(error)) error
	//DoneFunc    func() error
	//Result2Func func(io.Writer) error
	/*ReturnFunc  func(http.ResponseWriter)*/
}

type BMBrickFace interface {
	//FromJsonToInstance() interface{}
	BrickInstance() *BMBrick
	Prepare(ptr interface{}) error
	Exec(func(error)) error
	Done() error
	ResultTo(w io.Writer) error
	Return(w http.ResponseWriter)
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func NextBrickLocal(b BMBrickFace) {
	nxt := b.BrickInstance().Next
	nxt.Prepare(b.BrickInstance().Pr)
	nxt.Exec(nil)
	nxt.Done()
	b.BrickInstance().Err = nxt.BrickInstance().Err
	b.BrickInstance().Pr = nxt.BrickInstance().Pr
}

func NextBrickRemote(b BMBrickFace) {

	nxt := b.BrickInstance().Next
	if b.BrickInstance().Err != 0 || nxt == nil {
		return
	}

	host := nxt.BrickInstance().Host
	port := strconv.Itoa(nxt.BrickInstance().Port)
	router := nxt.BrickInstance().Router

	url := strings.Join([]string{"http://", host, ":", port, router}, "")
	contentType := "application/json;charset=utf-8"

	sb := strings.Builder{}
	b.ResultTo(&sb)

	bs := []byte(sb.String())
	body := bytes.NewBuffer(bs)

	resp, err := http.Post(url, contentType, body)
	if err != nil {
		log.Println("Post failed:", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read failed:", err)
		return
	}
	resp.Body.Close()

	res, err := jsonapi.FromJsonAPI(string(content))
	v := reflect.ValueOf(res)
	if v.Type().Name() == "BMErrorNode" {
		en := res.(bmerror.BMErrorNode)
		b.BrickInstance().Err = en.Code
	}
	b.BrickInstance().Pr = res
}

/*func (bk *BMBrick) BrickInstance() *BMBrick {*/
//return bk
//}

//func (bk *BMBrick) Prepare(ptr interface{}) error {
//if bk.PrepareFunc != nil {
//return bk.Prepare(ptr)
//} else {
//return errors.New("brick init error, this brick don't have prepare function")
//}
//}

//func (bk *BMBrick) Exec(f func(error)) error {
//if bk.ExecFunc != nil {
//return bk.ExecFunc(f)
//} else {
//return errors.New("brick init error, this brick don't have exec function")
//}
//}

//func (bk *BMBrick) Done() error {
//if bk.DoneFunc != nil {
//return bk.DoneFunc()
//} else {
//NextBrickRemote(bk)
//return errors.New("brick init error, this brick don't have done function")
//}
////NextBrickRemote(bk)
//}

//func (bk *BMBrick) ResultTo(w io.Writer) error {

//if bk.Result2Func != nil {
//return bk.Result2Func(w)
//} else {
//return errors.New("brick init error, this brick don't have result to function")
//}
//}

//func (b *BMBrick) Return(w http.ResponseWriter) {
//if b.ReturnFunc != nil {
//b.ReturnFunc(w)
//} else {
//errors.New("brick init error, this brick don't have return function")
//}
/*}*/
