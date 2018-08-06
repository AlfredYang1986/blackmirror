package bmmongo

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMMongo interface {
	InsertBMObject() error
	//UpdateBMObject(req)
	FindOne(req request.Request) error
	//FindMulti(req request.Request) ([]interface{}, error)
}
