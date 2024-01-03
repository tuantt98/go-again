package exercises

import (
	"log"
)

func fibonacci(c chan int, quit chan int) {
	log.Println("fibonacci")
	x, y := 0, 1
	for {
		select {
		case c <- x:
			log.Printf("Before: set value c: %d\n", x)
			x, y = y, x+y
		case <-quit:
			log.Println("quit")
			return
		}
	}
}

func Do2() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			log.Println("Get value: ", <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
