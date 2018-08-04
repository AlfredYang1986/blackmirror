package bmmongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"log"
	"errors"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/brand"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"reflect"
	"strings"
)

func InsertBMObject(ptr interface{}) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	v := reflect.ValueOf(ptr).Elem()
	rst, err := struct2map(v)

	cn := v.Type().Name()
	c := session.DB("test").C(cn)

	nExist, _ := c.FindId(rst["_id"]).Count()
	if nExist == 0 {
		rst["_id"] = bson.NewObjectId()
		fmt.Println(rst)
		err := c.Insert(rst)
		if err != nil {
			panic(err)
		}
		return nil
	} else {
		return errors.New("Only can instert not existed doc")
	}
}

func UpdateBMObject(ptr interface{}, nc []string) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	v := reflect.ValueOf(ptr).Elem()
	cn := v.Type().Name()

	rst, err := struct2map(v)
	fmt.Println(9999)
	fmt.Println(rst)
	//oid := rst["_id"].(string)
	var oid bson.ObjectId
	if bson.IsObjectIdHex(rst["_id"].(string)) {
		oid = bson.ObjectIdHex(rst["_id"].(string))
		fmt.Println(oid)
	} else {
		err = errors.New("need ObjectId to continue")
	}

	m := make(map[string]interface{})
	c := session.DB("test").C(cn)
	err = c.Find(bson.M{"_id": oid}).One(&m)
	//err = c.Find(bson.M{"name": "alfredyang"}).One(&m)
	if err != nil {
		return err
	}
	fmt.Println(m)

	for _, prop := range nc {
		m[prop] = rst[prop]
	}

	err = c.Update(bson.M{"_id": oid}, m)
	//err = c.Update(bson.M{"name": "alfredyang"}, m)
	if err != nil {
		return err
	}

	return nil
}

func FindOne(req request.Request) (interface{}, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C(req.Res)

	reval := brand.Brand{}
	//reflect.New(t).Elem().Interface()
	err = c.Find(req.Cond2QueryObj()).One(&reval)
	reval.Id = reval.Id_.Hex()
	if err != nil {
		panic(err)
		return 0, nil
	}
	fmt.Println(reval)

	return reval, nil
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

		if name == "id" {
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
