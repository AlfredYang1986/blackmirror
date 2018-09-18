package bmrouter

import (
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmrouter/bmoauth"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"io"
	"os"
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

		rt.HandleFunc("/upload", uploadFunc)

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
					//panic(err)
					w.Header().Add("Content-Type", "application/json")
					SimpleResponseForErr(err.Error(), w)
					return
				}

				face, _ := bmpkg.GetCurBrick(pkg, cur)
				InvokeSkeleton(w, r, face, pkg, cur)
			})
	})
	return rt
}

func uploadFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "GET" {
		errMsg := "upload request method error, please use POST."
		SimpleResponseForErr(errMsg, w)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			errMsg := "upload file key error, please use key 'file'."
			SimpleResponseForErr(errMsg, w)
			return
		}
		defer file.Close()
		localDir := "resource/" + handler.Filename
		f, err := os.OpenFile(localDir, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("OpenFile error")
			fmt.Println(err)
			errMsg := "upload local file open error."
			SimpleResponseForErr(errMsg, w)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		result := map[string]string{
			"file": handler.Filename,
		}
		response := map[string]interface{}{
			"status": "ok",
			"result": result,
			"error":  "",
		}
		jso := jsonapiobj.JsResult{}
		jso.Obj = response
		enc := json.NewEncoder(w)
		enc.Encode(jso.Obj)
	}

}

func SimpleResponseForErr(errMsg string, w io.Writer) {
	response := map[string]interface{}{
		"status": 401.1,
		"result": errMsg,
		"error":  "client error",
	}
	jso := jsonapiobj.JsResult{}
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
}
