package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

const (
	endSign  = "└───"
	dirSign  = "├───"
	pipeSign = "│\t"
)

var printFiles = flag.Bool("f", false, "print files")

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("usage: go run main.go [-f] [PATH]")
	}
	path := args[0]

	if err := dirTree(os.Stdout, path, *printFiles); err != nil {
		log.Fatalln("error:", err)
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	// TODO: ваша реализация
	return errors.New("не реализовано")
}
