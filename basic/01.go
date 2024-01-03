package basic

import (
	"fmt"
	"strconv"
)

func variable() {
	var number int
	number = 10

	fmt.Println(number)

	var email = "abc@gmail.com"
	fmt.Println(email)

	var a, b = 10, "11"

	// a = 1
	// b = 2

	fmt.Println(a, b)

	var (
		name = "kuro"
		age  int
	)

	name = "kuro"
	age = 27

	fmt.Println(name, age)

	language := "golang"
	fmt.Println(language)
}

func dataType() {
	var myName string = "Tuấn Trần"
	runes := []rune(myName)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	s := "123"

	// string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}

	fmt.Println(s, i)

	const NAME = "kuro"
}

func conditionAndLoop() {
	if a := addTwoNumber(1, 2); a <= 100 {
		fmt.Println("ok")
	}
	a := 10
	switch a {
	case 11:
		fmt.Println("1")
		fallthrough
	case 10:
		fmt.Println("2")
	}
}

func addTwoNumber(a int, b int) int {
	return a + b
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func myFunc(name string) {
	a := fmt.Sprintf("%s", name)
	fmt.Println(a)

	w, h, isSquare := squareInfo(100, 100)

	fmt.Println(
		w, h, isSquare,
	)
}

func squareInfo(w, h int) (width int, hight int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}

func myArray() {
	var m1 [4]int
	m1[0] = 1
	fmt.Println(m1)

	m2 := []int{1, 2, 3, 4}
	fmt.Println(m2)

	m3 := [3]int{1}
	fmt.Println(m3)

	fmt.Println("==========")
	fmt.Println(len(m3))

	m4 := [...]int{1, 2, 3}
	m4[1] = 10
	fmt.Println(m4)

	m5 := [...]int{1, 2, 3}
	m6 := m5
	m6[1] = 10

	fmt.Println("m5", m5)
	fmt.Println("m6", m6)

	m7 := m2

	m7[0] = 10

	fmt.Println("m2", m2)
	fmt.Println("m7", m7)
}

func myArray2() {
	m1 := []int{1, 2, 3, 4}

	for idx /*, value */ := range m1 {
		// fmt.Printf("Index: %d, value: %d\n", idx, value)
		fmt.Printf("Index: %d\n", idx)
	}

	// Declare an integer array of five elements.
	// Initialize index 1 and 2 with specific values.
	// The rest of the elements contain their zero value.
	array := [5]int{1: 10, 2: 20, 4: -1}

	fmt.Println("array", array)
}

func myArray3() {

	array := [5]int{0, 1, 2, 3, 4}
	fmt.Println("===========")
	fmt.Println(array)

	fmt.Println("===========1")
	handleArray(array)
	fmt.Println(array)

	fmt.Println("===========2")
	handleArray2(&array)
	fmt.Println(array)
}

func handleArray(array [5]int) {
	for i, v := range array {
		fmt.Println(i, v)
		array[i] *= 2
	}
	fmt.Println("==========handleArray==========")
	fmt.Println(array)
}

func handleArray2(array *[5]int) {
	for i := 0; i < len(*array); i++ {
		array[i] *= 2
	}
	fmt.Println("==========handleArray2==========")
	fmt.Println(array)
}

func mySlice() {
	var sl []int
	fmt.Println(len(sl), sl)

	var sl1 = []int{1, 2, 3, 4}
	fmt.Println(sl1)

	var m1 = [4]int{1, 2, 3, 4}

	sl2 := m1[1:3]
	fmt.Println(sl2)

	sl3 := m1[:]
	fmt.Println(sl3)

	var sl4 = []int{0, 1, 2, 3, 4}
	sl5 := sl4

	fmt.Println(sl5)

	sl6 := []int{0, 1, 2, 3}

	sl7 := sl6
	sl7[0] = 20
	fmt.Println("sl6: ", sl6)
	fmt.Println("sl7: ", sl7)
}

func mySlice2() {
	countries := [...]string{"VN", "USA", "UK", "CANADA", "China"}

	sl := countries[2:3]
	fmt.Println(sl)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))

	var sl1 []int
	fmt.Println(cap(sl1))

	fmt.Println("=========== 1")

	sl2 := make([]int, 2, 5)

	fmt.Println(sl2, len(sl2), cap(sl2))

	fmt.Println("=========== 2")

	sl3 := make([]int, 2)

	sl3 = append(sl3, 100)
	fmt.Println(sl3, len(sl3), cap(sl3))

	fmt.Println("=========== 3")

	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 2)

	copy(dest, src)
	fmt.Println(dest, len(dest), cap(dest))

	src = append(src[:1], src[2:]...)
	fmt.Println("src: ", src)
}

func mySlice3() {
	// sl4 := make([]int, 2, 3)
	var sl4 []int
	// sl4 = sl4[2:]
	fmt.Println(sl4, len(sl4), cap(sl4))

	for i := 0; i < 10; i++ {
		if i%5 == 0 {
			fmt.Println("for: ", sl4, len(sl4), cap(sl4))
		}
		sl4 = append(sl4, i)
	}
	fmt.Println(sl4, len(sl4), cap(sl4))

	sl5 := []string{99: "A"}
	fmt.Println(len(sl5[0]))

	sl6 := make([]int, 0, 2)
	// sl6 := []int{}

	for i := 0; i < 10; i++ {
		if i%5 == 0 {
			fmt.Println("for: ", sl6, len(sl6), cap(sl6))
		}
		sl6 = append(sl6, i)
	}
	fmt.Println(sl6, len(sl6), cap(sl6))
	fmt.Println("=========1=========")

	sl7 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl8 := sl7[1:3]

	sl8[0] = 100
	fmt.Println(sl7)
	fmt.Println(sl8)

	fmt.Println("=========2=========")
	// Tạo một slice kiểu integer.
	slice := []int{10, 20, 30, 40, 50}

	// Tạo mới một slice.
	newSlice := slice[1:3]
	fmt.Println(slice, newSlice)
	// Thay đổi giá trị ở index là 2 ở slice.
	newSlice[1] = 35
	fmt.Println(slice, newSlice)
}

func myMap() {
	m1 := map[string]string{"name": "kuro"}

	fmt.Println(m1)

	v, e := m1["abc"]

	fmt.Println(v, e)
}

func myVariadicFunction() {
	myVariadicFunction2(1, 2, 3, 4, 5, 6)

	var list = []int{1, 2, 3, 4}
	myVariadicFunction2(1, list...)

	myVariadicFunctionChange(list...)
	fmt.Println(list)
}

func myVariadicFunction2(item int, list ...int) {
	list = append(list, item)
	fmt.Println(list)
}

func myVariadicFunctionChange(list ...int) {
	list[0] = 100
}

func myMap1() {
	var mm = make(map[string]int)
	fmt.Println(mm)

	var mm1 map[string]int
	fmt.Println(mm1)

	mm2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	mm2["four"] = 4
	mm2["five"] = 5
	fmt.Println(mm2)

	delete(mm2, "five")
	fmt.Println(mm2)

	fmt.Println(len(mm2))
}

func myPointer() {
	a := 10
	var p1 *int
	p1 = &a

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println("============")
	b := 20
	p2 := new(int)
	p2 = &b

	fmt.Println(p2)
	fmt.Printf("%T\n", p2)
}

func myPointer2() {
	a := 100
	var p1 *int
	p1 = &a

	*p1 = 200

	fmt.Println(a, p1)
}

func myPointer3() {
	array := [3]int{1, 2, 3}
	var p1 *[3]int = &array

	fmt.Println(p1)
	fmt.Println("============")
	a := 30
	var p2 *int = &a
	myPointer4(p2)
	fmt.Println(a, *p2)
}

func myPointer4(pointer *int) {
	*pointer = 1000
}

func MBasic() {}
