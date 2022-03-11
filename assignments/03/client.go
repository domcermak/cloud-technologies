package main

import (
	"os"
	"time"

	cl "domcermak/ctc/assignments/03/cmd/client"
)

func main() {
	flags := cl.ParseFlags()
	client := cl.NewClient(time.Duration(flags.Timeout), flags.ServerAddr)
	cmd := cl.NewCommandLine(os.Stdin, os.Stdout, os.Stderr, client.CommandExecutioners()...)
	cmd.RenderAndAcceptCommands()
}
