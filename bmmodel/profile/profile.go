package profile

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMProfile struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	ScreenName  string `json:"screen_name" bson:"screen_name"`
	ScreenPhoto string `json:"screen_photo" bson:"screen_photo"`

	Company BMCompany `json:"company" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMProfile) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMProfile) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMProfile) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMProfile) QueryId() string {
	return bd.Id
}

func (bd *BMProfile) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMProfile) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMProfile) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "company":
		bd.Company = v.(BMCompany)
	}
	return bd
}

func (bd BMProfile) QueryConnect(tag string) interface{} {
	switch tag {
	case "company":
		return bd.Company
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMProfile) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMProfile) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMProfile) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
