package helper

import "fmt"

type Movement interface {
	move()
}

type Animal interface {
	speak()
	// eat()
}

type NextAnimal interface {
	Animal
	Movement
}

type Dog struct {
	Name string
}

func (d Dog) speak() {
	fmt.Println("Gau Gau....!")
}

func (d Dog) move() {
	fmt.Println("Run Run....!")
}

func goOut(i interface{}) {
	fmt.Println(i)
}

func MyInterface() {
	// var animal Animal

	// animal = Dog{Name: "Dog"}

	// animal.speak()

	dog := Dog{}

	var m Movement = dog
	m.move()
	var a Animal = dog
	a.speak()

	var na NextAnimal = dog
	na.move()
	na.speak()
	fmt.Println("==========")
	goOut(10)
	goOut("Hello")
}
