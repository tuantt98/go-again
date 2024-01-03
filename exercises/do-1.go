package exercises

import (
	"log"
)

func Do1() {
	log.Println("====")
	chan1 := make(chan int)

	go func() {
		chan1 <- 1
	}()

	value := <-chan1

	log.Println("okok", value)
}
