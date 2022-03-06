package day_32

import (
	"fmt"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {
	// decrement the WaitGroup by one
	defer wg.Done()

	// get execute after line 10 returns
	fmt.Println(message)
}

func SyncWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("Data %d", i)

		// add 1 goroutine in each loop
		wg.Add(1)

		// execute doPrint() as a goroutine
		go doPrint(&wg, data)
	}

	// block the program so that it wont execute the next line
	// before all of 5 goroutine had been done
	wg.Wait()
}

// The result will appear at the same time
// because it's already sync
