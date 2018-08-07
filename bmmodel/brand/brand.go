package brand

import (
	//"encoding/json"
	//"errors"
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	//"github.com/alfredyang1986/blackmirror/bmmodel/date"
	"github.com/alfredyang1986/blackmirror/bmmodel/location"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"reflect"
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

func (bd *Brand) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *Brand) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *Brand) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *Brand) QueryId() string {
	return bd.Id
}

func (bd *Brand) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *Brand) SetId(id string) {
	bd.Id = id
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
	return bmmodel.InsertBMObject(bd)
}

func (bd *Brand) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}
