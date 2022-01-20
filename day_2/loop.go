package day_2

import "fmt"

func Loop() {
	/** Normal 'for' loop */
	for i := 0; i < 10; i++ {
		fmt.Printf("The %v-loop\n", i)
	}

	/** Condition only 'for' loop */
	var i = 5
	for i >= 0 {
		fmt.Println("The 'condition only' loop")
		i--
	}

	/** No argument 'for' loop */
	var j = 3
	for {
		fmt.Println("The argumentless loop")
		j++

		if j == 5 {
			break
		}
	}

	fmt.Println("|------------------------------------------------|")

	/** Range 'for' loop */
	var languages = [3]string{"C++", "Java", "Rust"}
	var otherLanguages = [3]string{
		"Go",
		"Typescript",
		"Kotlin",
	}

	for i, language := range languages {
		fmt.Printf("%d. %s\n", i, language)
	}

	fmt.Println("|------------------------------------------------|")

	// we don't want to use the index from the loop
	for _, language := range otherLanguages {
		// i will have value of '-1'
		fmt.Printf("%d. %s\n", i, language)
	}

	/** Go doesn't have 'while' 'do while' loop style */
}
