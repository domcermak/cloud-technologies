package main

import (
	"fmt"

	"domcermak/ctc/assignments/05/client"
)

func main() {
	fmt.Println("starting...")

	flags := client.ParseFlags()
	fmt.Println(flags)

	c, err := client.NewClient(flags.ServerAddr)
	if err != nil {
		panic(err)
	}

	res, err := c.Post("foo", "bar")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	res, err = c.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	res, err = c.Post("foo", "something_else")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	res, err = c.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	res, err = c.Delete("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	res, err = c.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
