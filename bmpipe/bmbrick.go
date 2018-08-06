package bmpipe

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"net/http"
)

type BMBrick struct {
	Host string
	Port int

	Next *BMBrickFace

	Req  *request.Request
	Name string
	Pr   interface{}

	Err error
}

type BMBrickFace interface {
	BrickInstance() *BMBrick
	Prepare() error
	Exec() error
	Done(http.ResponseWriter) error
}
