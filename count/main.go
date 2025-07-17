package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
)

func main() {
	filename := flag.String("file", "", "path to input text file")
	flag.Parse()

	if *filename == "" {
		fmt.Println("Usage: go run count.go -file <path>")
		os.Exit(1)
	}
	// os.Open opens file and returns *os.File and an error, if it works correctly, err will be nil.
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
}