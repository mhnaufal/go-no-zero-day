package day_15

import "fmt"

func Channel() {
	// create a channel called message of type string
	var message = make(chan string)

	var greeting = func(whom string) {
		var data = fmt.Sprintf("Hello %v\n", whom)
		message <- data
	}

	go greeting("John Wick")
	go greeting("Grindelwald")
	go greeting("Joker")

	var message1 = <-message
	fmt.Printf("# %v", message1)

	var message2 = <-message
	fmt.Printf("# %v", message2)

	var message3 = <-message
	fmt.Printf("# %v", message3)
}