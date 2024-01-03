package helper

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup, numChan chan int) {
	var result int
	for i := 0; i <= 100; i++ {
		// fmt.Printf("%d ", i)
		result += i
	}
	numChan <- result
	wg.Done()
}

func printChar(wg *sync.WaitGroup, strChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		// fmt.Printf("%c ", i)
		result += string(i)
	}
	strChan <- result
	wg.Done()
}

func GoroutinesV2() {

	chanPrintNumber := make(chan int)
	chanPrintChar := make(chan string)

	fmt.Println("Begin!")
	var wg sync.WaitGroup

	wg.Add(2)

	go printNumber(&wg, chanPrintNumber)
	go printChar(&wg, chanPrintChar)

	fmt.Println(<-chanPrintNumber)
	fmt.Println(<-chanPrintChar)
	wg.Wait()
	fmt.Println("\nEnd!")
}
