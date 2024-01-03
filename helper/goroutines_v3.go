package helper

import (
	"fmt"
)

func printNumberV1(numChan chan int) {
	var result int
	for i := 0; i <= 100; i++ {
		// fmt.Printf("%d ", i)
		result += i
	}
	numChan <- result
}

func printCharV1(strChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		// fmt.Printf("%c ", i)
		result += string(i)
	}
	strChan <- result
}

func GoroutinesV3() {

	chanPrintNumber := make(chan int)
	chanPrintChar := make(chan string)

	fmt.Println("Begin!")

	go printNumberV1(chanPrintNumber)
	go printCharV1(chanPrintChar)

	fmt.Println(<-chanPrintNumber)
	fmt.Println(<-chanPrintChar)
	fmt.Println("\nEnd!")
}
