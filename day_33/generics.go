package day_33

import "fmt"

// Need Go version 1.18
func sum[V int | float32 | float64](numbers []V) V {
	var total V

	for _,e := range numbers {
		total += e
	}

	return total
}

func Generics() {
	total1 := sum([]int{1, 2, 3, 4})
	fmt.Println("Total: ", total1)
}
