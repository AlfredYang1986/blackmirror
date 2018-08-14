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
	//"strconv"
	"strings"
)

type BMBrick struct {
	Next BMBrickFace

	Req *request.Request
	Pr  interface{}

	Err int
}

type BMBrickFace interface {
	BrickInstance() *BMBrick
	Prepare(ptr interface{}) error
	Exec() error
	Done(pkg string, idx int64, e error) error
	ResultTo(w io.Writer) error
	Return(w http.ResponseWriter)
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

/*func NextBrickLocal(b BMBrickFace) {*/
//nxt := b.BrickInstance().Next
//nxt.Prepare(b.BrickInstance().Pr)
//nxt.Exec(nil)
//nxt.Done()
//b.BrickInstance().Err = nxt.BrickInstance().Err
//b.BrickInstance().Pr = nxt.BrickInstance().Pr
/*}*/

func NextBrickRemote(b BMBrickFace) {

	nxt := b.BrickInstance().Next
	if b.BrickInstance().Err != 0 || nxt == nil {
		return
	}

	host := "localhost" // nxt.BrickInstance().Host
	port := "8080"      // strconv.Itoa(nxt.BrickInstance().Port)
	router := "/api/v1" //nxt.BrickInstance().Router

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
