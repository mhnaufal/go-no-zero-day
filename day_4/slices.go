package day_4

import "fmt"

func Slices() {
	/** Slices */
	var languages = []string{"Rust", "C++"}           // Slice
	var languages2 = [3]string{"Rust", "C++", "Go"}   // Array
	var languages3 = [...]string{"Rust", "C++", "Go"} // Array

	fmt.Println("Slice = ", languages)
	fmt.Println("Array = ", languages2)
	fmt.Println("Array = ", languages3)

	fmt.Println("-----------------------------------------------")

	/** Take array languages from index 0 up to index before 2
	* languages2[2:3] means we make a reference to the third element in the array
	* NOT like languages[2] where we create a copy of it
	 */
	var newLanguages = languages2[0:2] //Slice
	fmt.Println("New slice= ", newLanguages)

	fmt.Println("-----------------------------------------------")

	/** Every change to slice will affect the original array */
	original := []string{"m", "o", "o", "n"}
	fmt.Println("Original = ", original)

	slice := original[:1]
	fmt.Println("Slice = ", slice)
	slice[0] = "n"

	fmt.Println("New Slice = ", slice)
	fmt.Println("New Original = ", original)

	fmt.Println("-----------------------------------------------")

	fmt.Println("Length = ", len(slice))
	fmt.Println("Capacity = ", cap(slice))
	fmt.Println("Append = ", append(slice, "s")) //append(slice, element)

	dst := make([]string, 2)
	fmt.Println("Copy = ", copy(dst, slice)) //copy(destination, source)
}
