package day_20

import "fmt"
import "log"
import "os/exec"

func Exec() {
	var command1, err = exec.Command("pwd").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("=> pwd\n%s\n", string(command1))
}
