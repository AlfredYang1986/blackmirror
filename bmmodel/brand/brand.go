package brand

import (
	//"encoding/json"
	"errors"
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	//"github.com/alfredyang1986/blackmirror/bmmodel/date"
	"github.com/alfredyang1986/blackmirror/bmmodel/location"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	//"strings"
)

type Brand struct {
	Id        string            `json:"id"`
	Id_       bson.ObjectId     `bson:"_id"`
	Name      string            `json:"name" bson:"name"`
	Slogan    string            `json:"slogan" bson:"slogan"`
	Highlight []string          `json:"highlights" bson:"heighlights"`
	About     string            `json:"about" bson:"about"`
	Awards    map[string]string `json:"awards"`
	Attends   map[string]string `json:"attends"`
	Qualifier map[string]string `json:"qualifier"`
	//Found     date.DDTime       `json:"found"`

	Locations []location.Location `json:"locations" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *Brand) resetIdWithId_() {
	if bd.Id != "" {
		return
	}

	if bd.Id_.Valid() {
		bd.Id = bd.Id_.Hex()
	} else {
		panic("no id with this object")
	}
}

func (bd *Brand) resetId_WithID() {
	if bd.Id_ != "" {
		return
	}

	if bson.IsObjectIdHex(bd.Id) {
		bd.Id_ = bson.ObjectIdHex(bd.Id)
	} else {
		panic("no id with this object")
	}
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd Brand) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "locations":
		var rst []location.Location
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(location.Location))
		}
		bd.Locations = rst
	}
	return bd
}

func (bd Brand) QueryConnect(tag string) interface{} {
	switch tag {
	case "locations":
		return bd.Locations
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *Brand) InsertBMObject() error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("Brand")
	bd.resetId_WithID()

	nExist, _ := c.FindId(bd.Id_).Count()
	if nExist == 0 {
		v := reflect.ValueOf(bd).Elem()
		rst, err := bmmodel.Struct2map(v)
		rst["_id"] = bd.Id
		err = c.Insert(rst)
		return err
	} else {
		return errors.New("Only can instert not existed doc")
	}
}

func (bd *Brand) FindOne(req request.Request) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C(req.Res)
	err = c.Find(req.Cond2QueryObj()).One(bd)
	if err != nil {
		panic(err)
	}
	bd.resetIdWithId_()

	return nil
}
