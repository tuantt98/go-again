package helper

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("ping %d", i)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func MyConcurrency() {
	var c chan string = make(chan string)

	go printer(c)
	go pinger(c)

	var input string
	fmt.Scanln(&input)
}
