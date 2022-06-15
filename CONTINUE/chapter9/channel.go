package chapter9

import (
	"fmt"
	"time"
)

func Channel() {
	// CHANNEL
	fmt.Println("[1] CHANNEL")

	// create a channel
	channelData := make(chan string, 5) // create a channel with a buffer of capacity 5
	defer close(channelData)

	// send data to a channel
	// channelData <- "Batman"
	// fmt.Println("send data to a channel: ", channelData)

	// receive data from a channel
	// name := <-channelData
	// fmt.Println("receive data from channel: ", name)

	// close channel
	// close(channelData)

	go func() {
		time.Sleep(2 * time.Second)
		channelData <- "Joker"              // will block until there is goroutine who receive this
		channelData <- "Joker"              // will block until there is goroutine who receive this
		channelData <- "Joker"              // will block until there is goroutine who receive this
		fmt.Println("send data to channel") // will come after channel finish send the data
	}()

	fmt.Println("buffer length: ", len(channelData))
	fmt.Println("buffer capacity: ", cap(channelData))

	name := <-channelData
	fmt.Println("receive data from a channel: ", name)

	go ResponseChannel("Batman", channelData)
	name = <-channelData

	time.Sleep(5 * time.Second)

	println()

}

func ResponseChannel(villain string, channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- villain
	fmt.Println("send data to channel")
}
