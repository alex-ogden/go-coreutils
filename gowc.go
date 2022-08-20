// A wc alternative written in go

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var countingLines bool
	targets := os.Args[1:]

	if sliceContains(targets, "-l") {
		countingLines = true
		targets = os.Args[2:]
	} else {
		countingLines = false
	}

	for _, target := range targets {
		// For each file we want to print the name of the file, then the number of words
		if strings.ContainsAny(target, "*[]") {
			if countingLines {
				getLineCountReg(target)
			} else {
				getWordCountReg(target)
			}
		} else {
			if countingLines {
				getLineCount(target)
			} else {
				getWordCount(target)
			}
		}
	}
}

func getWordCount(filename string) {
	fileContent, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// wordCount := len(strings.Fields(string(fileContent)))
	wordCount, err := wordCounter(fileContent)
	fmt.Printf("File: %s\nWord Count: %d\n", filename, wordCount)
}

func getWordCountReg(filename string) {
	matches, err := filepath.Glob(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		fileContent, err := os.Open(match)
		if err != nil {
			log.Fatal(err)
		}

		// wordCount := len(strings.Fields(string(fileContent)))
		wordCount, err := wordCounter(fileContent)
		fmt.Printf("File: %s\nWord Count: %d\n", filename, wordCount)
	}
}

func getLineCount(filename string) {
	fileContent, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	lineCount, err := lineCounter(fileContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File: %s\nLine Count: %d\n", filename, lineCount)
}

func getLineCountReg(filename string) {
	matches, err := filepath.Glob(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		fileContent, err := os.Open(match)
		if err != nil {
			log.Fatal(err)
		}

		lineCount, err := lineCounter(fileContent)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File: %s\nLine Count: %d\n", filename, lineCount)
	}
}

// Self implemented lineCounting method using bytes.Count to find newline characters
func lineCounter(r io.Reader) (int, error) {
	// Create a 32k buffer
	buf := make([]byte, 32*1024)
	count := 0
	// Define our end-of-line (newline char in this case)
	lineSep := []byte{'\n'}

	for {
		// c is the number of bytes read into the buffer (number of bytes in the file)
		c, err := r.Read(buf)
		// we the use bytes.Count to count the number of non-overlapping instances of our line seperator
		// in the buffer (:c means every byte from 0 -> length of buffer)
		count += bytes.Count(buf[:c], lineSep)

		// Check for our errors
		switch {
		// If we get the EOF error, we've hit the end of the file - return our count
		case err == io.EOF:
			return count, nil
		// Otherwise there's an issue, return the error!
		case err != nil:
			return count, err
		}
	}
}

func wordCounter(r io.Reader) (int, error) {
	// Create a 32k buffer
	buf := make([]byte, 64*1024)
	count := 0
	// Define our end-of-line (newline char in this case)
	lineSep := []byte{' '}

	for {
		// c is the number of bytes read into the buffer (number of bytes in the file)
		c, err := r.Read(buf)
		// we the use bytes.Count to count the number of non-overlapping instances of our line seperator
		// in the buffer (:c means every byte from 0 -> length of buffer)
		count += bytes.Count(buf[:c], lineSep)

		// Check for our errors
		switch {
		// If we get the EOF error, we've hit the end of the file - return our count
		case err == io.EOF:
			return count, nil
		// Otherwise there's an issue, return the error!
		case err != nil:
			return count, err
		}
	}
}

// Have to write my own contains method for checking for flags
func sliceContains(targetSlice []string, pattern string) bool {
	for _, slice := range targetSlice {
		if slice == pattern {
			return true
		}
	}
	return false
}
