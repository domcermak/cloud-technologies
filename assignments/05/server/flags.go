package server

import (
	"encoding/json"
	"flag"
)

type Flags struct {
	EtcdServerAddr string
	ServerAddr     string
}

func ParseFlags() *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.EtcdServerAddr, "etcd_server_address", "localhost:2379", "Address of the etcd server")
	flag.StringVar(&flags.ServerAddr, "server_address", ":8080", "Address of the server")
	flag.Parse()

	return flags
}

func (f Flags) String() string {
	data, err := json.MarshalIndent(f, "", "\t")
	if err != nil {
		panic(err)
	}

	return string(data)
}
