package location

import (
	//"reflect"
	"encoding/json"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/test"
)

type Location struct {
	Id       string `json:"id" mongo:"_id"`
	Title    string `json:"title" mongo:"title"`
	Address  string `json:"address" mongo:"address"`
	District string `json:"district" mongo:"district"`

	Test test.Test `json:"test" jsonapi:"relationships"`
}

func FromJson(data string) (Location, error) {
	var rst Location
	if err := json.Unmarshal([]byte(data), &rst); err != nil {
		panic(err)
	}

	return rst, nil
}

func (loc *Location) GetTitle() string {
	rst, _ := bmmodel.AttrWithName(loc, "title", "")
	return rst.(string)
}

func (loc *Location) GetAddress() string {
	rst, _ := bmmodel.AttrWithName(loc, "address", "")
	return rst.(string)
}

func (loc *Location) GetDistrict() string {
	rst, _ := bmmodel.AttrWithName(loc, "district", "")
	return rst.(string)
}

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
