package main

import "fmt"

func main() {
	// slice1()
	// slice2()
	// slice3()
	// slice4()
}

func slice1() {
	var a1 [100]byte
	var a2 [100]int32

	fmt.Printf("array 1 length:   %d\n", len(a1))
	fmt.Printf("array 2 length:   %d\n", len(a2))
	fmt.Println()

	var slice0 []byte = a1[:]
	fmt.Printf("slice 0 length:   %d\n", len(slice0))
	fmt.Printf("slice 0 capacity: %d\n", cap(slice0))
	fmt.Println()

	var slice1 []byte = a1[10:20]
	fmt.Printf("slice 1 length:   %d\n", len(slice1))
	fmt.Printf("slice 1 capacity: %d\n", cap(slice1))
	fmt.Println()

	var slice2 = a1[20:30]
	fmt.Printf("slice 2 length:   %d\n", len(slice2))
	fmt.Printf("slice 2 capacity: %d\n", cap(slice2))
	fmt.Println()

	slice3 := a1[30:40]
	fmt.Printf("slice 3 length:   %d\n", len(slice3))
	fmt.Printf("slice 3 capacity: %d\n", cap(slice3))
}

//- slice can be created from any array
//- but the slice does not contain any data, just a pointer to array element
//- so any modify in slice element is reflected in an array as well
func slice2() {
	var a [10]int

	slice := a[:]

	fmt.Printf("array before modification: %v\n", a)
	fmt.Printf("slice before modification: %v\n", slice)
	fmt.Println()

	// cont
	for i := 0; i < len(a); i++ {
		a[i] = i * 2
	}

	fmt.Printf("array after modification:  %v\n", a)
	fmt.Printf("slice after modification:  %v\n", slice)
	fmt.Println()

	for i := 0; i < len(slice); i++ {
		slice[i] = 42
	}

	fmt.Printf("array after modification:  %v\n", a)
	fmt.Printf("slice after modification:  %v\n", slice)
}

//- modify array elements
//- then modify the same elements, but via slice
func slice3() {
	var a [10]int

	slice1 := a[4:9]
	slice2 := slice1[3:]

	fmt.Printf("array:            %v\n", a)
	fmt.Printf("array length:     %d\n\n", len(a))

	fmt.Printf("slice 1:          %v\n", slice1)
	fmt.Printf("slice 1 length:   %d\n", len(slice1))
	fmt.Printf("slice 1 capacity: %d\n\n", cap(slice1))

	fmt.Printf("slice 2:          %v\n", slice2)
	fmt.Printf("slice 2 length:   %d\n", len(slice2))
	fmt.Printf("slice 2 capacity: %d\n\n", cap(slice2))

	// cont

	slice2[0] = 99
	slice2[1] = 99

	fmt.Printf("array:            %v\n", a)
	fmt.Printf("slice 1:          %v\n", slice1)
	fmt.Printf("slice 2:          %v\n", slice2)
}

func slice4() {
	var slice []int

	for i := 1; i < 11; i++ {
		slice = append(slice, i)
		fmt.Printf("slice:          %v\n", slice)
		fmt.Printf("slice length:   %d\n", len(slice))
		fmt.Printf("slice capacity: %d\n\n", cap(slice))
	}
}
