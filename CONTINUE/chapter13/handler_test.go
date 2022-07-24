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

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello mom")
	})

	mux.HandleFunc("/healthcheck", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Healthy")
	})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Server failed to start")
	}
}

func TestRequest(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Body)
		fmt.Fprint(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	})
}