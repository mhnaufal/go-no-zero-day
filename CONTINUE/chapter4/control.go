package chapter4

import (
	"fmt"
	"math/rand"
)

func Control() {
	// COMPOSITE TYPE
	fmt.Println("CONTROL STRUCTURE")

	// SHADOWING
	fmt.Println("[1] SHADOWING")
	x := 10
	if x > 5 {
		fmt.Println("shadow [x]: ", x)
		x := 13 // var x int = 13
		fmt.Println("shadow [x]: ", x)
	}
	fmt.Println("shadow [x]: ", x)

	println()

	// IF
	fmt.Println("[2] IF")

	if n := rand.Intn(10); n == 3 || n == 10 { // only short declarations allowed
		fmt.Println("if: WOW")
	} else {
		fmt.Println("if: MEH")
	}

	println()

	// SWITCH
	fmt.Println("[3] SWITCH")

	switch name := "Darkseid"; name {
	case "Joker":
		fmt.Println("switch: I am Joker")
		fmt.Println("switch: I am Joker")
	case "Darkseid":
		fmt.Println("switch: I am Darkseid")
		fmt.Println("switch: I am Darkseid")
	default:
		fmt.Println("switch: I am not villain")
		fmt.Println("switch: I am not villain")
	}

	println()

	// switch without specifying condition in the switch statement
	var name string = "Batman"
	switch {
	case len(name) < 2:
		fmt.Println("switch: WHAT???")
	case len(name) >= 3 && len(name) <= 10:
		fmt.Println("switch: OK")
	default:
		fmt.Println("switch: interesting")
	}

	println()

	// FOR FOUR LOOP
	fmt.Println("[4] FOR LOOP")

	i := 0
	for i <= 5 {
		fmt.Printf("%v ", i)
		i++
	}

	println()

	for i := 5; i >= 0; i-- {
		fmt.Printf("%v ", i)
	}

	println()

	villains := []string{"Joker", "Deadshoot", "Darkseid"}
	for _, villain := range villains {
		fmt.Printf("%v ", villain)
	}

	println()

	fmt.Println("+------------------+")
}
