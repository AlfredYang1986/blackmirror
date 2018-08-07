package auth

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BMPhone struct {
	Id    string        `json:"id"`
	Id_   bson.ObjectId `bson:"_id"`
	Phone string        `json:"phone" bson:"phone"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMPhone) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMPhone) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMPhone) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMPhone) QueryId() string {
	return bd.Id
}

func (bd *BMPhone) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMPhone) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMPhone) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMPhone) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMPhone) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMPhone) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

/*------------------------------------------------
 * phone interface
 *------------------------------------------------*/

func (bd BMPhone) IsPhoneRegisted() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("BMPhone")
	n, err := c.Find(bson.M{"phone": bd.Phone}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd BMPhone) Valid() bool {
	return bd.Phone != ""
}
