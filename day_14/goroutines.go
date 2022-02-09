package day_14

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func threadLike(counter int, message string) {
	for i := 1; i <= counter; i++ {
		fmt.Printf("[%v] %v\n", i, message)
		amt := time.Duration(rand.Intn(800))
		time.Sleep(time.Millisecond * amt)
	}
}

func Goroutines() {
	runtime.GOMAXPROCS(2) //how many core(s) will be used

	threadLike(5, "Before")
	go threadLike(5, "Go") //going to run after the 'Outside'
	threadLike(5, "After")

	var input string
	fmt.Scanln(&input)
}
