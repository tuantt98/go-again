package exercises

import (
	"log"
)

func firstInput(numbers ...int) chan int {

	result := make(chan int)

	go func() {
		for i := 0; i < len(numbers); i++ {
			result <- numbers[i]
		}
		close(result)
	}()

	return result
}

func secondInput(chanNumbers chan int) chan int {
	result := make(chan int)
	go func() {
		for item := range chanNumbers {
			result <- item * item
		}
		close(result)
	}()
	return result
}

func Pipeline() {
	log.Println("Begin!")

	firstChan := firstInput(1, 2, 3, 4, 5, 6)
	secondChan := secondInput(firstChan)

	for item := range secondChan {
		log.Println("receive: ", item)
	}

	log.Println("End!")
}
