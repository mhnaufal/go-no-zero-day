package day_35

import (
	"fmt"
	"net/http"
)

func Method() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte("post"))
		case "GET":
			w.Write([]byte("get"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("Server started")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
