package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {

	http.HandleFunc("/api", HelloHandler)
	http.ListenAndServe(":7890", nil)

}
