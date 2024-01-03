package helper

import "fmt"

type StudentInfo struct {
	class string
	email string
	age   int
}

type Student struct {
	id   int
	name string
	info StudentInfo
}

func (s Student) setName(name string) {
	s.name = name
}

func HandleStruct() {
	// st1 := Student{
	// 	id:   01,
	// 	name: "kuro",
	// }

	// fmt.Println(st1)
	// fmt.Println(st1.id)

	// fmt.Println("=======")

	// st2 := Student{2, "kuro"}
	// fmt.Println("🚀 ~ file: my-struct.go:22 ~ funcHandleStruct ~ st2:", st2)

	// var st3 = struct {
	// 	email string
	// 	age   int
	// }{
	// 	email: "kuro@gmail.com",
	// 	age:   18,
	// }

	// fmt.Println("🚀 ~ file: my-struct.go:25 ~ funcHandleStruct ~ st3:", st3)

	// st4 := &Student{
	// 	id:   02,
	// 	name: "kuro_02",
	// }
	// fmt.Println(&st4)
	// fmt.Println(st4.id)

	student1 := Student{
		id:   001,
		name: "kuro",
		info: StudentInfo{
			class: "A2",
			email: "kuro@gmail.com",
			age:   18,
		},
	}
	student2 := Student{
		id:   001,
		name: "kuro",
		info: StudentInfo{
			class: "A2",
			email: "kuro@gmail.com",
			age:   18,
		},
	}
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student1 == student2)
}
