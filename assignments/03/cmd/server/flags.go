package server

import (
	"flag"
)

type Flags struct {
	ServerAddr, PgHost, PgDatabase, PgUsername, PgPassword string
	PgPort                                                 int
}

func ParseFlags() *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.ServerAddr, "server_address", "localhost:1234", "Address for the server to be published on")
	flag.StringVar(&flags.PgHost, "pg_host", "localhost:5432", "Host of running Postgres instance")
	flag.IntVar(&flags.PgPort, "pg_port", 5432, "Port of running Postgres instance")
	flag.StringVar(&flags.PgDatabase, "pg_database", "postgres", "Database name in running Postgres instance")
	flag.StringVar(&flags.PgUsername, "pg_username", "postgres", "Username to running Postgres instance")
	flag.StringVar(&flags.PgPassword, "pg_password", "postgres", "Password to running Postgres instance")
	flag.Parse()

	return flags
}
