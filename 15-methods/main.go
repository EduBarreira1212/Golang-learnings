package main

import "fmt"

type student struct {
	name   string
	age    uint8
	course string
}

func (student student) register() {
	fmt.Printf("%s registred sucessfully!\n", student.name)
}

func (student *student) birthday() {
	student.age++
}

func (student *student) changeCourse(newCourse string) {
	student.course = newCourse
}

func main() {
	student1 := student{"Edu", 21, "Laws"}
	fmt.Println(student1)
	student1.register()
	student1.birthday()
	fmt.Println(student1)
	student1.changeCourse("Medicine")
	fmt.Println(student1)
}
