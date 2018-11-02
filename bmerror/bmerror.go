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
	Id     string `json:"id"`
	Code   int    `json:"code"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type BMErrorObject struct {
	Errors []BMErrorNode `json:"errors"`
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
				-9999: BMErrorNode{Code: -9999, Title: "unknown error"},
				-1:    BMErrorNode{Code: -1, Title: "This phone already registered"},
				-2:    BMErrorNode{Code: -2, Title: "This WeChat already registered"},
				-3:    BMErrorNode{Code: -3, Title: "This course or experience_class already registered, please change name"},
				-4:    BMErrorNode{Code: -4, Title: "This company already registered, please change name"},
				-5:    BMErrorNode{Code: -5, Title: "This brand already registered, please change name"},
				-6:    BMErrorNode{Code: -6, Title: "No company found!"},
				-7:    BMErrorNode{Code: -7, Title: "No brand found!"},
				-8:    BMErrorNode{Code: -8, Title: "This account already registered!"},
				-101:  BMErrorNode{Code: -101, Title: "This user already registered"},
				-102:  BMErrorNode{Code: -102, Title: "User not found"},
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
	var errobjs []BMErrorNode
	if e.IsErrorDefined(ec) {
		errobjs = append(errobjs, e.m[ec])
	} else {
		panic("cannot return no defined error")
		errobjs = append(errobjs, e.m[-9999])
	}
	tmp := BMErrorObject{
		Errors:errobjs,
	}
	jsonapi.ToJsonAPIForError(&tmp, w)
}

func (e *tBMError) ErrorReval2(ec int, w http.ResponseWriter) {
	if e.IsErrorDefined(ec) {
		tmp := e.m[ec]
		jsonapi.ToJsonAPI(&tmp, w)
		//jsonapi.ToJsonAPIForError(&tmp, w)
	} else {
		panic("cannot return no defined error")
		tmp := e.m[-9999]
		jsonapi.ToJsonAPI(&tmp, w)
		//jsonapi.ToJsonAPIForError(&tmp, w)
	}
}
