package main

import "fmt"

func main() {
	//pointer1()
	//pointer2()
	//pointer3()
	//pointer4()
	//pointer5()
}

func pointer1() {
	var i int = 42

	fmt.Println(i)

	var ptrToInteger *int
	fmt.Println(ptrToInteger)

	ptrToInteger = &i

	fmt.Println(ptrToInteger)
	fmt.Println(*ptrToInteger)

	*ptrToInteger++

	fmt.Println(i)
	fmt.Println(*ptrToInteger)
}

func pointer2() {
	var u User

	u.id = 1
	u.name = "Linus"
	u.surname = "Torvalds"

	fmt.Println(u)

	var pUser *User
	pUser = &u

	fmt.Println(pUser)
	fmt.Println(*pUser)

	(*pUser).id = 10000
	fmt.Println(*pUser)

	pUser.id = 20000
	fmt.Println(*pUser)
}

func pointer3() {
	var u User
	u.id = 1
	u.name = "Linus"
	u.surname = "Torvalds"

	fmt.Println(u)
	fmt.Println("------------------")

	var pName *string
	var pSurname *string
	pName = &u.name
	pSurname = &u.surname

	fmt.Println(pName)
	fmt.Println(pSurname)
	fmt.Println(*pName)
	fmt.Println(*pSurname)
	fmt.Println("------------------")

	(*pName) = "Rob"
	(*pSurname) = "Pike"
	fmt.Println(*pName)
	fmt.Println(*pSurname)
}

func getPointer() *int {
	var i int = 42
	return &i
}

func pointer4() {
	p := getPointer()
	fmt.Printf("%v\n", p)
	fmt.Printf("%d\n", *p)
}

func pointer5() {
	month := [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December"}

	fmt.Println(month)

	// cont
	pThird := &month[2]
	*pThird = "Březen"

	fmt.Println(month)
}
