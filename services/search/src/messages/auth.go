package messages

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/A-Siam/bracker/search/src/common/loggers"
	"github.com/A-Siam/bracker/search/src/dto"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type AuthCallback func(userDto dto.UserDto, ok chan bool)

func ListenOnAuthTopic(callbacks []AuthCallback, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	loggers.InfoLogger.Println("start listening on auth topic")

	kafkaClient := GetConsumer()
	kafkaClient.SubscribeTopics([]string{os.Getenv("AUTH_TOPIC")}, nil)
	for {
		msg, err := kafkaClient.ReadMessage(time.Second)
		if err != nil && !err.(kafka.Error).IsTimeout() {
			loggers.ErrorLogger.Println("inside receiving message", err.Error())
		} else {
			if msg == nil {
				continue
			}
			for _, cb := range callbacks {
				okChan := make(chan bool)
				var userDto = dto.UserDto{}
				err := json.Unmarshal(msg.Value, &userDto)
				if err != nil {
					return
				}
				go cb(userDto, okChan)
				<-okChan
			}
		}
	}
}
