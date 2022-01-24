package day_3

import "fmt"

func Array() {
	/** Array */
	var empty [3]int
	fmt.Println("Empty array: ", empty)

	var languages [4]string
	languages[0] = "en"
	languages[1] = "id"
	languages[2] = "fr"
	languages[3] = "en" // duplicate value allowed

	fmt.Println("Languages: "+languages[0], languages[1], languages[2], languages[3])

	fmt.Println("-----------------------------------------------")

	var programming_language = [3]string{"Rust", "Go", "C++"}
	// var programming_language = [...]string{"Rust", "Go", "C++"}
	fmt.Println("Programming Languages: "+programming_language[0], programming_language[1], programming_language[2])
	fmt.Println("Length of the array: ", len(programming_language))

	var frameworks = make([]string, len(programming_language))
	frameworks[0] = "Laravel"
	frameworks[1] = "Rocket"
	fmt.Println(frameworks)

	fmt.Println("-----------------------------------------------")

	/** Array 'for' loop will result index and element */
	for index, framework := range frameworks {
		fmt.Printf("%v. %v\n", index, framework)
	}

	/** We can skip the array element */
	for index, _ := range frameworks {
		fmt.Printf("Index = %v\n", index)
	}

	/** We can also skip the index */
	for _, framework := range frameworks {
		fmt.Printf("%v\n", framework)
	}
}
