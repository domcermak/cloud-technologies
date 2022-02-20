package main

import "fmt"

// - a.k.a. associative array or hash
// - container for key-value pairs
// - "nil map":
// 		var m1 map[int]string
// - empty map:
// 		var m3 map[int]string = make(map[int]string)
// 		m1 := make(map[int]string)
// - three basic operations: add/put, get, and delete
// - add/put items to a map:
// 		m3[0] = "zero"
// 		m3[1] = "one"
// 		m3[2] = "two"
// - get item from a map:
// 		m3[2]
// - delete from a map
// 		delete(m3, 0)

func main() {
	// map1()
	// map2()
	// map3()
	// map4()
	// map5()
	// map6()
}

func map1() {
	var m1 map[int]string
	fmt.Println(m1)

	m1[0] = "zero"
}

func map2() {
	var m1 map[int]string = make(map[int]string)
	fmt.Println(m1)

	m1[1] = "one"
	m1[2] = "two"

	fmt.Println(m1)
}

func map3() {
	m1 := map[int]string{
		1: "one",
		2: "two",
	}

	fmt.Println(m1)
}

func map4() {
	m := make(map[string]string)
	m["first"] = "1st"
	m["second"] = "2nd"
	m["third"] = "3rd"
	fmt.Printf("'%s'\n", m["first"])
	fmt.Printf("'%s'\n", m["second"])
	fmt.Printf("'%s'\n", m["third"])
	fmt.Printf("'%s'\n", m["xyzzy"])
}

func map5() {
	m1 := make(map[string]string)

	m1["one"] = "I"
	m1["two"] = "II"

	value, exist := m1["two"]
	if exist {
		fmt.Println("found:", value)
	} else {
		fmt.Println("not found")
	}

	delete(m1, "two")

	value, exist = m1["two"]
	if exist {
		fmt.Println("found:", value)
	} else {
		fmt.Println("not found")
	}
}

func map6() {
	type Key struct {
		id   uint32
		role string
	}

	m1 := make(map[Key]User)
	fmt.Println(m1)

	m1[Key{1, "root"}] = User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds"}

	m1[Key{2, "gopher"}] = User{
		id:      2,
		name:    "Rob",
		surname: "Pike"}

	fmt.Println(m1)

	delete(m1, Key{1, "root"})
	fmt.Println(m1)

	delete(m1, Key{1000, "somebody else"})
	fmt.Println(m1)
}
