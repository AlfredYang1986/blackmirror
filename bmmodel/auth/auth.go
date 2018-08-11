package auth

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMAuth struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Phone   BMPhone           `json:"phone" jsonapi:"relationships"`
	Wechat  BMWechat          `json:"wechat" jsonapi:"relationships"`
	Profile profile.BMProfile `json:"profile" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMAuth) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMAuth) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMAuth) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMAuth) QueryId() string {
	return bd.Id
}

func (bd *BMAuth) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMAuth) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMAuth) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "phone":
		bd.Phone = v.(BMPhone)
	case "wechat":
		bd.Wechat = v.(BMWechat)
	case "profile":
		bd.Profile = v.(profile.BMProfile)
	}
	return bd
}

func (bd BMAuth) QueryConnect(tag string) interface{} {
	switch tag {
	case "phone":
		return bd.Phone
	case "wechat":
		return bd.Wechat
	case "profile":
		return bd.Profile
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMAuth) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMAuth) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMAuth) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
