package main

import (
	"fmt"
	"os"
	"strings"

	"io.witcher.self-help/Functions"
)

func main() {
	args := os.Args

	if len(args) == 0 {
		fmt.Println("Please provide a string to search for")
		return
	}

	if strings.ToLower(args[1]) == "-mk-ins-str" {
		readpath := args[2]
		writepath := args[3]
		fmt.Printf("Reading from %s and write to %s", readpath, writepath)
		err := Functions.CreateMutilpleInsertStringFromFile(readpath, writepath)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Create insert string success")
		}
	}

}
