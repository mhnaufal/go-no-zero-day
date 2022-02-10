package day_16

import "fmt"

func Buffer() {
	messages := make(chan int, 2) //total there is 3 buffer

	// Run asynchronously
	go func() {
		for {
			i := <-messages
			fmt.Println("Receive data ", i)
		}
	}()

	// Run synchronously after previous buffer has been received
	for i := 0; i < 5; i++ {
		fmt.Println("Send data ", i)
		messages <- i
	}

	/** Buffers are going to send asyn
	 * then, when buffer already had enough space
	 * the buffer will be received syn until there are some empty space (at least 1 empty space)
	 * then the async send will run again
	 */
}
