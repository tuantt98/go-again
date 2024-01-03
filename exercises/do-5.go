package exercises

import (
	"io"
	"log"
	"net/http"
)

type UrlInfo struct {
	url     string
	content chan string
}

func getContentFromUrl(url string) (string, error) {

	var result string
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return result, err
	}

	urlContext, errReadBody := io.ReadAll(resp.Body)

	if errReadBody != nil {
		log.Println(err)
		return result, err
	}

	result = string(urlContext)

	defer resp.Body.Close()

	return result, nil
}

func getContent(url UrlInfo) {
	log.Printf("Start: get content from: %s", url.url)

	result, err := getContentFromUrl(url.url)

	if err != nil {
		log.Println("Error")
		panic(err)
	}
	url.content <- result

	log.Printf("End: get content from: %s", url.url)

}

func Do5() {

	listMapUrl := map[string]UrlInfo{
		"google.com":  {url: "https://google.com/", content: make(chan string)},
		"youtube.com": {url: "https://youtube.com/", content: make(chan string)},
	}

	for i, v := range listMapUrl {
		log.Println("Web: ", i)
		go getContent(v)
		<-v.content
	}

}
