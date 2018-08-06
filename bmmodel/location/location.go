package location

import (
	"errors"
	//"fmt"
	"reflect"
	//"strings"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmmodel/test"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Location struct {
	Id       string        `json:"id"`
	Id_      bson.ObjectId `bson:"_id"`
	Title    string        `json:"title" bson:"title"`
	Address  string        `json:"address" bson:"address"`
	District string        `json:"district" bson:"district"`

	Test test.Test `json:"test" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (loc *Location) resetIdWithId_() {
	if loc.Id != "" {
		return
	}

	if loc.Id_.Valid() {
		loc.Id = loc.Id_.Hex()
	} else {
		panic("no id with this object")
	}
}

func (loc *Location) resetId_WithID() {
	if loc.Id_ != "" {
		return
	}

	if bson.IsObjectIdHex(loc.Id) {
		loc.Id_ = bson.ObjectIdHex(loc.Id)
	} else {
		panic("no id with this object")
	}
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/

func (loc Location) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "test":
		loc.Test = v.(test.Test)
	}
	return loc
}

func (loc Location) QueryConnect(tag string) interface{} {
	//return loc.Relationships[tag]
	switch tag {
	case "test":
		return loc.Test
	}
	return loc
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (loc *Location) InsertBMObject() error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("Brand")
	loc.resetId_WithID()

	nExist, _ := c.FindId(loc.Id_).Count()
	if nExist == 0 {
		v := reflect.ValueOf(loc).Elem()
		rst, err := bmmodel.Struct2map(v)
		rst["_id"] = loc.Id
		err = c.Insert(rst)
		return err
	} else {
		return errors.New("Only can instert not existed doc")
	}
}

func (loc *Location) FindOne(req request.Request) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C(req.Res)
	err = c.Find(req.Cond2QueryObj()).One(loc)
	if err != nil {
		panic(err)
	}
	loc.resetIdWithId_()

	return nil
}
