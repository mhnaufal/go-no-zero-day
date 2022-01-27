package day_7

import "fmt"

func Function2() {
	/** Variadic function */
	fmt.Println("Variadic function")
	gpa(5, 1, 2, 3, 4)
}

/** func functionName(paramA typeA, paramB ...typeB) returnType */
func gpa(div int, marks ...int) int {
	total := 0

	for _, mark := range marks {
		total += mark
	}

	total = total / div

	fmt.Println("Total", total)
	return total
}
