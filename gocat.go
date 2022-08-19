// A cat alternative written in go

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get a string slice of filenames
	targets := os.Args[1:]

	for _, file := range targets {
		// Check if file provided is regex or not
		if strings.ContainsAny(file, "*[]") {
			// Regex
			outputFileReg(file)
		} else {
			// Not regex
			outputFile(file)
		}
	}
}

func outputFile(filename string) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Convert byte slice to string and print
	fmt.Printf("%s\n", string(fileContent))
}

func outputFileReg(filename string) {
	files, err := filepath.Glob(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileContent, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", string(fileContent))
	}
}
