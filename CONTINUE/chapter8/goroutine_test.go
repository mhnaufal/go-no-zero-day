package chapter8

import (
	"fmt"
	"testing"
	"time"
)

func LogStatus() {
	fmt.Println("[X] status...")
}

func TestGorout(t *testing.T) {
	go LogStatus()
	fmt.Println("END")

	time.Sleep(1 * time.Second)
}
