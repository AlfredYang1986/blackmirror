package bmmongo

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"gopkg.in/mgo.v2"
	"os"
	"sync"
)

var e error
var onceConfig sync.Once
var mgoSession *mgo.Session
var databaseInUse string

func GetSessionInstance() (*mgo.Session, string, error) {

	onceConfig.Do(func() {
		configPath := os.Getenv("BM_MONGO_CONF_HOME")
		profileItems := bmconfig.BMGetConfigMap(configPath)

		host := profileItems["Host"].(string)
		port := profileItems["Port"].(string)
		user := profileItems["User"].(string)
		pass := profileItems["Pass"].(string)
		database := profileItems["Database"].(string)

		url := fmt.Sprint(host, ":", port, "/", database)
		if user != "" && pass != "" {
			url = fmt.Sprint("mongodb://", user, ":", pass, "@", host, ":", port, "/", database)
		}

		session, err := mgo.Dial(url)
		mgoSession = session
		databaseInUse = database
		e = err
	})

	return mgoSession, databaseInUse, e
}