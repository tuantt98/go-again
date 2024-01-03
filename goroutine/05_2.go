package goroutine

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func generatePipeline2(numbers []int) chan int {
	out := make(chan int)

	go func() {
		for _, n := range numbers {
			// log.Println("generatePipeline: ", n)
			out <- n
		}
		close(out)
	}()

	return out
}

func fanOut2(in chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			// log.Println("fanOut: ", n)
			out <- n * n
		}
		close(out)
	}()

	return out
}

func fanIn2(inputChan ...chan int) chan int {
	in := make(chan int)

	go func() {
		for _, ch := range inputChan {
			for n := range ch {
				// log.Println("fanIn: ", n)
				in <- n
			}
		}
		close(in)
	}()

	return in
}

func MChannelV3_2() {
	randomNumbers := []int{}
	numberP := runtime.NumCPU()
	totalChan := []chan int{}

	for i := 1; i <= 1_000_000; i++ {
		randomNumbers = append(randomNumbers, i)
	}

	// Pipeline
	inputChan := generatePipeline2(randomNumbers)

	// Fan-out
	for i := 1; i <= numberP; i++ {
		a := fanOut2(inputChan)
		totalChan = append(totalChan, a)
	}
	// fmt.Println("Len(totalChan): ", len(totalChan))

	// Fan-in
	c := fanIn2(totalChan...)

	sum := 0

	for i := range c {
		sum += i
	}

	fmt.Println("Sum (MChannelV3_2): ", sum)

	defer timeTrack(time.Now(), "MChannelV3_2")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
