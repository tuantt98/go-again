package exercises

import (
	"log"
	"time"
)

// google search

func googleSearch(result chan string) {
	time.Sleep(time.Second)
	result <- "Found from Google"
}

// bing search
func bingSearch(result chan string) {
	time.Sleep(time.Second)
	result <- "Found from Bing"
}

func MSelect() {

	log.Println("Begin!")

	chanGoogle := make(chan string)
	chanBing := make(chan string)

	go googleSearch(chanGoogle)
	go bingSearch(chanBing)

	select {
	case result := <-chanGoogle:
		log.Println(result)
	case result := <-chanBing:
		log.Println(result)
	}

	log.Println("End main!")

}
