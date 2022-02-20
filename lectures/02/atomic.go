package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//atomic1()
	//atomic2()
	//atomic3()
}

func atomic1() {
	i := 0

	wg := sync.WaitGroup{}
	wg.Add(10)

	for j := 0; j < 10; j++ {
		go func() {
			defer wg.Done()
			for k := 0; k < 1e5; k++ {
				i++
			}
		}()
	}

	wg.Wait()

	// should be 1e6
	fmt.Println(i)
}

func atomic2() {
	i := uint64(0)

	wg := sync.WaitGroup{}
	wg.Add(10)

	for j := 0; j < 10; j++ {
		go func() {
			defer wg.Done()
			for k := 0; k < 1e5; k++ {
				atomic.AddUint64(&i, 1)
			}
		}()
	}

	wg.Wait()

	// should be 1e6
	// why not atomic.LoadUint64 ???
	fmt.Println(i)
}

func atomic3() {
	i := uint64(0)

	fmt.Println(atomic.SwapUint64(&i, 5))
	fmt.Println(atomic.CompareAndSwapUint64(&i, 4, 3))
	fmt.Println(atomic.CompareAndSwapUint64(&i, 5, 3))
	fmt.Println(atomic.LoadUint64(&i))
}
