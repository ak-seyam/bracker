package main

import (
	"fmt"
	"log"

	"github.com/A-Siam/bracker/auth/src/api"
)

const (
	port = 7890
)

func main() {
	app := api.InitApi()
	log.Println("🚀", "Starting the server on port", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
