package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	errHandle()
	errWrap()
}

func errHandle() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	// ignore error
	_ = file.Close()
}

var myError = errors.New("my error")

func errWrap() {
	fmt.Println("is myError", errors.Is(myError, myError))

	upd := fmt.Errorf("with context: %w", myError)
	fmt.Println("is myError", errors.Is(upd, myError))

	fmt.Println("unwrap", errors.Unwrap(upd))
	fmt.Println("unwrap root", myError)
}
