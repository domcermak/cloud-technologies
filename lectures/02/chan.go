package main

import (
	"fmt"
	"time"
)

func main() {
	chan1()
	//chan2()
	//chan3()
	//chan4()
}

func receiveAndPrint(c chan int) {
	a := <-c
	fmt.Println("Received", a)
}

func chan1() {
	c := make(chan int)
	go receiveAndPrint(c)
	c <- 5
	c <- 6
}

func chan2() {
	c := make(chan int)
	go func() {
		for i := range c {
			fmt.Println(i)
		}
		fmt.Println("channel closed")
	}()

	for i := 0; i < 5; i++ {
		c <- i * 2
	}

	close(c)
	time.Sleep(time.Second)
}

func chan3() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for {
			select {
			case x := <-c:
				fmt.Println(x)
			case <-quit:
				quit <- 1
				return
			}
		}
	}()
	quit <- 1
	<-quit
	fmt.Println("all done")
}

func chan4() {
	quit := make(chan int)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("Sleeping")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	time.Sleep(time.Second * 2)
	quit <- 1
}

func chan5() {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		c <- i
	}

	for el := range c {
		fmt.Println(el)
	}
}
