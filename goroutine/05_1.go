package goroutine

import (
	"fmt"
	"time"
)

func generatePipeline(numbers []int) <-chan int {
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

func fanOut(in <-chan int) <-chan int {
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

func fanIn(inputChan ...<-chan int) <-chan int {
	in := make(chan int)

	go func() {
		for _, ch := range inputChan {
			for n := range ch {
				// log.Println("fanIn: ", n)
				in <- n
			}
		}
	}()

	return in
}

func MChannelV3() {
	randomNumbers := []int{}

	for i := 1; i <= 1_000_000; i++ {
		randomNumbers = append(randomNumbers, i)
	}

	// Pipeline
	inputChan := generatePipeline(randomNumbers)

	// Fan-out
	c1 := fanOut(inputChan)
	c2 := fanOut(inputChan)
	c3 := fanOut(inputChan)
	c4 := fanOut(inputChan)

	// Fan-in
	c := fanIn(c1, c2, c3, c4)

	sum := 0

	for i := 0; i < len(randomNumbers); i++ {
		sum += <-c
	}

	fmt.Println("Sum (MChannelV3): ", sum)
	defer timeTrack(time.Now(), "MChannelV3")
}
