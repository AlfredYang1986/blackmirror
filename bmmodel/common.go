package bmmodel

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/alfredyang1986/blackmirror/bmmate"
	"github.com/alfredyang1986/blackmirror/bmmodel/bmmongo"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BMObject interface {
	ResetIdWithId_()
	ResetId_WithID()

	QueryObjectId() bson.ObjectId
	QueryId() string
	SetObjectId(bson.ObjectId)
	SetId(string)

	bmmongo.BMMongo
}

type NoPtr struct {
}

const (
	BMJson    string = "json"
	BMJsonAPI string = "jsonapi"
	BMMongo   string = "bson"
)

func ResetIdWithId_(ptr BMObject) {
	if ptr.QueryId() != "" {
		return
	}

	tmp := ptr.QueryObjectId()
	if tmp.Valid() {
		ptr.SetId(tmp.Hex())
	} else {
		panic("no id with this object")
	}
}

func ResetId_WithID(ptr BMObject) {
	if ptr.QueryObjectId() != "" {
		return
	}

	tmp := ptr.QueryId()
	if bson.IsObjectIdHex(tmp) {
		ptr.SetObjectId(bson.ObjectIdHex(tmp))
	} else {
		panic("no id with this object")
	}
}

func InsertBMObject(ptr BMObject) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	v := reflect.ValueOf(ptr).Elem()
	cn := v.Type().Name()
	c := session.DB("test").C(cn)
	ptr.ResetId_WithID()

	//nExist, _ := c.FindId(ptr.Id_).Count()
	nExist, _ := c.FindId(ptr.QueryObjectId).Count()
	if nExist == 0 {
		rst, err := Struct2map(v)
		rst["_id"] = ptr.QueryObjectId()
		err = c.Insert(rst)
		return err
	} else {
		return errors.New("Only can instert not existed doc")
	}
}

func FindOne(req request.Request, ptr BMObject) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C(req.Res)
	err = c.Find(req.Cond2QueryObj(req.Res)).One(ptr)
	if err != nil {
		return err
	}
	ptr.ResetIdWithId_()

	return nil
}

func DeleteOne(req request.Request, ptr BMObject) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C(req.Res)
	err = c.Find(req.Cond2QueryObj(req.Res)).One(ptr)
	if err != nil {
		return err
	}
	err = c.Remove(req.Cond2QueryObj(req.Res))
	if err != nil {
		return err
	}

	return nil
}

func FindMutil(req request.Request, ptr interface{}) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	fmu := req.FmCond[0]
	skip := (fmu.Page - 1) * fmu.Take

	c := session.DB("test").C(req.Res)
	err = c.Find(req.Cond2QueryObj(req.Res)).Skip(skip).Limit(fmu.Take).All(ptr)

	return err
}

func UpdateOne(req request.Request, ptr BMObject) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return errors.New("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C(req.Res)
	err = c.Find(req.Cond2QueryObj(req.Res)).One(ptr)
	if err != nil {
		return err
	}

	up := req.Cond2UpdateObj()
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a.reflect.struct.field
		fieldValue := v.Field(i)
		tag := fieldInfo.Tag // a.reflect.tag

		var name string
		if tag.Get(BMMongo) != "" {
			name = tag.Get(BMMongo)
		} else {
			name = strings.ToLower(fieldInfo.Name)
		}

		if up[name] != nil {
			fieldValue.Set(reflect.ValueOf(up[name]))
		}
	}
	ptr.ResetIdWithId_()
	err = c.Update(bson.M{"_id": ptr.QueryObjectId()}, ptr)

	return err

}

func AttrWithName(ptr interface{}, attr string, tagN string) (interface{}, error) {
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a.reflect.struct.field
		fieldValue := v.Field(i)
		tag := fieldInfo.Tag // a.reflect.tag

		var name string
		if tagN == BMJson {
			name = tag.Get(BMJson)
		} else if tagN == BMMongo {
			name = tag.Get(BMMongo)
		} else {
			name = strings.ToLower(fieldInfo.Name)
		}

		if name == attr {
			return AttrValue(fieldValue)
		}
	}

	return NoPtr{}, nil
}

func AttrValue(v reflect.Value) (interface{}, error) {
	switch v.Kind() {
	case reflect.Invalid:
		return nil, nil
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	case reflect.String:
		return v.String(), nil
	case reflect.Array, reflect.Slice:
		var rst []interface{}
		for i := 0; i < v.Len(); i++ {
			tmp, _ := AttrValue(v.Index(i))
			rst = append(rst, tmp)
		}
		return rst, nil
	case reflect.Map:
		rst := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			kv := v.MapIndex(key)
			tmp, _ := AttrValue(kv)
			rst[key.String()] = tmp
		}
		return rst, nil
	case reflect.Interface:
		if bmmate.IsStruct(v) {
			if reflect.TypeOf(v.Interface()).Kind() == reflect.String {
				return AttrValue(reflect.ValueOf(v.Interface()))
			} else {
				return AttrValue(reflect.ValueOf(v.Interface()))
			}
		} else {
			return AttrValue(reflect.ValueOf(v.Interface()))
		}
	}

	return NoPtr{}, errors.New("not implement")
}

func Struct2map(v reflect.Value) (map[string]interface{}, error) {
	rst := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a.reflect.struct.field
		fieldValue := v.Field(i)
		tag := fieldInfo.Tag // a.reflect.tag

		var name string
		if tag.Get(BMMongo) != "" {
			name = tag.Get(BMMongo)
		} else {
			name = strings.ToLower(fieldInfo.Name)
		}

		if name == "id" || name == "_id" {
			continue
		}

		ja, ok := tag.Lookup(BMJsonAPI)
		if ok && ja == "relationships" {
			//NOTE: relationships
			//rst[name] = "TODO"
			continue
		}

		tmp, _ := AttrValue(fieldValue)
		rst[name] = tmp
	}
	fmt.Println(rst)

	return rst, nil
}
