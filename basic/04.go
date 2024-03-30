package basic

import (
	"fmt"
)

type Animal interface {
	run()
	eat()
	setAge(age int)
}

type Chicken struct {
	Name string
	Age  int
}

// setAge implements Animal.
func (c *Chicken) setAge(age int) {
	c.Age = age
}

func (c *Chicken) run() {
	fmt.Println(&c)
	fmt.Println("GO.....")
}

func (c *Chicken) eat() {
	fmt.Println(&c)
	fmt.Println("EAT.....")
}
func (c *Chicken) printAge() {
	fmt.Println(&c)
	fmt.Println("Age: ", c.Age)
}

func Basic4() {
	var a Animal

	b := Chicken{Name: "abc"}

	a = &b

	a.eat()
	a.run()
	a.setAge(10)
	fmt.Println(a)

	fmt.Println("==========")

	b.eat()
	b.run()
	fmt.Println(b)
	b.setAge(15)

	b.printAge()

}
