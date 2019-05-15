package bmtests

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"os"
	"testing"
	"time"
)

func TestKafkaProducer(t *testing.T) {

	os.Setenv("BM_KAFKA_CONF_HOME", "../resource/kafkaconfig.json")

	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "test"
	bkc.Produce(&topic, []byte("LaoDeng => TestKafkaProducer"))

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
	time.Sleep(10 * time.Second)
	fmt.Println("subscribeFunc DONE!")
}

