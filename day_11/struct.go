package day_11

import "fmt"

type Student struct {
	name  string
	grade int
}

func Struct() {
	var student1 Student
	student1.name = "John Wick"
	student1.grade = 99
	fmt.Println("Init => ", student1.name)
	fmt.Println("Init => ", student1.grade)

	fmt.Println("-----------------------------------------------")

	// Anonymous struct
	var student2 = Student{"Johny Deep", 32}
	fmt.Println(student2)

	var student3 = struct {
		Student
		class string
	}{
		Student: Student{"Sky", 33},
		class:   "A",
	}

	fmt.Println(student3)
	fmt.Println(student3.Student.name)
	fmt.Println(student3.class)

	fmt.Println("-----------------------------------------------")

	// Pointer struct
	var student4 *Student = &student2
	fmt.Println("Pointer student = ", student4)
}
