package brand

import (
	"encoding/json"
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/date"
	"github.com/alfredyang1986/blackmirror/bmmodel/location"
	"gopkg.in/mgo.v2/bson"
	//"reflect"
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
	Found     date.DDTime       `json:"found"`

	Locations []location.Location `json:"locations" jsonapi:"relationships"`
}

func FromJson(data string) (Brand, error) {
	var rst Brand
	if err := json.Unmarshal([]byte(data), &rst); err != nil {
		panic(err)
	}

	return rst, nil
}

func (bd *Brand) getMap(name string) map[string]string {
	rst, _ := bmmodel.AttrWithName(bd, name, bmmodel.BMJson)
	reval := make(map[string]string)
	for k, v := range rst.(map[string]interface{}) {
		reval[k] = v.(string)
	}
	return reval
}

func (bd *Brand) GetName() string {
	rst, _ := bmmodel.AttrWithName(bd, "name", "")
	return rst.(string)
}

func (bd *Brand) GetSlogan() string {
	rst, _ := bmmodel.AttrWithName(bd, "slogan", "")
	return rst.(string)
}

func (bd *Brand) GetHighlights() []string {
	rst, _ := bmmodel.AttrWithName(bd, "highlights", bmmodel.BMJson)
	var reval []string
	for _, item := range rst.([]interface{}) {
		reval = append(reval, item.(string))
	}
	return reval
}

func (bd *Brand) GetAbout() string {
	rst, _ := bmmodel.AttrWithName(bd, "about", "")
	return rst.(string)
}

func (bd *Brand) GetAwards() map[string]string {
	return bd.getMap("awards")
}

func (bd *Brand) GetAttends() map[string]string {
	return bd.getMap("attends")
}

func (bd *Brand) GetQualifier() map[string]string {
	return bd.getMap("qualifier")
}

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
