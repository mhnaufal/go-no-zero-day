package day_5

import "fmt"

func Map() {
	/** Map
	* map[key]value
	 */

	/* 1st style */
	var months map[int]string //until here map is still zero value or nil
	months = map[int]string{
		1: "January",
		2: "February",
	}

	fmt.Println("Months = ", months)

	/* 2nd style */
	days := make(map[string]string)
	days["Monday"] = "Senin"
	days["Tuesday"] = "Selasa"
	days["Wednesday"] = "Rabu"
	fmt.Println("Days = ", days)
	fmt.Println("What day is it today? = ", days["Monday"]) //if we provide with wrong key, it will return null

	fmt.Println("-----------------------------------------------")

	/* Delete key-value on a map */
	delete(days, "Tuesday")

	/* Detect if a key-value is exist in a map */
	value, isExist := days["Tuesday"]
	if isExist {
		fmt.Println("Value = ", value)
	} else {
		fmt.Println("It doens't to be exist =>", isExist)
	}
}
