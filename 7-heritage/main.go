package main

import "fmt"

type person struct {
	name string
	age  int8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	person1 := person{"Eduardo", 21}

	student1 := student{person1, "Law", "Harvard"}

	fmt.Println(student1)
	fmt.Println(student1.age)
}
