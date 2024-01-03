package helper

import (
	"context"
	"fmt"
	"time"
)

func Demo() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		for {
			// Thực hiện một tác vụ dài hạn
			if ctx.Err() != nil {
				fmt.Println("Tác vụ bị hủy bỏ")
				return
			}
		}
	}()
	time.Sleep(time.Second * 10)
}

func s1() {
	fmt.Println("s1")
}

func s2() {
	fmt.Println("s2")
}
func s3() {
	fmt.Println("s3")
}
func s4() {
	fmt.Println("s4")
}

func MyScheduler() {
	// runtime.GOMAXPROCS(2)
	// numberP := runtime.NumCPU()
	// fmt.Println(numberP)

	// numberMax := runtime.GOMAXPROCS(0)
	// fmt.Println(numberMax)
	go s1()
	go s2()
	go s3()
	go s4()
}
