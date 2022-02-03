package day_12

import "fmt"

type Student struct {
	name  string
	grade int
}

// func structName methodName() {}
func (s Student) getStudentData() {
	fmt.Printf("Name = '%v' Grade = '%v'\n", s.name, s.grade)
}

func (s *Student) changeName(newName string) {
	s.name = newName
}

func Method() {
	var student1 = Student{"Johny Deep", 99}
	student1.getStudentData()

	student1.changeName("James Deep")
	student1.getStudentData()
}
