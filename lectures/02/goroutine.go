package main

import (
	"fmt"
	"time"
)

func goroutinesDo(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go goroutinesDo("first")
	go goroutinesDo("second")
	goroutinesDo("main")
}
