package main

import (
	"fmt"

	"github.com/A-Siam/bracker/auth/src/api"
	"github.com/A-Siam/bracker/auth/src/common/loggers"
	"github.com/A-Siam/bracker/auth/src/message"
)

const (
	port = 7890
)

func main() {
	app := api.InitApi()
	producer := message.GetProducer()
	defer producer.Close()
	loggers.InfoLogger.Println("Starting the server on port", port)
	loggers.ErrorLogger.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
