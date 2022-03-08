package main

import (
	"domcermak/ctc/assignments/03/cmd/server"
	"fmt"
)

func main() {
	fmt.Println("starting...")
	if err := server.NewServer().ListenAndServe(); err != nil {
		_ = fmt.Errorf("err: %v", err)
	}
	fmt.Println("quitting...")
}
