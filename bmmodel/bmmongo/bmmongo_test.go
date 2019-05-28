package bmmongo

import (
	"fmt"
	"os"
	"testing"
)

func TestGetSessionInstance(t *testing.T) {

	fmt.Println("TEST")

	os.Setenv("BM_MONGO_CONF_HOME", "../../resource/mongoconfig.json")

	s, db, err := GetSessionInstance()
	if err != nil {
		panic(err.Error())
	}
	defer s.Close()
	dbNames, err := s.DatabaseNames()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(dbNames)

	collNames, err := s.DB(db).CollectionNames()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(collNames)

}
