//Package bmkafka_test is kafka-interface's tests in BlackMirror's GoLib
package bmkafka_test

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"os"
)

func ExampleBmKafkaConfig_Produce() {

	os.Setenv("BM_KAFKA_CONF_HOME", "../resource/kafkaconfig.json")

	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "test"
	bkc.Produce(&topic, []byte("LaoDeng => TestKafkaProducer"))

}

func ExampleBmKafkaConfig_SubscribeTopics() {
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
	//time.Sleep(10 * time.Second)
	fmt.Println("subscribeFunc DONE!")
}
