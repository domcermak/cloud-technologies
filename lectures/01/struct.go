package main

import "fmt"

//- a.k.a. records
//- user-defined data type (so called named structure)
//- or anonymous structure
//- dot operator to access struct members
//- initialization using {}
//- or by using named members (which is more explicit and better)
//- structs are comparable
//- pass to functions as value or via pointer (by reference)
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	//struct1()
	//struct2()
	//struct3()
}

func struct1() {
	user1 := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds",
	}

	fmt.Println(user1)

	user1.id = 2
	user1.name = "Rob"
	user1.surname = "Pike"

	fmt.Println(user1)

	user2 := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds"}

	fmt.Println(user1 == user1)
	fmt.Println(user1 == user2)

	printUserByValue(user1)
	printUserByPointer(&user1)
}

func printUserByValue(u User) {
	fmt.Println(u)
}

func printUserByPointer(u *User) {
	fmt.Println(u)
}

func struct2() {
	user1 := &User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds",
	}

	printUserByPointer(user1)
	// dereference
	printUserByValue(*user1)
}

// anonymous struct
func struct3() {
	var employee struct {
		firstName, lastName string
		age                 int
	}

	fmt.Println(employee)
	printAnonymousUser(employee)
}

func printAnonymousUser(u struct {
	firstName, lastName string
	age                 int
}) {
	fmt.Println(u)
}
