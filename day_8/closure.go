package day_8

import "fmt"

func Closure() {
	index := 0

	incrementIndex := func() int {
		index *= 2 //can access the outer variables
		return index
	}

	fmt.Printf("Index value = %v\n", incrementIndex())
	fmt.Printf("Index value = %v\n", incrementIndex())

	fmt.Println("-----------------------------------------------")

	next := sequence()

	fmt.Printf("Sequence = %v\n", next())
	fmt.Printf("Sequence = %v\n", next())
	fmt.Printf("Sequence = %v\n", next())

	newNext := sequence() //restart from 1
	fmt.Printf("New Sequence = %v\n", newNext())

	fmt.Println("-----------------------------------------------")

}

// func functionName returnFunc returnFuncType {}
func sequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
