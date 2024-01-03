package exercises

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func Do4() {
	n := 10000
	maxWorker := 5
	queueCh := make(chan int, n)
	var wg sync.WaitGroup
	f, err := os.OpenFile("exercises/data/save_v2.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 1; i <= n; i++ {
		queueCh <- i
	}
	close(queueCh)

	wg.Add(maxWorker)
	for i := 1; i <= maxWorker; i++ {
		go func(count int) {
			for v := range queueCh {
				time.Sleep(time.Millisecond * 10)
				log.Printf("Worker %d is crawling web url %d\n", count, v)

				// Write result to file save_v2.txt
				content := fmt.Sprintf("Worker %d is crawling web url %d\n", count, v)
				_, errWrite := f.WriteString(content)
				if errWrite != nil {
					log.Println(err)
					return
				}

			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
