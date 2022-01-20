package main

import "fmt"

func Variable() {
	// long way to create variable i64
	var score int = 90

	// short way
	// score := 90
	fmt.Printf("score = %d\n", score)

	// bool
	isPass := false
	fmt.Printf("Is user pass test? %v\n", isPass)
}