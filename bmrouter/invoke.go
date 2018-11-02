package bmrouter //bmser

import (
	"bytes"
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once
var bmRouter bmconfig.BMRouterConfig

func InvokeSkeleton(w http.ResponseWriter, r *http.Request,
	bks bmpipe.BMBrickFace, pkg string, idx int64) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	sjson := string(body)

	w.Header().Add("Content-Type", "application/json")

	rst, _ := jsonapi.FromJsonAPI(sjson)
	bks.Prepare(rst)
	err = bks.Exec()
	bks.Done(pkg, idx, err)
	bks.Return(w)
}

//func InvokeSkeleton2(w http.ResponseWriter, r *http.Request,
//	bks bmpipe.BMBrickFace, pkg string, idx int64, wg *sync.WaitGroup, mt *sync.Mutex) {
//
//	mt.Lock()
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		log.Printf("Error reading body: %v", err)
//		http.Error(w, "can't read body", http.StatusBadRequest)
//		return
//	}
//	sjson := string(body)
//	w.Header().Add("Content-Type", "application/json")
//	rst, _ := jsonapi.FromJsonAPI(sjson)
//	bks.Prepare(rst)
//	err = bks.Exec()
//	bks.Done(pkg, idx, err)
//	bks.Return(w)
//
//	mt.Unlock()
//	wg.Done()
//}

func NextBrickRemote(pkg string, idx int64, face bmpipe.BMBrickFace) {

	once.Do(bmRouter.GenerateConfig)
	//TODO: 各个bricks待抽离
	host := bmRouter.Host // nxt.BrickInstance().Host
	port := bmRouter.Port // strconv.Itoa(nxt.BrickInstance().Port)
	router := "/api/v1/"  //nxt.BrickInstance().Router
	router += pkg
	router += "/"
	router += strconv.Itoa(int(idx))

	url := strings.Join([]string{"http://", host, ":", port, router}, "")
	contentType := "application/json;charset=utf-8"

	sb := strings.Builder{}
	face.ResultTo(&sb)

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
		face.BrickInstance().Err = en.Code
	}
	face.BrickInstance().Pr = res
}

func ForWardNextBrick(host string, port string, pkg string, idx int64, face bmpipe.BMBrickFace) {

	router := "/api/v1/" //nxt.BrickInstance().Router
	router += pkg
	router += "/"
	router += strconv.Itoa(int(idx))

	url := strings.Join([]string{"http://", host, ":", port, router}, "")
	contentType := "application/json;charset=utf-8"

	sb := strings.Builder{}
	face.ResultTo(&sb)

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
		face.BrickInstance().Err = en.Code
	}
	face.BrickInstance().Pr = res
}
