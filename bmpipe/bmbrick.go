package bmpipe

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"net/http"
)

type BMBrick struct {
	Host string
	Port int

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
	Done(http.ResponseWriter) error
}
