// Package bmkafka is kafka-interfaces in BlackMirror's GoLib
package bmkafka

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"os"
	"sync"
)

// BmKafkaConfig is BlackMirror's KafkaConfig.
// SSL used by default.
type BmKafkaConfig struct {
	Broker              string
	Group               string
	CaLocation          string
	CaSignedLocation    string
	SslKeyLocation      string
	Pass                string
	Topics              []string
	SchemaRepositoryUrl string
}

var e error
var onceConfig sync.Once
var config *BmKafkaConfig

// GetConfigInstance get the KafkaConfigInstance from config file.
func GetConfigInstance() (*BmKafkaConfig, error) {
	onceConfig.Do(func() {
		configPath := os.Getenv("BM_KAFKA_CONF_HOME")
		profileItems := bmconfig.BMGetConfigMap(configPath)
		topics := make([]string, 0)
		for _, t := range profileItems["Topics"].([]interface{}) {
			topics = append(topics, t.(string))
		}
		config = &BmKafkaConfig{
			Broker:              profileItems["Broker"].(string),
			SchemaRepositoryUrl: profileItems["SchemaRepositoryUrl"].(string),
			Group:               profileItems["Group"].(string),
			CaLocation:          profileItems["CaLocation"].(string),
			CaSignedLocation:    profileItems["CaSignedLocation"].(string),
			SslKeyLocation:      profileItems["SslKeyLocation"].(string),
			Pass:                profileItems["Pass"].(string),
			Topics:              topics,
		}
		e = nil
	})
	return config, e
}
