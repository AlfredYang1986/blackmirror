package bmerror

import (
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/hashicorp/go-uuid"
	"net/http"
	"sync"
)

type tBMError struct {
	m map[int]BmErrorNode
}

type BmErrorNode struct {
	Id     string `json:"id"`
	Code   int    `json:"code"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type BMErrorObject struct {
	Errors []BmErrorNode `json:"errors"`
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmErrorNode) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmErrorNode) QueryConnect(tag string) interface{} {
	return bd
}

var e *tBMError
var o sync.Once

func ErrInstance() *tBMError {
	o.Do(func() {
		e = &tBMError{
			m: map[int]BmErrorNode{
				-9999: BmErrorNode{Code: -9999, Title: "unknown error"},
				-9998: BmErrorNode{Code: -9998, Title: "Error! JsonAPI resolve error!", Detail: "Please check your input data!"},
				-1:    BmErrorNode{Code: -1, Title: "This phone already registered"},
				-2:    BmErrorNode{Code: -2, Title: "This WeChat already registered"},
				-3:    BmErrorNode{Code: -3, Title: "This course or experience_class already registered, please change name"},
				-4:    BmErrorNode{Code: -4, Title: "This company already registered, please change name"},
				-5:    BmErrorNode{Code: -5, Title: "This brand already registered, please change name"},
				-6:    BmErrorNode{Code: -6, Title: "No company found!"},
				-7:    BmErrorNode{Code: -7, Title: "No brand found!"},
				-8:    BmErrorNode{Code: -8, Title: "This account already registered!"},
				-9:    BmErrorNode{Code: -9, Title: "This applyee already registered! or Wrong wechat info!"},
				-10:   BmErrorNode{Code: -10, Title: "No wechat_openid found!", Detail: "No wechat_openid found!"},
				-11:   BmErrorNode{Code: -11, Title: "Account not found!", Detail: "Account not found!"},
				-101:  BmErrorNode{Code: -101, Title: "This user already registered"},
				-102:  BmErrorNode{Code: -102, Title: "User not found!"},
				-103:  BmErrorNode{Code: -103, Title: "Password error!"},
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
	var errobjs []BmErrorNode
	eid, _ := uuid.GenerateUUID()
	enode := BmErrorNode{}
	if e.IsErrorDefined(ec) {
		enode = e.m[ec]
	} else {
		enode = e.m[-9999]
	}
	enode.Id = eid
	enode.Status = 500
	errobjs = append(errobjs, enode)
	tmp := BMErrorObject{
		Errors: errobjs,
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

func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
