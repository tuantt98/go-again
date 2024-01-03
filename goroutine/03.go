package goroutine

import (
	"log"
	"sync"
)

func MChannel() {
	log.Println("Start!")

	var wg sync.WaitGroup
	ch := make(chan int, 10)

	wg.Add(2)

	go func(ch <-chan int) { // channel only receive
		// for i := range ch {
		// 	log.Println("Show value - 1: ", i)
		// }

		for {
			if i, ok := <-ch; ok {
				log.Println("Show value: ", i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // channel only send
		ch <- 2
		ch <- 3
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

	log.Println("End!")
}
