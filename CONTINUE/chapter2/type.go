package chapter2

import "fmt"

func Type() {
	// NUMBER
	num := 3_13
	fmt.Println("integer: ", num)

	floatNum := 6.2e3
	fmt.Println("float: ", floatNum)

	complexNum := complex(2.3, 3.000)
	fmt.Println("complex: ", complexNum)
	fmt.Println("complex: ", real(complexNum))

	fmt.Println("+------------------+")

	// BOOLEAN
	var pass bool = true
	if pass {
		fmt.Println("is pass? YES")
	} else {
		fmt.Println("NOPE")
	}

	fmt.Println("+------------------+")

	// STRING
	var name string = "MRX"
	fmt.Println("the name: ", name)
	fmt.Println("length: ", len(name))

	// will return the byte representation of the char, not the actual char
	// panic if out of bound
	fmt.Println("get char: ", name[0])
	fmt.Println("get char: ", string(name[0])) // change byte to string

	fmt.Println("+------------------+")

	// TYPE CONVERSION
	var x int = 13
	var y float32 = 2.2
	var z float32 = float32(x) - y
	fmt.Println("convert z: ", z)

	fmt.Println("+------------------+")

	// TYPE DECLARATIONS
	type PassStatus bool

	var isPass PassStatus = true
	fmt.Println("type declarations: ", isPass)

	fmt.Println("+------------------+")
}
