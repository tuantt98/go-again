package goroutine

import (
	"log"
	"sync"
)

var wg sync.WaitGroup
var counter = 0
var mt sync.RWMutex

func sayHello() {
	log.Printf("Hello, %d\n", counter)
	mt.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	mt.Unlock()
	wg.Done()
}

func MGoRoutineV2() {
	log.Println("Start!")

	for i := 0; i < 10; i++ {
		wg.Add(2)
		mt.RLock()
		go sayHello()
		mt.Lock()
		go increment()
	}

	wg.Wait()

	log.Println("End!")
}
