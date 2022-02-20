package main

import "fmt"

func main() {
	if1()
	if2()
	switch1()
	switch2()
	switch3()
}

func if1() {
	var i = 5

	if i < 6 {
		fmt.Println(i, "< 6")
	}

	if j := 10; j < 11 {
		fmt.Println(j, "< 11")
	}
}

func if2() {
	if j := 10; j < 9 {
		fmt.Println(j, "< 9")
	} else {
		fmt.Println(j, "< 9 is false")
	}
}

func switch1() {
	switch i := 5; i {
	case 4:
		fmt.Println(i, "= 4")
	case 5:
		fmt.Println(i, "= 5")
	default:
		fmt.Println(i, "didn't match any rule")
	}
}

func switch2() {
	i := 5

	switch {
	case i < 5:
		fmt.Println(i, "< 5")
	case i > 5:
		fmt.Println(i, "> 5")
	default:
		fmt.Println(i, "is not less or more than 5")
	}
}

func switch3() {
	f := func() int {
		return 5
	}

	var value = 6

	switch i := 5; i {
	case value:
		fmt.Println(i, "= value")
	case f():
		fmt.Println(i, "= f()")
	case 4 + 1:
		fmt.Println(i, "= 4 + 1")
	}
}
