package exercises

import (
	"fmt"
	"log"
)

func printSomethingV2(msg string) chan string {

	result := make(chan string)

	go func() {
		for i := 0; i <= 5; i++ {
			result <- fmt.Sprintf("%s: %d", msg, i)
		}
	}()

	return result
}

func fanIn(chan1, chan2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case <-chan1:
				c <- <-chan1
			case <-chan2:
				c <- <-chan2
			}
		}
	}()

	return c
}

func FanInPattern() {
	log.Println("Begin!")

	coffee := printSomethingV2("Coffee order")
	bread := printSomethingV2("Bread order")

	serve := fanIn(coffee, bread)

	for i := 0; i <= 5; i++ {
		log.Println("receive from", <-serve)
	}

	log.Println("End!")
}
