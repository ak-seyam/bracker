package main

import (
	"sync"

	"github.com/A-Siam/bracker/search/src/callbacks"
	"github.com/A-Siam/bracker/search/src/messages"
)

func main() {
	var wg sync.WaitGroup
	messages.ListenOnAuthTopic([]messages.AuthCallback{
		callbacks.OnUserCreated,
	}, &wg)
}
