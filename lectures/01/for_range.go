package main

import "fmt"

func main() {
	// forrange1()
	// forrange2()
	// forrange3()
	// forrange4()
}

func forrange1() {
	array := [4]string{"one", "two", "three", "four"}

	for index, item := range array {
		println(index, item)
	}
}

func forrange2() {
	array := [4]string{"one", "two", "three", "four"}

	for _, item := range array {
		println(item)
	}
}

func forrange3() {
	for _, ch := range "dobrý den" {
		fmt.Printf("%c", ch)
	}
}

func forrange4() {
	var m1 = make(map[int]string)
	m1[0] = "one"
	m1[1] = "two"
	m1[2] = "three"
	m1[3] = "four"

	for key, val := range m1 {
		println(key, val)
	}
}
