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
