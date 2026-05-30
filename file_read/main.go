package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Please provide a file path!")
		os.Exit(1)
	}

	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w, er := io.Copy(os.Stdout, file)

	if w == 0 || er != nil {
		fmt.Println("File is empty!")
	}
}
