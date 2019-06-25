package bmkafka

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var consumerMap map[string]*kafka.Consumer

func init() {
	consumerMap = nil
}

// GetConsumerInstanceByTopics get one KafkaConsumerInstance by topics.
func (bkc *BmKafkaConfig) GetConsumerInstanceByTopics(topics []string) (*kafka.Consumer, error) {

	topicsStr := strings.Join(topics, "##")

	if consumerMap == nil {
		consumerMap = make(map[string]*kafka.Consumer, 0)
	}

	if consumerMap[topicsStr] == nil {
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": bkc.Broker,
			// Avoid connecting to IPv6 brokers:
			// This is needed for the ErrAllBrokersDown show-case below
			// when using localhost brokers on OSX, since the OSX resolver
			// will return the IPv6 addresses first.
			// You typically don't need to specify this configuration property.
			"broker.address.family": "v4",
			"group.id":              bkc.Group + "_" + topicsStr,
			"session.timeout.ms":    6000,
			//"auto.offset.reset":        "earliest",
			"auto.offset.reset":        "latest",
			"security.protocol":        "SSL", //默认使用SSL
			"ssl.ca.location":          bkc.CaLocation,
			"ssl.certificate.location": bkc.CaSignedLocation,
			"ssl.key.location":         bkc.SslKeyLocation,
			"ssl.key.password":         bkc.Pass,
		})

		if err != nil {
			fmt.Printf("Failed to create consumer: %s\n", err)
			e = err
		} else {
			fmt.Printf("Created Consumer %v\n", c)
			consumerMap[topicsStr] = c
			e = nil
		}


		return c, e

	} else {
		return consumerMap[topicsStr], e
	}

}

// SubscribeTopics subscribe some topics from args or config.
func (bkc *BmKafkaConfig) SubscribeTopics(topics []string, subscribeFunc func(interface{})) {
	if len(bkc.Topics) == 0 {
		panic("no Topics in config")
	}
	var topicsTmp []string
	if len(topics) == 0 {
		topicsTmp = bkc.Topics
	} else {
		topicsTmp = topics
	}

	c, err := bkc.GetConsumerInstanceByTopics(topicsTmp)
	//defer c.Close()
	bmerror.PanicError(err)
	err = c.SubscribeTopics(topicsTmp, nil)
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
				continue
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()

}

// SubscribeTopicsOnce subscribe some topics from args or config.
// Only once!
func (bkc *BmKafkaConfig) SubscribeTopicsOnce(topics []string, duration time.Duration, subscribeFunc func(interface{})) {
	if len(bkc.Topics) == 0 {
		panic("no Topics in config")
	}
	var topicsTmp []string
	if len(topics) == 0 {
		topicsTmp = bkc.Topics
	} else {
		topicsTmp = topics
	}

	c, err := bkc.GetConsumerInstanceByTopics(topicsTmp)
	//defer c.Close()
	bmerror.PanicError(err)
	err = c.SubscribeTopics(topicsTmp, nil)
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

	fmt.Printf("Closing consumer\n")
	c.Close()

}
