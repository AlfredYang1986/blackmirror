package bmrouter

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	//"github.com/alfredyang1986/blackmirror/bmser"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"sync"
)

var rt *mux.Router
var o sync.Once

func BindRouter() *mux.Router {
	o.Do(func() {
		rt = mux.NewRouter()

		rt.HandleFunc("/api/v1/{package}/{cur}",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Println(1234)
				vars := mux.Vars(r)
				var cur int64 = 0
				pkg := vars["package"] // the book title slug
				fmt.Println(pkg)
				strcur := vars["cur"] // the page
				fmt.Println(strcur)
				if strcur != "" {
					cur, _ = strconv.ParseInt(strcur, 10, 0)
				}

				face, _ := bmpkg.GetCurBrick(pkg, cur)
				InvokeSkeleton(w, r, face, pkg, cur)
			})
	})
	return rt
}
