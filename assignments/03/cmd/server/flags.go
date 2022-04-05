package server

import (
	"encoding/json"
	"flag"
)

const (
	Debug string = "debug"
	Info         = "info"
)

type Flags struct {
	ServerAddr string `json:"server_addr"`
	PgHost     string `json:"pg_host"`
	PgDatabase string `json:"pg_database"`
	PgUsername string `json:"pg_username"`
	PgPassword string `json:"pg_password"`
	PgPort     int    `json:"pg_port"`
	LogLevel   string `json:"log_level"`
}

func ParseFlags() *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.ServerAddr, "server_address", "localhost:1234", "Address for the server to be published on")
	flag.StringVar(&flags.PgHost, "pg_host", "localhost:5432", "Host of running Postgres instance")
	flag.IntVar(&flags.PgPort, "pg_port", 5432, "Port of running Postgres instance")
	flag.StringVar(&flags.PgDatabase, "pg_database", "postgres", "Database name in running Postgres instance")
	flag.StringVar(&flags.PgUsername, "pg_username", "postgres", "Username to running Postgres instance")
	flag.StringVar(&flags.PgPassword, "pg_password", "postgres", "Password to running Postgres instance")
	flag.StringVar(&flags.LogLevel, "log_level", Info, "Log level of the application")
	flag.Parse()

	return flags
}

func (flags Flags) String() string {
	data, err := json.MarshalIndent(flags, "", "\t")
	if err != nil {
		panic(err)
	}

	return string(data)
}
