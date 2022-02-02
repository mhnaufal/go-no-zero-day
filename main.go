package main

import (
	"fmt"
	"no-zero-day/day_1"
	"no-zero-day/day_2"
	"no-zero-day/day_3"
	"no-zero-day/day_4"
	"no-zero-day/day_5"
	"no-zero-day/day_6"
	"no-zero-day/day_7"
	"no-zero-day/day_8"
	"no-zero-day/day_9"
	"no-zero-day/day_10"
	"no-zero-day/day_11"
)

func main() {
	fmt.Println("Day One")
	day_1.Variable()
	day_1.Constant()
	day_1.Condition()

	fmt.Println("\nDay Two")
	day_2.Loop()

	fmt.Println("\nDay Three")
	day_3.Array()

	fmt.Println("\nDay Four")
	day_4.Slices()

	fmt.Println("\nDay Five")
	day_5.Map()

	fmt.Println("\nDay Six")
	day_6.Function()

	fmt.Println("\nDay Seven")
	day_7.Function2()

	fmt.Println("\nDay Eight")
	day_8.Closure()

	fmt.Println("\nDay Nine")
	day_9.Defer()

	fmt.Println("\nDay Ten")
	day_10.Pointer()
	
	fmt.Println("\nDay Eleven")
	day_11.Struct()
}
