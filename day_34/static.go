package day_34

import (
	"fmt"
	"net/http"
)

func Static() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server run on 127.0.0.1:3000")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
