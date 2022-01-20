package day_1

import "fmt"

func Variable() {
	/** Integer */
	// explicitly declare a variable of i64 or i32
	var score int = 90

	// implicitly declare a variable with type auto detected by Go
	// score := 90
	fmt.Printf("score = %d\n\n", score)

	/** Bool */
	isPass := false
	fmt.Printf("Is user pass test? %v\n\n", isPass)

	/** Floating point */
	var liter float64 = 5.5
	fmt.Printf("Remaining liters = %v\n\n", liter)

	/** String */
	fullName := "My full name"
	fmt.Printf("Here it is my full name = %v\n\n", fullName)

	// get the length of the string
	nameLength := len(fullName)
	fmt.Printf("My full name length is approximately = %v\n\n", nameLength)

	// get the nth character from a string
	// the nth char will return a byte representation of the nth char
	// if the nth char is beyond the length of the string, there will be panic at runtime.
	var nChar uint8 = fullName[0]
	// var nChar byte = fullName[0]		// panic: runtime error: index out of range
	fmt.Printf("The first character from my full name is = %v\n\n", nChar)

	theByte := string(nChar)
	fmt.Printf("Byte %v is equal to string '%v'\n\n", nChar, theByte)

	/** Multiple variable declaration */
	var (
		x int64   = 10
		y float64 = 3.
		z int64   = 2
	)

	fmt.Printf("z is %d, y is %.2f, z is %d\n\n", x, y, z)
}
