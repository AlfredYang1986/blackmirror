package bmpipe

import (
	"bytes"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/jsonapi"
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
}

type BMBrickFace interface {
	BrickInstance() *BMBrick
	Prepare(ptr interface{}) error
	Exec() error
	Done() error
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func HttpPost(b BMBrickFace) {

	bm := b.BrickInstance().Next
	host := bm.BrickInstance().Host
	port := strconv.Itoa(bm.BrickInstance().Port)
	fmt.Println(port)
	router := bm.BrickInstance().Router

	url := strings.Join([]string{"http://", host, ":", port, router}, "")
	contentType := "application/json;charset=utf-8"

	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BMAuth)

	sb := strings.Builder{}
	jsonapi.ToJsonAPI(&tmp, &sb)

	bs := []byte(sb.String())
	body := bytes.NewBuffer(bs)

	resp, err := http.Post(url, contentType, body)
	if err != nil {
		log.Println("Post failed:", err)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read failed:", err)
		return
	}

	log.Println("content:", string(content))

	res, err := jsonapi.FromJsonAPI(string(content))
	fmt.Println("res")
	fmt.Println(res)
	v := reflect.ValueOf(res)
	if v.Type().Name() == "BMErrorNode" {
		en := res.(bmerror.BMErrorNode)
		b.BrickInstance().Err = en.Code
		b.BrickInstance().Pr = res
	}
}
