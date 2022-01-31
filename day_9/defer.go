package day_9

import "fmt"

func Defer() {
	/** Defer */
	defer second() // called after func Defer() has finished
	first()

	/** Panic */
	// panic("PANIC")
	str := recover() //not gonna called because of panic
	fmt.Println(str)

	/** Recover */
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func first() {
	fmt.Println("Called 1st")
}
func second() {
	fmt.Println("Called 2nd")
}
