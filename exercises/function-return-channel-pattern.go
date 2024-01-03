package exercises

import (
	"fmt"
	"log"
)

func printSomething(msg string) chan string {

	result := make(chan string)

	go func() {
		for i := 0; i <= 100; i++ {
			result <- fmt.Sprintf("%s: %d", msg, i)
		}
	}()

	return result
}

func FRCP() {
	log.Println("Begin!")

	bridge := printSomething("Hello")

	for i := 0; i <= 100; i++ {
		log.Println("receive from", <-bridge)
	}

	log.Println("\nEnd!")
}
