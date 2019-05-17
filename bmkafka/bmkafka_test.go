package bmkafka

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	"github.com/hashicorp/go-uuid"
	"io/ioutil"
	"os"
	"testing"
)

func TestKafkaProducer(t *testing.T) {

	os.Setenv("BM_KAFKA_CONF_HOME", "../resource/kafkaconfig.json")

	bkc, err := GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "test"
	bkc.Produce(&topic, []byte("LaoDeng => TestKafkaProducer"))

}

func TestKafkaConsumer(t *testing.T) {

	os.Setenv("BM_KAFKA_CONF_HOME", "../resource/kafkaconfig.json")

	bkc, err := GetConfigInstance()
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

func TestKafkaProduceToXmpp(t *testing.T) {

	os.Setenv("BM_KAFKA_CONF_HOME", "../resource/kafkaconfig.json")

	var schemaRepositoryUrl = "http://59.110.31.50:8081"
	var rawMetricsSchema = `{"namespace": "net.elodina.kafka.metrics","type": "record","name": "XmppCmd","fields": [{"name": "id", "type": "string"},{"name": "reportUser",  "type": "string" },{"name": "msg",  "type": "string" }]}`

	encoder := kafkaAvro.NewKafkaAvroEncoder(schemaRepositoryUrl)
	schema, err := avro.ParseSchema(rawMetricsSchema)
	bmerror.PanicError(err)
	record := avro.NewGenericRecord(schema)
	tmpUUID, err := uuid.GenerateUUID()
	bmerror.PanicError(err)
	record.Set("id", tmpUUID)
	record.Set("reportUser", "test@max.logic")
	record.Set("msg", "hello!!!")
	recordByteArr, err := encoder.Encode(record)

	bkc, err := GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "xmpp-topic"
	bkc.Produce(&topic, recordByteArr)

}

func TestKafkaProduceToOss(t *testing.T) {

	os.Setenv("BM_KAFKA_CONF_HOME", "../resource/kafkaconfig.json")

	var schemaRepositoryUrl = "http://59.110.31.50:8081"
	var rawMetricsSchema = `{"namespace": "net.elodina.kafka.metrics","type": "record","name": "OssCmd","fields": [{"name": "id", "type": "string"},{"name": "bucketName",  "type": "string" },{"name": "objectKey",  "type": "string" },{"name": "objectValue",  "type": "bytes" }]}`

	encoder := kafkaAvro.NewKafkaAvroEncoder(schemaRepositoryUrl)
	schema, err := avro.ParseSchema(rawMetricsSchema)
	bmerror.PanicError(err)
	record := avro.NewGenericRecord(schema)
	tmpUUID, err := uuid.GenerateUUID()
	bmerror.PanicError(err)
	fmt.Println(tmpUUID)
	record.Set("id", tmpUUID)
	record.Set("bucketName", "pharbers-resources")
	record.Set("objectKey", tmpUUID)

	localDir := "/home/jeorch/work/test/temp/test.jpeg"
	f, err := os.Open(localDir)
	bmerror.PanicError(err)
	defer f.Close()

	objectValue, err := ioutil.ReadAll(f)
	bmerror.PanicError(err)
	record.Set("objectValue", objectValue)

	recordByteArr, err := encoder.Encode(record)

	bkc, err := GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "oss-topic"
	bkc.Produce(&topic, recordByteArr)

}

