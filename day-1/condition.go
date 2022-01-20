package main

import "fmt"

func Condition() {
	gpa := 3.84

	/** If else statement */
	if gpa > 3.5 {
		fmt.Println("Cumlaude 👨‍🎓")
	} else if gpa > 3. {
		fmt.Println("Awesome work! 🚀")
	} else {
		fmt.Println("Keep it up!")
	}
			
			/** Switch statement */
	switch {
	case gpa > 3.5:
		fmt.Println("Cumlaude 👨‍🎓")
	case gpa > 3.:
		fmt.Println("Awesome work! 🚀")
	default:
		fmt.Println("Keep it up!")
	}
}