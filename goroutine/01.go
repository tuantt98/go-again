package goroutine

import (
	"log"
	"sync"
	"time"
)

func count(name string) {
	for i := 0; i <= 5; i++ {
		log.Println(name, i)
		time.Sleep(time.Second)
	}
}

func MGoRoutine() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		count("Sheep")
		wg.Done()
	}()
	go func() {
		count("fish")
		wg.Done()
	}()

	wg.Wait()
}
