package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func swap2(a, b int) (x, y int) {
	y = a
	x = b
	return
}

func main() {
	fmt.Println(sum(3, 4))
	fmt.Println(swap(3, 4))
	fmt.Println(swap2(3, 4))
}
