package exercises

import (
	"fmt"
	"os"
	"strings"
)

func countFirst(result chan int, filePath string, keyword string) {
	fmt.Println("countFirst")
	var numberOfOcc int
	fileContext, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
		result <- 0
		return
	}

	numberOfOcc = strings.Count(string(fileContext), keyword)

	result <- numberOfOcc

	defer close(result)
}

func countSecond(result chan int, filePath string, keyword string) {
	fmt.Println("countSecond")
	var numberOfOcc int
	fileContext, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
		result <- 0
		return
	}

	numberOfOcc = strings.Count(string(fileContext), keyword)

	result <- numberOfOcc

	defer close(result)
}

func CountText() {
	fmt.Println("Begin")
	countFirstChan := make(chan int)
	countSecondChan := make(chan int)

	go countFirst(countFirstChan, "exercises/data/text_1.txt", "go")
	go countSecond(countSecondChan, "exercises/data/text_2.txt", "Concurrency")

	fmt.Println("First: ", <-countFirstChan)
	fmt.Println("Second: ", <-countSecondChan)

	fmt.Println("End!")
}
