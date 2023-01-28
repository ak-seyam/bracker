package messages

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
	"bootstrap.servers": os.Getenv("BROKER_SERVERS"),
	"group.id":          os.Getenv("MESSAGE_GROUP"),
	"auto.offset.reset": "earliest",
})

func GetConsumer() *kafka.Consumer {
	if err != nil {
		panic(err)
	}
	return consumer
}
