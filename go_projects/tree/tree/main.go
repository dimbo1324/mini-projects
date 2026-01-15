package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
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
	return walk(out, path, printFiles, "")
}

func walk(out io.Writer, path string, printFiles bool, prefix string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var filtered []os.DirEntry
	for _, entry := range entries {
		if entry.Name() == ".DS_Store" {
			continue
		}
		if !printFiles && !entry.IsDir() {
			continue
		}
		filtered = append(filtered, entry)
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name() < filtered[j].Name()
	})

	for i, entry := range filtered {
		isLast := i == len(filtered)-1

		var connector string
		if isLast {
			connector = endSign
		} else {
			connector = dirSign
		}

		var sizeInfo string
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return err
			}
			if info.Size() == 0 {
				sizeInfo = " (empty)"
			} else {
				sizeInfo = fmt.Sprintf(" (%db)", info.Size())
			}
		}

		_, err := fmt.Fprintf(out, "%s%s%s%s\n", prefix, connector, entry.Name(), sizeInfo)
		if err != nil {
			return err
		}

		if entry.IsDir() {
			newPrefix := prefix
			if isLast {
				newPrefix += "\t"
			} else {
				newPrefix += pipeSign
			}

			subPath := filepath.Join(path, entry.Name())

			err := walk(out, subPath, printFiles, newPrefix)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
