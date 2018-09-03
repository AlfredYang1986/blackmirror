package bmerror

import (
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"sync"
)

type tBMError struct {
	m map[int]BMErrorNode
}

type BMErrorNode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMErrorNode) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMErrorNode) QueryConnect(tag string) interface{} {
	return bd
}

var e *tBMError
var o sync.Once

func ErrInstance() *tBMError {
	o.Do(func() {
		e = &tBMError{
			m: map[int]BMErrorNode{
				-9999: BMErrorNode{Code: -9999, Message: "unknown error"},
				-1:    BMErrorNode{Code: -1, Message: "This phone already registered"},
				-2:    BMErrorNode{Code: -2, Message: "This WeChat already registered"},
				-3:    BMErrorNode{Code: -3, Message: "This course or experience_class already registered, please change name"},
				-4:    BMErrorNode{Code: -4, Message: "This company already registered, please change name"},
				-5:    BMErrorNode{Code: -5, Message: "This brand already registered, please change name"},
				-6:    BMErrorNode{Code: -6, Message: "No company found!"},
				-7:    BMErrorNode{Code: -7, Message: "No brand found!"},
				-101:  BMErrorNode{Code: -101, Message: "This user already registered"},
				-102:  BMErrorNode{Code: -102, Message: "User not found"},
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
