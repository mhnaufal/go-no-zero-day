package day_18

import "fmt"
import "strconv"

func Con() {
	var str = "123"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num)
	}

	var num2 = 321
	var str2 = strconv.Itoa(num2)

    fmt.Println(str2)
}
