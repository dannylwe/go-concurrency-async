package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	urls := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://cnn.com",
		"https://netflix.com",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		// increment waitgroup counter
		wg.Add(1)

		go func(url string) {
			//decrement waitgroup counter
			defer wg.Done()
			check(url)
		}(url)
	}
	// wait for all goroutines to complete
	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("check took %v", elapsed) // 3.9252245s; approx 50% decrease in wait time
}

// check is a functions that pings a website to determine if is up or down
func check(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is not online")
	}
	fmt.Println(url, "is online")
}
