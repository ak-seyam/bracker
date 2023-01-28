package message

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/A-Siam/bracker/auth/src/common/loggers"
	"github.com/A-Siam/bracker/auth/src/dto"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ProduceNewUserMessage(user dto.UserDto) error {
	loggers.InfoLogger.Println("🛸 Sending user with id = ", user.ID, " creation message")
	producer := GetProducer()
	authTopic := os.Getenv("AUTH_TOPIC")
	if len(authTopic) == 0 {
		return errors.New("invalid auth topic")
	}
	message, err := json.Marshal(user)
	if err != nil {
		return err
	}
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &authTopic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(user.ID),
		Value: message,
	}, nil)
	return nil
}
