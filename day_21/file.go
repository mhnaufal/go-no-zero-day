package day_21

import (
	"fmt"
	"os"
)

var path = "C:/Users/Ryzen/Desktop/Projects4/Go/no-zero-day/day_21/file.txt"

func File() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)

		if err == nil {
			return
		}

		defer file.Close()
	}

	fmt.Println("File => ", path)
}
