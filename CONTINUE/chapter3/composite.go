package chapter3

import "fmt"

func Composite() {
	// COMPOSITE TYPE
	var arr_1 [3]int // zero value array
	fmt.Println("array: ", arr_1)

	var arr_2 = [3]int{10, 10, 10}
	fmt.Println("array: ", arr_2)

	var arr_3 = [...]int{1, 2, 3}
	fmt.Println("array: ", arr_3)

	fmt.Println("+------------------+")

}
