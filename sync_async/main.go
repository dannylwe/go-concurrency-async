package main

import (
	"fmt"
	"log"
	"net/http"
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
	
	for _, url := range urls {
		go check(url)
	}

	elapsed := time.Since(start)
	log.Printf("check took %v", elapsed) //19.6841258s
}

// check is a functions that pings a website to determine if is up or down
func check(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is not online")
	}
	fmt.Println(url, "is online")
}
