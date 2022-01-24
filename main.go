package main

import (
	"fmt"
	"no-zero-day/day_1"
	"no-zero-day/day_2"
	"no-zero-day/day_3"
)

func main() {
	fmt.Println("Day One")
	day_1.Variable()
	day_1.Constant()
	day_1.Condition()

	fmt.Println("Day Two")
	day_2.Loop()
	
	fmt.Println("Day Three")
	day_3.Array()
}
