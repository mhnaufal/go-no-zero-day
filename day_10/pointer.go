package day_10

import "fmt"

func Pointer() {
	/** Pointer */
	x := 5
	fmt.Printf("Value before %v\n", x)
	// &x -> used to get the 'address' of a variable
	fmt.Printf("Value after %v", x)
	
	// Another way to create a pointer
	y := new(int) //same as &y
	fmt.Printf("Value before %v\n", *y)
	changeValue(y)
	fmt.Printf("Value after %v\n", *y)
}

// create a variable of type *int with the name of valuePointer
func changeValue(valuePointer *int) {
	// *valuePointer -> used to get the 'value' from a pointer
	*valuePointer = 99
}
