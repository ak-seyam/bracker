package main

import (
	"fmt"
	"sync"

	"github.com/A-Siam/bracker/search/src/api"
	"github.com/A-Siam/bracker/search/src/callbacks"
	"github.com/A-Siam/bracker/search/src/common/loggers"
	"github.com/A-Siam/bracker/search/src/messages"
)

const (
	port   = 7893
	n_apis = 1
)

func main() {
	loggers.InfoLogger.Println("start searching service")
	var wg sync.WaitGroup
	authCallbacks := []messages.AuthCallback{
		callbacks.OnUserCreated,
	}
	wg.Add(n_apis + len(authCallbacks))
	go func() {
		messages.ListenOnAuthTopic(authCallbacks, &wg)
	}()
	go func() {
		defer wg.Done()
		app := api.InitApi()
		loggers.InfoLogger.Printf("start listening on port :%d", port)
		app.Listen(fmt.Sprintf(":%d", port))
	}()
	wg.Wait()
	loggers.InfoLogger.Println("search system exited closed")
}
