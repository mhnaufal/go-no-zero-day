package chapter9

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i <= 11; i++ {
			channel <- "[" + strconv.Itoa(i) + "] loops"
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Accepts : ", data)
	}

	fmt.Println("Finish")
}
