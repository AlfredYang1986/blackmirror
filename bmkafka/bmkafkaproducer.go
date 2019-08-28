package bmkafka

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"sync"
)

var producer *kafka.Producer
var onceProducer sync.Once

// GetProducerInstance get one KafkaProducerInstance.
func (bkc *Config) GetProducerInstance() (*kafka.Producer, error) {
	onceProducer.Do(func() {
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": bkc.Broker,
			"security.protocol": "SSL",	//默认使用SSL
			"ssl.ca.location": bkc.CaLocation,
			"ssl.certificate.location": bkc.CaSignedLocation,
			"ssl.key.location": bkc.SslKeyLocation,
			"ssl.key.password": bkc.SslPass,
		})

		if err != nil {
			fmt.Printf("Failed to create producer: %s\n", err)
			e = err
		} else {
			fmt.Printf("Created Producer %v\n", p)
			producer = p
			e = nil
		}

	})
	return producer, e
}
// Produce use to produce kafka msg.
func (bkc *Config) Produce(topic *string, value []byte)  {

	p, err := bkc.GetProducerInstance()
	bmerror.PanicError(err)

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
		Value:          value,
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}

	err = p.Produce(&msg, deliveryChan)
	bmerror.PanicError(err)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)

}