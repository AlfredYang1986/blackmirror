package bmmongo

import (
	"github.com/alfredyang1986/blackmirror-modules/bmmodel/request"
)

type BMMongo interface {
	InsertBMObject() error
	UpdateBMObject(request.Request) error
	FindOne(request.Request) error
}

type BMMongoMulti interface {
	FindMulti(req request.Request) error
	//RetsetID
}
