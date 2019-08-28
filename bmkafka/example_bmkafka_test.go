//Package bmkafka_test is kafka-interface's tests in BlackMirror's GoLib
package bmkafka_test

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"os"
)

func ExampleBmKafkaConfig_Produce() {

	_ = os.Setenv("BM_KAFKA_BROKER", "123.56.179.133:9092")
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", "http://123.56.179.133:8081")
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", "test20190828")
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt")
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem")
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key")
	_ = os.Setenv("BM_KAFKA_SSL_PASS", "pharbers")

	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "test"
	bkc.Produce(&topic, []byte("LaoDeng => TestKafkaProducer"))

}

func ExampleBmKafkaConfig_SubscribeTopics() {
	_ = os.Setenv("BM_KAFKA_BROKER", "123.56.179.133:9092")
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", "http://123.56.179.133:8081")
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", "test20190828")
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt")
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem")
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key")
	_ = os.Setenv("BM_KAFKA_SSL_PASS", "pharbers")

	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topics := []string{"test"}
	bkc.SubscribeTopics(topics, subscribeFunc)
}

func subscribeFunc(a interface{}) {
	fmt.Println("subscribeFunc => ", a)
	//time.Sleep(10 * time.Second)
	fmt.Println("subscribeFunc DONE!")
}
