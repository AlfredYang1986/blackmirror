package bmrouter

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmrouter/bmoauth"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"github.com/colinmarc/hdfs"
	"github.com/hashicorp/go-uuid"
	"html/template"
	"io"
	"os"
	"time"

	//"github.com/alfredyang1986/blackmirror/bmser"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var rt *mux.Router
var o sync.Once

func BindRouter() *mux.Router {
	o.Do(func() {
		rt = mux.NewRouter()

		rt.HandleFunc("/upload", upload)

		rt.HandleFunc("/api/v1/{package}/{cur}",
			func(w http.ResponseWriter, r *http.Request) {
				vars := mux.Vars(r)
				var cur int64 = 0
				pkg := vars["package"] // the book title slug
				strcur := vars["cur"]  // the page
				if strcur != "" {
					cur, _ = strconv.ParseInt(strcur, 10, 0)
				}

				var err error
				bauth := bmpkg.IsNeedAuth(pkg, cur)
				if bauth {
					fmt.Println("need oauth")
					bearer := r.Header.Get("Authorization")
					tmp := strings.Split(bearer, " ")
					fmt.Println(tmp)
					if len(tmp) < 2 {
						err = errors.New("not authorized")
					} else {
						err = bmoauth.CheckToken(tmp[1])
					}
				}
				if err != nil {
					panic(err)
				}

				face, _ := bmpkg.GetCurBrick(pkg, cur)
				InvokeSkeleton(w, r, face, pkg, cur)
			})
	})
	return rt
}

func upload(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("method:", r.Method)
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "GET" {
		ct := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(ct, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)

	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		localDir := "resource/" + handler.Filename
		desName, _ := uuid.GenerateUUID()
		fmt.Println("des:" + desName)
		result := map[string]string{
			"file": desName,
		}
		resMap := map[string]interface{}{
			"status": "ok",
			"result": result,
		}
		jso := jsonapiobj.JsResult{}
		jso.Obj = resMap
		enc := json.NewEncoder(w)
		enc.Encode(jso.Obj)
		//fmt.Fprintf(w, "%v", jso.Obj)
		f, err := os.OpenFile(localDir, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("OpenFile error")
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		client,_ := hdfs.New("192.168.100.137:9000")
		err = client.CopyToRemote(localDir, "/client/"+desName)
		//os.Remove(localDir)
		if err != nil {
			fmt.Println("CopyToRemote error")
			fmt.Println(err)
			return
		}

	}

}
