package bmrouter

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmrouter/bmoauth"
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
