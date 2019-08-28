// Package bmkafka is kafka-lib in BlackMirror's GoLibs
package bmkafka

import (
	"os"
	"sync"
)

// BmKafkaConfig is BlackMirror's KafkaConfig.
// SSL used by default.
type Config struct {
	Broker            string
	Group             string
	CaLocation        string
	CaSignedLocation  string
	SslKeyLocation    string
	SslPass           string
	SchemaRegistryUrl string
}

var e error
var onceConfig sync.Once
var config *Config

// GetConfigInstance get the KafkaConfigInstance from config file.
func GetConfigInstance() (*Config, error) {
	onceConfig.Do(func() {
		Broker := os.Getenv("BM_KAFKA_BROKER")
		SchemaRegistryUrl := os.Getenv("BM_KAFKA_SCHEMA_REGISTRY_URL")
		Group := os.Getenv("BM_KAFKA_CONSUMER_GROUP")
		CaLocation := os.Getenv("BM_KAFKA_CA_LOCATION")
		CaSignedLocation := os.Getenv("BM_KAFKA_CA_SIGNED_LOCATION")
		SslKeyLocation := os.Getenv("BM_KAFKA_SSL_KEY_LOCATION")
		SslPass := os.Getenv("BM_KAFKA_SSL_PASS")

		config = &Config{
			Broker:            Broker,
			SchemaRegistryUrl: SchemaRegistryUrl,
			Group:             Group,
			CaLocation:        CaLocation,
			CaSignedLocation:  CaSignedLocation,
			SslKeyLocation:    SslKeyLocation,
			SslPass:           SslPass,
		}
		e = nil
	})
	return config, e
}
