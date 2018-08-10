package bmser

import (
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io/ioutil"
	"log"
	"net/http"
)

func InvokeSkeleton(w http.ResponseWriter, r *http.Request,
	bks bmpipe.BMBrickFace, f func(error)) {

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
	bks.Exec(f)
	bks.Done()

	bks.Return(w)
}
