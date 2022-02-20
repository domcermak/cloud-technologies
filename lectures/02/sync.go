package main

import (
	"fmt"
	"sync"
)

func main() {
	//sync1()
	//sync2()
	sync3()
}

// multiple producers/consumers with wait.Group
func sync1() {
	c := make(chan string, 2)

	producersCount := 5
	consumersCount := 2

	pwg := sync.WaitGroup{}
	pwg.Add(producersCount)

	for i := 0; i < producersCount; i++ {
		// ???
		go func(i int) {
			defer pwg.Done()
			for j := 0; j < 5; j++ {
				c <- fmt.Sprintf("p%v-%v", i, j)
			}
		}(i)
	}

	cwg := sync.WaitGroup{}
	cwg.Add(consumersCount)

	for i := 0; i < consumersCount; i++ {
		go func(i int) {
			defer cwg.Done()
			for el := range c {
				fmt.Println(fmt.Sprintf("consumer %v: %v", i, el))
			}
		}(i)
	}

	pwg.Wait()
	close(c)

	cwg.Wait()
	fmt.Println("done")
}

// sync.Cond
func sync3() {
	sharedRsc := []int{}

	var wg sync.WaitGroup
	wg.Add(2)

	m := sync.Mutex{}
	c := sync.NewCond(&m)

	go func() {
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc[0])
		sharedRsc = sharedRsc[1:]
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc[0])
		sharedRsc = sharedRsc[1:]
		c.L.Unlock()
		wg.Done()
	}()

	c.L.Lock()
	sharedRsc = append(sharedRsc, 1, 2)
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}
