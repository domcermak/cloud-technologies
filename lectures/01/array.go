package main

import "fmt"

//- basic data type in the Go programming language
//- all array items has the same type
//- (well, you can use `interface{}` to allow _dynamic_typing_behaviour_)
//- type of array is derived from type of items *and* array size
//- (unlike slices)
//- index in range 0..length-1
//- items indexing via [] (as in most other languages)
func main() {
	arrays1()
	arrays2()
}

func arrays1() {
	var a1 [10]byte
	var a2 [10]int32
	a3 := [10]int32{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}

	fmt.Printf("array 1 length: %d\n", len(a1))
	fmt.Printf("array 2 length: %d\n", len(a2))
	fmt.Printf("array 3 length: %d\n", len(a3))

	var a [10]int

	fmt.Printf("Original array: %v\n", a)

	for i := 0; i < len(a1); i++ {
		a[i] = i * 2
	}

	fmt.Printf("Modified array: %v\n", a)
}

func arrays2() {
	var matrix [4][3]float32
	fmt.Printf("matrix: %v\n", matrix)

	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(matrix[j]); i++ {
			matrix[j][i] = 1.0 * float32(i+1) * float32(j+1)
		}
	}
	fmt.Printf("matrix: %v\n", matrix)
}

// unlike slices, arrays can be copied
func array3() {
	var a1 [10]int

	a2 := a1

	fmt.Printf("array 1: %v\n", a1)
	fmt.Printf("array 2: %v\n", a2)

	for i := 0; i < len(a1); i++ {
		a1[i] = i * 2
	}

	fmt.Println("-------------------------------")

	fmt.Printf("array 1: %v\n", a1)
	fmt.Printf("array 2: %v\n", a2)
}
