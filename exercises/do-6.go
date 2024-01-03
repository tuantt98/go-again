package exercises

import (
	"log"
	"sync"
)

func handleNumber(maxNumber int, result chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 1; i <= maxNumber; i++ {
			log.Printf("Handle set number: %d\n", i)
			result <- i * i
		}
		close(result)
	}()

}

func Do6() {

	log.Println("Start!")

	n := 100
	var wg sync.WaitGroup
	// wg.Add(n)

	result := make(chan int)
	go handleNumber(n, result, &wg)

	for v := range result {
		log.Printf("Handle show number: %d\n", v)
	}

	log.Println("End!")
	// wg.Wait()
}
