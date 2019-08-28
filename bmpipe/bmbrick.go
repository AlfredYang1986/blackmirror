package bmpipe

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"io"
	"net/http"
)

type BMBrick struct {
	Next BMBrickFace

	Req *request.Request
	Pr  interface{}

	Err int

	face BMBrickFace
}

type BMBrickFace interface {
	BrickInstance() *BMBrick
	Prepare(ptr interface{}) error
	Exec() error
	Done(pkg string, idx int64, e error) error
	ResultTo(w io.Writer) error
	Return(w http.ResponseWriter)
}
