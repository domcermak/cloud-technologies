package client

import (
	"flag"
	"time"
)

type Flags struct {
	Timeout    int64
	ServerAddr string
}

func ParseFlags() *Flags {
	flags := &Flags{}
	flag.Int64Var(&flags.Timeout, "timeout", 10, "Connection timeout in milliseconds")
	flag.StringVar(&flags.ServerAddr, "server_address", "localhost:1234", "Address of the server the client is connecting to")
	flags.Timeout *= int64(time.Second)

	return flags
}
