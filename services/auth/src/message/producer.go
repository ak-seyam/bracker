package message

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var producer, err = kafka.NewProducer(&kafka.ConfigMap{
	"bootstrap.servers": os.Getenv("BROKER_SERVERS"),
})

func GetProducer() *kafka.Producer {
	if err != nil {
		panic(err)
	}
	return producer
}
