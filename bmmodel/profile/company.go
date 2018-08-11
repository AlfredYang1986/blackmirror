package profile

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMCompany struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Name string `json:"name"`

	Profiles []BMProfile `json:"profiles" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMCompany) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMCompany) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMCompany) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMCompany) QueryId() string {
	return bd.Id
}

func (bd *BMCompany) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMCompany) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMCompany) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "profiles":
		var rst []BMProfile
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(BMProfile))
		}
		bd.Profiles = rst
	}
	return bd
}

func (bd BMCompany) QueryConnect(tag string) interface{} {
	switch tag {
	case "profiles":
		return bd.Profiles
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMCompany) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMCompany) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMCompany) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
