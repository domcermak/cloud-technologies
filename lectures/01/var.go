package main

import "fmt"

const ac = 5

var (
	vara int    = 3
	vars string = "Hello world"
)

func main() {
	const bc = 6

	var l int = 4
	k := 5

	a := float64(5)

	u64 := uint64(a)

	fmt.Println(l, k, a, u64, bc, ac)
}
