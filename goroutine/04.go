package goroutine

import (
	"log"
	"sync"
)

func MChannelV2() {
	log.Println("Start!")

	var wg sync.WaitGroup
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)
	quit := make(chan int, 10)
	wg.Add(2)

	go func(ch <-chan int) { // channel only receive
		defer wg.Done()

		// for i := range ch {
		// 	log.Println("Show value - 1: ", i)
		// }

		for {
			select {
			case i := <-ch:
				log.Println("ch - show value 1: ", i)
			case i := <-ch2:
				log.Println("ch - show value 2: ", i)
			case q := <-quit:
				log.Println("ch - show value 3: ", q)
				return
			default:
				return
			}
		}
	}(ch)

	go func(ch chan<- int) { // channel only send
		ch <- 2
		// close(ch)

		ch2 <- 3
		// close(ch2)

		quit <- 0
		close(quit)
		wg.Done()
	}(ch)

	wg.Wait()

	log.Println("End!")
}
