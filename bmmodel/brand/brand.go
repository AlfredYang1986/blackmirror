package brand

import (
	//"encoding/json"
	"errors"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/date"
	"github.com/alfredyang1986/blackmirror/bmmodel/location"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
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
		rst, err := struct2map(v)
		rst["_id"] = bd.Id
		err = c.Insert(rst)
		return err
	} else {
		return errors.New("Only can instert not existed doc")
	}
}

//UpdateBMObject(req)
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

func attrValue(v reflect.Value) (interface{}, error) {
	switch v.Kind() {
	case reflect.Invalid:
		return nil, nil
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), nil
	case reflect.String:
		return v.String(), nil
	case reflect.Array, reflect.Slice:
		var rst []interface{}
		for i := 0; i < v.Len(); i++ {
			tmp, _ := attrValue(v.Index(i))
			rst = append(rst, tmp)
		}
		return rst, nil
	case reflect.Map:
		rst := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			kv := v.MapIndex(key)
			tmp, _ := attrValue(kv)
			rst[key.String()] = tmp
		}
		return rst, nil
	}

	return 0, errors.New("not implement")
}

func struct2map(v reflect.Value) (map[string]interface{}, error) {
	rst := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a.reflect.struct.field
		fieldValue := v.Field(i)
		tag := fieldInfo.Tag // a.reflect.tag

		var name string
		if tag.Get(bmmodel.BMMongo) != "" {
			name = tag.Get(bmmodel.BMMongo)
		} else {
			name = strings.ToLower(fieldInfo.Name)
		}

		if name == "id" || name == "_id" {
			continue
		}

		ja, ok := tag.Lookup(bmmodel.BMJsonAPI)
		if ok && ja == "relationships" {
			//NOTE: relationships
			//rst[name] = "TODO"
			continue
		}

		tmp, _ := attrValue(fieldValue)
		rst[name] = tmp
	}
	fmt.Println(rst)

	return rst, nil
}
