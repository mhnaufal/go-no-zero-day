package chapter10

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	A := context.Background()
	B := context.WithValue(A, "b", "value B")
	C := context.WithValue(A, "c", "value C")
	D := context.WithValue(C, "d", "value D")

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)

	fmt.Println(A.Value("a")) // <nil> because A doenst have "value A"
	fmt.Println(D.Value("c")) // got from the parent
	fmt.Println(D.Value("d")) // got from herself
}
