package location

import (
	"encoding/json"
	"github.com/alfredyang1986/blackmirror/bmmodel"
)

type Location struct {
	Id       string `json:"_id", mongo:"_id"`
	Title    string `json:"title", mongo:"title"`
	Address  string `json:"address", mongo:"address"`
	District string `json:"district", mongo:"district"`

	Relationships map[string]interface{}
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
	if loc.Relationships == nil {
		loc.Relationships = make(map[string]interface{})
	}
	loc.Relationships[tag] = v
	return loc
}

func (loc Location) QueryConnect(tag string) interface{} {
	return loc.Relationships[tag]
}
