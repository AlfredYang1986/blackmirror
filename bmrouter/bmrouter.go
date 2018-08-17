package bmrouter

import (
	"github.com/alfredyang1986/blackmirror-modules/bmcommon/bmsingleton/bmpkg"
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
				vars := mux.Vars(r)
				var cur int64 = 0
				pkg := vars["package"] // the book title slug
				strcur := vars["cur"]  // the page
				if strcur != "" {
					cur, _ = strconv.ParseInt(strcur, 10, 0)
				}

				face, _ := bmpkg.GetCurBrick(pkg, cur)
				InvokeSkeleton(w, r, face, pkg, cur)
			})
	})
	return rt
}
