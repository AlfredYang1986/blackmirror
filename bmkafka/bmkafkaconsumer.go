package bmkafka

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var consumer *kafka.Consumer
var onceConsumer sync.Once

func (bkc *bmKafkaConfig) GetConsumerInstance() (*kafka.Consumer, error) {
	onceConsumer.Do(func() {
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": bkc.Broker,
			// Avoid connecting to IPv6 brokers:
			// This is needed for the ErrAllBrokersDown show-case below
			// when using localhost brokers on OSX, since the OSX resolver
			// will return the IPv6 addresses first.
			// You typically don't need to specify this configuration property.
			"broker.address.family": "v4",
			"group.id":              bkc.Group,
			"session.timeout.ms":    6000,
			"auto.offset.reset":     "earliest"})

		if err != nil {
			fmt.Printf("Failed to create consumer: %s\n", err)
			e = err
		} else {
			fmt.Printf("Created Consumer %v\n", c)
			consumer = c
			e = nil
		}

		//err = c.SubscribeTopics(bkc.Topics, nil)

	})
	return consumer, e
}

func (bkc *bmKafkaConfig) SubscribeTopics(topics []string, subscribeFunc func(interface{})) {
	if len(bkc.Topics) == 0 {
		panic("no Topics in config")
	}
	c, err := bkc.GetConsumerInstance()
	bmerror.PanicError(err)
	if len(topics) == 0 {
		err = c.SubscribeTopics(bkc.Topics, nil)
	} else {
		err = c.SubscribeTopics(topics, nil)
	}
	bmerror.PanicError(err)

	run := true
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				subscribeFunc(e.Value)
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

}

func (bkc *bmKafkaConfig) SubscribeTopicsOnce(topics []string, duration time.Duration, subscribeFunc func(interface{})) {
	if len(bkc.Topics) == 0 {
		panic("no Topics in config")
	}
	c, err := bkc.GetConsumerInstance()
	bmerror.PanicError(err)
	if len(topics) == 0 {
		err = c.SubscribeTopics(bkc.Topics, nil)
	} else {
		err = c.SubscribeTopics(topics, nil)
	}
	bmerror.PanicError(err)

	run := true
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	timeout := time.After(duration)

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		case <-timeout:
			fmt.Println("SubscribeTopicsOnce timeout ", duration)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				subscribeFunc(e.Value)
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
				run = false
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

}
