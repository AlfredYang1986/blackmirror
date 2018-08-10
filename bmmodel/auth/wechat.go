package auth

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BMWechat struct {
	Id      string        `json:"id"`
	Id_     bson.ObjectId `bson:"_id"`
	Open_id string        `json:"open_id" bson:"open_id"`
	Name    string        `json:"name" bson:"name"`
	Photo   string        `json:"photo" bson:"photo"`

	//TODO: 其它微信信息
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMWechat) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMWechat) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMWechat) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMWechat) QueryId() string {
	return bd.Id
}

func (bd *BMWechat) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMWechat) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMWechat) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMWechat) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMWechat) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMWechat) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMWechat) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * wechat interface
 *------------------------------------------------*/

func (bd BMWechat) IsWechatRegisted() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("BMWechat")
	n, err := c.Find(bson.M{"open_id": bd.Open_id}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd BMWechat) Valid() bool {
	return bd.Open_id != ""
}
