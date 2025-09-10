package main

import (
	"fmt"
	"time"
)

func ping(ch chan string) {
	for {
		ch <- "ping"
		time.Sleep(500 * time.Millisecond)
	}
}

func pong(ch chan string) {
	for {
		ch <- "pong"
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch := make(chan string)

	go ping(ch)
	go pong(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
