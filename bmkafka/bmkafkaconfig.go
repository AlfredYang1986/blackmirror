package bmkafka

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"os"
	"sync"
)

type bmKafkaConfig struct {
	Broker              string
	Group               string
	Topics              []string
	SchemaRepositoryUrl string
}

var e error
var onceConfig sync.Once
var config *bmKafkaConfig

func GetConfigInstance() (*bmKafkaConfig, error) {
	onceConfig.Do(func() {
		configPath := os.Getenv("BM_KAFKA_CONF_HOME")
		profileItems := bmconfig.BMGetConfigMap(configPath)
		topics := make([]string, 0)
		for _, t := range profileItems["Topics"].([]interface{}) {
			topics = append(topics, t.(string))
		}
		config = &bmKafkaConfig{
			Broker: profileItems["Broker"].(string),
			SchemaRepositoryUrl: profileItems["SchemaRepositoryUrl"].(string),
			Group:  profileItems["Group"].(string),
			Topics: topics,
		}
		e = nil
	})
	return config, e
}
