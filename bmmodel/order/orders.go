package order

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"gopkg.in/mgo.v2/bson"
)

type Orders struct {

	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Orders	  []Order 			`json:"orders"`

}

func (bd *Orders) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Orders)
	for i,r := range bd.Orders {
		r.ResetIdWithId_()
		bd.Orders[i] = r
	}
	return  err
}