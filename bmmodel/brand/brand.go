package brand

import (
	"encoding/json"
	"github.com/alfredyang1986/blackmirror/bmmodel"
)

type Brand struct {
	Id        string            `json:"_id", mongo:"_id"`
	Name      string            `json:"name", mongo:"name"`
	Slogan    string            `json:"slogan", mongo:"slogan"`
	Highlight []string          `json:"highlights", mongo:"heighlights"`
	About     string            `json:"about", mongo:"about"`
	Awards    map[string]string `json:"awards"`
	Attends   map[string]string `json:"attends"`
	Qualifier map[string]string `json:"qualifier"`
}

func FromJson(data string) (Brand, error) {
	var rst Brand
	if err := json.Unmarshal([]byte(data), &rst); err != nil {
		panic(err)
	}

	return rst, nil
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
