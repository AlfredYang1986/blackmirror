package bmtests

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"os"
	"testing"
)

func TestKafkaProducer(t *testing.T) {

	os.Setenv("BM_KAFKA_CONF_HOME", "../resource/kafkaconfig.json")

	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "test"
	bkc.Produce(&topic, []byte("Su => TestKafkaProducer"))

}

func TestKafkaConsumer(t *testing.T) {

	os.Setenv("BM_KAFKA_CONF_HOME", "../resource/kafkaconfig.json")

	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topics := []string{"test"}
	bkc.SubscribeTopics(topics, subscribeFunc)

}

func subscribeFunc(a interface{}) {
	fmt.Println("subscribeFunc => ", a)
}

