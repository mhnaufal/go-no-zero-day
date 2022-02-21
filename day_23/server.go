package day_23

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

/* Function callback = Routing & controller */
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Server() {
	/* Similiar to route callback in NodeJS */
	/* Take a route and a function callback */
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hallo!")
	// })
	fmt.Println(os.Getwd())
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"name":    "John Wick",
			"message": "Have a good day!",
		}

		var t, err = template.ParseFiles("C:/Users/Ryzen/Desktop/Projects4/Go/no-zero-day/day_23/template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)

	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
