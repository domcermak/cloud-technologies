package main

import (
	"fmt"

	"domcermak/ctc/assignments/03/cmd/server"
)

func main() {
	fmt.Println("starting...")

	flags := server.ParseFlags()
	if flags.LogLevel == server.Debug {
		fmt.Println("starting with flags:")
		fmt.Println(flags)
	}

	postgres, err := server.NewPostgres(
		flags.PgHost,
		uint16(flags.PgPort),
		flags.PgDatabase,
		flags.PgUsername,
		flags.PgPassword,
	)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Println("listening to incoming connections...")
	if err := server.NewServer(flags.ServerAddr, postgres).ListenAndServe(); err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Println("quitting...")
}
