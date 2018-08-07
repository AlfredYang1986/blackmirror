package bmerror

import (
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"sync"
)

type tBMError struct {
	m map[int]tBMErrorNode
}

type tBMErrorNode struct {
	code    int    `json:"id"`
	message string `json:"message"`
}

var e *tBMError
var o sync.Once

func ErrInstance() *tBMError {
	o.Do(func() {
		e = &tBMError{
			m: map[int]tBMErrorNode{
				-9999: tBMErrorNode{code: -9999, message: "unknown error"},
				-1:    tBMErrorNode{code: -1, message: "phone already regisated"},
				-2:    tBMErrorNode{code: -2, message: "wechat already regisated"},
			},
		}
	})

	return e
}

func (e *tBMError) IsErrorDefined(ec int) bool {
	for k, _ := range e.m {
		if k == ec {
			return true
		}
	}
	return false
}

func (e *tBMError) ErrorReval(ec int, w http.ResponseWriter) {
	if e.IsErrorDefined(ec) {
		tmp := e.m[ec]
		jsonapi.ToJsonAPI(&tmp, w)
	} else {
		//panic("cannot return no defined error")
		tmp := e.m[-9999]
		jsonapi.ToJsonAPI(&tmp, w)
	}
}
