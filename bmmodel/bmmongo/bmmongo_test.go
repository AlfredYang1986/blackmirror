package bmmongo

import (
	"fmt"
	"os"
	"testing"
)

func TestGetSessionInstance(t *testing.T) {

	fmt.Println("TEST")

	os.Setenv("BM_MONGO_CONF_HOME", "../../resource/mongoconfig.json")

	//s, db, err := GetSessionInstance()
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer s.Close()
	////dbNames, err := s.DatabaseNames()
	////if err != nil {
	////	panic(err.Error())
	////}
	////fmt.Println(dbNames)
	//
	//collNames, err := s.DB(db).CollectionNames()
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(collNames)
	test1()
	test2()
	fmt.Println("DONE!")

}

func test1() {
	s, bmMgoConfig, err := GetSessionInstance()
	if err != nil {
		panic(err.Error())
	}
	defer s.Close()
	collNames, err := s.DB(bmMgoConfig.Database).CollectionNames()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("1=", collNames)
}

func test2() {
	s, bmMgoConfig, err := GetSessionInstance()
	if err != nil {
		panic(err.Error())
	}
	defer s.Close()
	collNames, err := s.DB(bmMgoConfig.Database).CollectionNames()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("2=", collNames)
}

