package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// BinaryOp represents any function that takes two int parameters and returns int as a result
type BinaryOp func(x, y int) int

func applyBinaryOp(a, b int, operation BinaryOp) int {
	return operation(a, b)
}

func testBinaryOps(a, b int) {
	f1 := func(a, b int) int { return a + b }
	f2 := func(a, b int) int { return a - b }
	f3 := func(a, b int) int { return a * b }
	f4 := func(a, b int) int { return a / b }
	fmt.Printf("%d + %d = %d\n", a, b, applyBinaryOp(a, b, f1))
	fmt.Printf("%d - %d = %d\n", a, b, applyBinaryOp(a, b, f2))
	fmt.Printf("%d * %d = %d\n", a, b, applyBinaryOp(a, b, f3))
	fmt.Printf("%d / %d = %d\n", a, b, applyBinaryOp(a, b, f4))
}

// - lambda functions that refer to variables defined outside the function
// - it "closes over" another function -> the name closure
// - code pointer and environment pointer internally
// - useful for function that needs to store its state "somewhere"
func testBinaryOpsLambda(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, func(a, b int) int { return a + b }(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, func(a, b int) int { return a - b }(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, func(a, b int) int { return a * b }(a, b))
	fmt.Printf("%d / %d = %d\n", a, b, func(a, b int) int { return a / b }(a, b))
}

func testSort() {
	numbers := make([]int, 20)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Int() % 100
	}
	fmt.Println(numbers)

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	fmt.Println(numbers)
}

func main() {
	testBinaryOps(1, 2)
	//testBinaryOpsLambda(1, 2)
	testSort()
}
