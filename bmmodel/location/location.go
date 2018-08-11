package location

import (
	//"errors"
	//"fmt"
	//"reflect"
	//"strings"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Location struct {
	Id       string        `json:"id"`
	Id_      bson.ObjectId `bson:"_id"`
	Title    string        `json:"title" bson:"title"`
	Address  string        `json:"address" bson:"address"`
	District string        `json:"district" bson:"district"`

	//Test test.Test `json:"test" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (loc *Location) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(loc)
}

func (loc *Location) ResetId_WithID() {
	bmmodel.ResetId_WithID(loc)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *Location) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *Location) QueryId() string {
	return bd.Id
}

func (bd *Location) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *Location) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/

func (loc Location) SetConnect(tag string, v interface{}) interface{} {
	/* switch tag {*/
	//case "test":
	//loc.Test = v.(test.Test)
	/*}*/
	return loc
}

func (loc Location) QueryConnect(tag string) interface{} {
	/* switch tag {*/
	//case "test":
	//return loc.Test
	/*}*/
	return loc
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (loc *Location) InsertBMObject() error {
	return bmmodel.InsertBMObject(loc)
}

func (loc *Location) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, loc)
}

func (loc *Location) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, loc)
}
