package common

import "fmt"

func LogRequest(name string, fn func() (interface{}, error)) (interface{}, error) {
	fmt.Printf("running %s\n", name)
	defer fmt.Printf("finishing %s\n", name)

	return fn()
}
