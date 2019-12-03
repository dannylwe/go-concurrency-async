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
		check(url)
	}

	elapsed := time.Since(start)
	log.Printf("check took %v", elapsed)
}

func check(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is not online")
	}
	fmt.Println(url, "is online")
}
