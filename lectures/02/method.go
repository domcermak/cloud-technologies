package main

import "fmt"

type MyStruct struct {
	a, b int64
}

func (s *MyStruct) add(a, b int64) {
	s.a += a
	s.b += b
}

func (s MyStruct) addValue(a, b int64) {
	s.a += a
	s.b += b
}

func main() {
	// methods1()
	// methods2()
	// method3()
	method4()
}

func method1() {
	line1 := MyStruct{}

	fmt.Println(line1)

	line1.add(5, 5)
	fmt.Println(line1)

	line1.addValue(5, 5)
	fmt.Println(line1)
}

func method2() {
	line1 := MyStruct{}

	fmt.Println(line1)

	// copy by value
	cpy := line1

	line1.add(5, 5)
	fmt.Println(line1)
	fmt.Println(cpy)
}

func method3() {
	line1 := &MyStruct{}

	fmt.Println(line1)

	// copy pointer
	cpy := line1

	line1.add(5, 5)
	fmt.Println(line1)
	fmt.Println(cpy)
}

func method4() {
	line1 := &MyStruct{}

	fmt.Println(line1)

	line1.addValue(5, 5)
	// ???
	fmt.Println(line1)
}
