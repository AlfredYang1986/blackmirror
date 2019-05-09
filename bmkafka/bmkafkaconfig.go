package bmkafka

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"os"
	"sync"
)

type bmKafkaConfig struct {
	broker string
	group  string
	topics []string
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
			broker: profileItems["Broker"].(string),
			group: profileItems["Group"].(string),
			topics: topics,
		}
		e = nil
	})
	return config, e
}

func panicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}