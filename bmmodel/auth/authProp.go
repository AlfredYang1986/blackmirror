package auth

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMAuthProp struct {
	Id        string        `json:"Id"`
	Id_       bson.ObjectId `bson:"_id"`
	Auth_id   string        `json:"auth_id" bson:"auth_id"`
	Phone_id  string        `json:"phone_id" bson:"phone_id"`
	Wechat_id string        `json:"wechat_id" bson:"wechat_id"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMAuthProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMAuthProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMAuthProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMAuthProp) QueryId() string {
	return bd.Id
}

func (bd *BMAuthProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMAuthProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMAuthProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMAuthProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMAuthProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMAuthProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}
