package bmmongo

import (
	"fmt"
	"blackmirror/bmconfighandle"
	"gopkg.in/mgo.v2"
	"os"
	"sync"
)

type BmMongoConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

var onceConfig sync.Once
var mgoSession *mgo.Session
var bmMgoConfig *BmMongoConfig

func GetSessionInstance() (*mgo.Session, *BmMongoConfig, error) {

	onceConfig.Do(func() {
		configPath := os.Getenv("BM_MONGO_CONF_HOME")
		profileItems := bmconfig.BMGetConfigMap(configPath)

		tempConfig := BmMongoConfig{}
		tempConfig.Host = profileItems["Host"].(string)
		tempConfig.Port = profileItems["Port"].(string)
		tempConfig.User = profileItems["User"].(string)
		tempConfig.Pass = profileItems["SslPass"].(string)
		tempConfig.Database = profileItems["Database"].(string)
		bmMgoConfig = &tempConfig
	})

	url := fmt.Sprint(bmMgoConfig.Host, ":", bmMgoConfig.Port, "/", bmMgoConfig.Database)
	if bmMgoConfig.User != "" && bmMgoConfig.Pass != "" {
		url = fmt.Sprint("mongodb://", bmMgoConfig.User, ":", bmMgoConfig.Pass, "@", bmMgoConfig.Host, ":", bmMgoConfig.Port, "/", bmMgoConfig.Database)
	}

	session, err := mgo.Dial(url)
	mgoSession = session

	return mgoSession, bmMgoConfig, err
}
