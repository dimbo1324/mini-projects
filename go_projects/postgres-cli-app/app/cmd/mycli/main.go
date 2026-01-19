package main

import (
	"flag"
)

var flags struct {
	ConfigPath    string
	DirectoryPath string
}

func init() {
	flag.StringVar(&flags.ConfigPath, "c", "", "path to yaml schema config files")
	flag.StringVar(&flags.DirectoryPath, "d", "", "path to directory with files")
}

const postgresqlConnString = "postgres://user:password@localhost:5432/postgres"

func main() {
	// Инициализация
}
