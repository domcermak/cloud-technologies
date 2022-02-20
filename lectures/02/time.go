package main

import (
	"fmt"
	"time"
)

func main() {
	//time1()
	//time2()
	//time3()
	//time4()
}

func time1() {
	fmt.Println(time.Second)
	fmt.Println(time.Minute + time.Second)

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.UnixMilli())
	fmt.Println(now.Add(time.Hour))

	fmt.Println(now.Format(time.RFC3339))
}

func time2() {
	now := time.Now()

	fmt.Println(now)
	<-time.After(2 * time.Second)
	fmt.Println(time.Now().Sub(now))
}

func time3() {
	timer := time.NewTimer(time.Second)

	for tick := range timer.C {
		// monotonic clock
		fmt.Println("tick", tick)
		timer.Reset(time.Second)
	}
}

func time4() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for tick := range ticker.C {
		fmt.Println("tick", tick)
	}
}
