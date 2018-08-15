package order

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

/*
    Replace entityname && Entityname
    Define Attibute1/2/... && attibute1/2/...
    Case-sensitive
*/

type Order struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`

	Title	  string 			`json:"title" bson:"title"`
	//Goods

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *Order) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *Order) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *Order) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *Order) QueryId() string {
	return bd.Id
}

func (bd *Order) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *Order) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd Order) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd Order) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *Order) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *Order) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *Order) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
