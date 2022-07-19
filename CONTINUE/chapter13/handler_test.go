package chapter13

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		fmt.Fprintln(writer, "Hello Mom!")
	}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Server failed to start")
	}
}
