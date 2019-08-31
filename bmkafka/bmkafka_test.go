package bmkafka

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	"github.com/hashicorp/go-uuid"
	"os"
	"testing"
	"time"
)

func TestKafkaProducer(t *testing.T) {

	_ = os.Setenv("BM_KAFKA_BROKER", "123.56.179.133:9092")
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", "http://123.56.179.133:8081")
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", "test20190828")
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt")
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem")
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key")
	_ = os.Setenv("BM_KAFKA_SSL_PASS", "pharbers")

	bkc, err := GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "test"
	bkc.Produce(&topic, []byte("NEW BLOOD!!!"))

}

func TestKafkaConsumer(t *testing.T) {

	_ = os.Setenv("BM_KAFKA_BROKER", "123.56.179.133:9092")
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", "http://123.56.179.133:8081")
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", "test20190828")
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt")
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem")
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key")
	_ = os.Setenv("BM_KAFKA_SSL_PASS", "pharbers")

	bkc, err := GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topics := []string{"test"}
	bkc.SubscribeTopics(topics, subscribeFunc)

	time.Sleep(10 * time.Minute)

}

func subscribeFunc(a interface{}) {
	fmt.Println("subscribeFunc => ", string(a.([]byte)))
	//time.Sleep(10 * time.Second)
	fmt.Println("subscribeFunc DONE!")
}

func TestKafkaConsumerMap(t *testing.T) {

	_ = os.Setenv("BM_KAFKA_BROKER", "123.56.179.133:9092")
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", "http://123.56.179.133:8081")
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", "test20190830")
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt")
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem")
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key")
	_ = os.Setenv("BM_KAFKA_SSL_PASS", "pharbers")

	bkc, err := GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}

	go func() {
		topics := []string{"test"}
		bkc.SubscribeTopics(topics, subscribeFunc)
	}()

	go func() {
		topics := []string{"test2"}
		bkc.SubscribeTopics(topics, subscribeFunc2)
	}()

	time.Sleep(10 * time.Minute)

}

func subscribeFunc2(a interface{}) {
	fmt.Println("subscribeFunc2 => ", string(a.([]byte)))
	//time.Sleep(10 * time.Second)
	fmt.Println("subscribeFunc2 DONE!")
}

func TestKafkaAvroProducer(t *testing.T) {

	_ = os.Setenv("BM_KAFKA_BROKER", "123.56.179.133:9092")
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", "http://123.56.179.133:8081")
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", "test20190828")
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt")
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem")
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key")
	_ = os.Setenv("BM_KAFKA_SSL_PASS", "pharbers")

	var schemaRepositoryUrl = os.Getenv("BM_KAFKA_SCHEMA_REGISTRY_URL")
	var rawMetricsSchema = `{"type": "record","name": "RecordDemo","namespace": "com.pharbers.kafka.schema","fields": [{"name": "id", "type": "string"},{"name": "name",  "type": "string" }]}`

	encoder := kafkaAvro.NewKafkaAvroEncoder(schemaRepositoryUrl)
	schema, err := avro.ParseSchema(rawMetricsSchema)
	bmerror.PanicError(err)
	record := avro.NewGenericRecord(schema)
	tmpUUID, err := uuid.GenerateUUID()
	bmerror.PanicError(err)
	record.Set("id", tmpUUID)
	record.Set("name", "test@max.logic")
	//record.Set("msg", "hello1")
	//record.Set("msg", "hello2")
	recordByteArr, err := encoder.Encode(record)
	bmerror.PanicError(err)

	bkc, err := GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "test6"
	bkc.Produce(&topic, recordByteArr)

}

func TestKafkaConsumerWithAvro(t *testing.T) {

	_ = os.Setenv("BM_KAFKA_BROKER", "123.56.179.133:9092")
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", "http://123.56.179.133:8081")
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", "test20190830")
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt")
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem")
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key")
	_ = os.Setenv("BM_KAFKA_SSL_PASS", "pharbers")

	bkc, err := GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topics := []string{"ConnectResponse"}
	bkc.SubscribeTopics(topics, subscribeAvroFunc)

}

func subscribeAvroFunc(a interface{}) {
	fmt.Println("subscribeFunc => ", a)
	var schemaRepositoryUrl = os.Getenv("BM_KAFKA_SCHEMA_REGISTRY_URL")
	decoder := kafkaAvro.NewKafkaAvroDecoder(schemaRepositoryUrl)
	record, err := decoder.Decode(a.([]byte))
	bmerror.PanicError(err)
	fmt.Println("ConnectResponse => ", record.(*avro.GenericRecord))
	fmt.Println("subscribeFunc DONE!")
}

