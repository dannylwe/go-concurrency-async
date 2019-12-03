package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type urlStatus struct {
	url    string
	status bool
}

func main() {
	start := time.Now()

	urls := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://cnn.com",
		"https://netflix.com",
	}

	c := make(chan urlStatus)

	for _, url := range urls {
		go check(url, c)
	}

	result := make([]urlStatus, len(urls))
	for i := range result {
		result[i] = <-c
		if result[i].status {
			fmt.Println(result[i].url, "is up.")
		} else {
			fmt.Println(result[i].url, "is down !!")
		}
	}

	elapsed := time.Since(start)
	log.Printf("check took %v", elapsed) //3.1611808s
}

// check is a functions that pings a website to determine if is up or down
func check(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		c <- urlStatus{url, false}
	}
	c <- urlStatus{url, true}
}
