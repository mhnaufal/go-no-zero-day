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

	fmt.Println("+------------------+")
}
