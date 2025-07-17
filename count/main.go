package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func processFile(file *os.File) {
	scanner := bufio.NewScanner(file)

	var (
		lineCount   int
		wordCount   int
		byteCount   int
		longestLine string
	)

	for scanner.Scan() {
		line := scanner.Text()

		lineCount++
		wordCount += len(splitWords(line))
		byteCount += len(line) + 1 // +1 for newline char

		if len(line) > len(longestLine) {
			longestLine = line
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Lines: %d\n", lineCount)
	fmt.Printf("Words: %d\n", wordCount)
	fmt.Printf("Bytes: %d\n", byteCount)
	fmt.Printf("Longest line: %q\n", longestLine)
}

func splitWords(s string) []string {
	return strings.Fields(s)
}

func main() {
	// create -file flag
	filename := flag.String("file", "", "path to input text file")
	flag.Parse()

	if *filename == "" {
		fmt.Println("Usage: go run ./count -file <path>")
		os.Exit(1)
	}
	// os.Open opens file and returns *os.File and an error, if it works correctly, err will be nil.
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	processFile(file)
	defer file.Close()
}