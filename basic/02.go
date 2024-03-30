package basic

import "fmt"

func Basic2() {
	fmt.Println("Hello")
	a := [5]int{1, 2, 3, 4, 5}

	b := a[:3]

	fmt.Println(a, b)

	b[0] = 10
	fmt.Println(a, b)
}
