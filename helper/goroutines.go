package helper

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func g1() {
	fmt.Println("g1")
	wg.Done()
}

func g2() {
	fmt.Println("g2")
	wg.Done()
}

var wg sync.WaitGroup

func Goroutines() {
	go g1()

	ng := runtime.NumGoroutine()
	fmt.Println(ng)
	time.Sleep(time.Second)

	fmt.Println("begin")
	wg.Add(2)
	go g1()
	go g2()

	wg.Wait()

	fmt.Println("end")

}
