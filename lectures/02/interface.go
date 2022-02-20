package main

import "fmt"

type Adder interface {
	Add(a, b int) int
}

// AdderImpl is a user-defined data types that satisfy Adder interface
type AdderImpl struct {
}

func (a AdderImpl) Add(x, y int) int {
	return x + y
}

func main() {
	//interface1()
	//interface2()
	//interface3()
	//interface4()
	interface5()
}

func interface1() {
	a := AdderImpl{}
	fmt.Println(a.Add(1, 2))
}

type AddFunc func(a, b int) int

func (f AddFunc) Add(x, y int) int {
	return x + y
}

func doAdd(add Adder, a, b int) int {
	return add.Add(a, b)
}

func interface2() {
	res := doAdd(AddFunc(func(a, b int) int { return a + b }), 1, 2)
	fmt.Println(res)
}

func interface3() {
	var res []interface{}

	for i := 0; i < 5; i++ {
		res = append(res, int64(5))
	}

	for _, v := range res {
		fmt.Println(v.(int64))
	}
}

func interface4() {
	var i interface{} = "go...go...go!"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)
}

func interface5() {
	var p1 *int    /* = nil */
	var p2 *string /* = nil */

	var i1 interface{} /* = nil */
	var i2 interface{} = p1
	var i3 interface{} = p2

	fmt.Printf("%T\t%v\n", i1, i1)
	fmt.Printf("%T\t%v\n", i2, i2)
	fmt.Printf("%T\t%v\n", i3, i3)

	fmt.Println()
	fmt.Printf("%v\n", i1 == i2)
	fmt.Printf("%v\n", i1 == i3)
	fmt.Printf("%v\n", i2 == i3)
}
