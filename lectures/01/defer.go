package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

func main() {

	//defer1()
	//defer2()
	//defer3()
	//defer4()
	//defer5()
	//defer6()

	// practicalDefer()
}

// * defer statement
// - `defer` is a keyword in the Go programming language
// - used to "remember" commands that will be called before `return` or exit
// - based on LIFO (stack) of remembered commands
// - parameters are evaluated when `defer` is declared (ie. in runtime)
// - (not when the specified code is called)
// - it is possible to change function return value(s) via `defer`

func defer1OnFinish() {
	fmt.Println("finished")
}

func defer1() {
	defer defer1OnFinish()

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("finishing defer1 function")
}

// using lambda function and closure (will be described later)
func defer2() {
	defer (func() { fmt.Println("finished") })()

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("finishing defer2 function")
}

// LIFO behaviour
func defer3() {
	fn := func(i int) { fmt.Println(i) }

	for i := 0; i <= 10; i++ {
		// defer in for loop is almost always bad practice
		defer fn(i)
	}
	fmt.Println("finishing defer3 function")
}

// actual parameters are evaluated in runtime
// LIFO behaviour
func defer4() {
	function := func(i int) {
		fmt.Printf("defer %2d\n", i)
	}

	x := 0

	fmt.Printf("value = %2d\n", x)
	defer function(x)

	x++

	fmt.Printf("value = %2d\n", x)
	defer function(x)

	x++
	fmt.Printf("value = %2d\n", x)

	fmt.Println("finishing defer4 function")
}

// - arrays are a bit tricky
// - (call by value vs. call by reference)
func defer5() {
	function := func(a []int) {
		fmt.Printf("defer %v\n", a)
	}

	var x = []int{1, 2, 3}
	fmt.Printf("value = %v\n", x)
	defer function(x)

	x[0] = 0
	fmt.Printf("value = %v\n", x)
	defer function(x)

	x[1] = 0
	fmt.Printf("value = %v\n", x)
	defer function(x)

	x[2] = 0
	fmt.Printf("value = %v\n", x)
}

// accessing and updating named return value in defer
func defer6Function() (i int) {
	defer func() { i += 2 }()
	return 1
}

func defer6() {
	fmt.Println("defer6", defer6Function())
}

func practicalDefer() {
	path := "cond.go"

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(fmt.Errorf("open file %v: %w", path, err))
		return
	}

	// we can ignore close error in this case - why?
	defer f.Close()

	// file is closed on all return paths

	if rand.Int() > 5_000 {
		return
	}

	body, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(fmt.Errorf("read file %v: %w", path, err))
		return
	}

	fmt.Println(fmt.Sprintf("file %v size is %v bytes", path, len(body)))
}
