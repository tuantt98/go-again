package helper

import "fmt"

func (s Student) getName() (name string) {
	return s.name
}

// Value receiver
func (s Student) changeName(name string) {
	s.name = name
}

// Pointer receiver
func (s *Student) updateName(name string) {
	s.name = name
}

func MyFunc() {
	student := Student{
		id: 1, name: "kuro",
		info: StudentInfo{
			class: "A2",
			email: "kuro@gmail.com",
			age:   18,
		},
	}

	name := student.getName()
	fmt.Println(name)

	fmt.Println("================1")
	student.changeName("kuro_v2")
	fmt.Println(student)

	fmt.Println("================2")
	student.setName("kuro_v2")
	fmt.Println(student)

	fmt.Println("================3")
	student.updateName("kuro_v2")
	fmt.Println(student)
}
