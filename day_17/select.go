package day_17

import "fmt"

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func Select() {
	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("Numbers: ", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	for i := 0; i < 1; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		}
	}
}
