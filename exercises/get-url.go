package exercises

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type SGetUrl struct {
	urlChan chan string
	url     string
}

func getUrlContent(url string, result chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	urlContext, errReadBody := io.ReadAll(resp.Body)

	if errReadBody != nil {
		fmt.Println(err)
		return
	}

	result <- string(urlContext)
	defer resp.Body.Close()
}

func GetUrl() {
	fmt.Println("Begin")
	f, err := os.OpenFile("exercises/data/save.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		return
	}

	var listGetUrl = []SGetUrl{
		{urlChan: make(chan string), url: "https://www.youtube.com/"},
		{urlChan: make(chan string), url: "http://google.com/"},
		{urlChan: make(chan string), url: "https://www.github.com/"},
	}

	listGetUrl = append(listGetUrl, SGetUrl{
		urlChan: make(chan string),
		url:     "https://bard.google.com/",
	})

	for i := 0; i < len(listGetUrl); i++ {

		go getUrlContent(listGetUrl[i].url, listGetUrl[i].urlChan)
	}
	for i := 0; i < len(listGetUrl); i++ {
		_, err := f.WriteString(<-listGetUrl[i].urlChan)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	defer f.Close()

	fmt.Println("End!")
}
