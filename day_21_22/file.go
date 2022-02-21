package day_21_22

import (
	"fmt"
	"os"
)

var path = "C:/Users/Ryzen/Desktop/Projects4/Go/no-zero-day/day_21/file.txt"

func File() {
	// Creating a new file
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)

		if err != nil {
			return
		}

		defer file.Close()
	}

	fmt.Println("File => ", path)

	/* Day 22 */
	// Writing into file
	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)

	if err2 != nil {
		return
	}

	defer file.Close()

	_, err2 = file.WriteString("Hello all!\n") //write into the file

	if err2 != nil {
		return
	}

	err2 = file.Sync() //save the changes

	if err2 != nil {
		return
	}

	fmt.Println("File updated => ", file)
}
