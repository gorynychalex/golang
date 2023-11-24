package main

import "fmt"

type Person struct {
	name string
}

type Student struct {
	Person
	group string
}

func (p *Person) Talk() {
	fmt.Println("My name is ", p.name)
}

func main() {

	person := new(Person)
	person.name = "Michael"
	person.Talk()

	stud1 := new(Student)
	stud1.Person.name = "Alex"
	stud1.Talk()
}