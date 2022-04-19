package client

import (
	"encoding/json"
	"flag"
)

type Flags struct {
	ServerAddr string `json:"server_addr"`
}

func ParseFlags() *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.ServerAddr, "server_address", "localhost:8080", "Address of the server")
	flag.Parse()

	return flags
}

func (f Flags) String() string {
	data, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}

	return string(data)
}
