package bmmongo

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type BMMongo interface {
	InsertBMObject() error
	UpdateBMObject(request.Request) error
	FindOne(request.Request) error
	//FindMulti(req request.Request) ([]interface{}, error)
}

type BMMongoColl interface {
	FindMulti(req request.Request) ([]interface{}, error)
	//retset
}
