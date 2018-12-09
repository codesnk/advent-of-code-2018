package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var history map[int]bool
var done bool
var result int

func main() {
	done = false
	history = make(map[int]bool)
	history[0] = true
	for !done {
		readFile()
	}
}

func readFile() {
	// Open the file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open the input file: ", err)
	}
	defer inputFile.Close()

	lineScanner := bufio.NewScanner(inputFile)

	for lineScanner.Scan() {
		num, err := strconv.Atoi(lineScanner.Text()) // Probably not the most safe way to do this.
		if err != nil {
			log.Fatal("Failed to read the input as integer: ", err)
		}
		result += num
		// Check if the resulting frequency has been seen before
		if history[result] {
			fmt.Println("The first frequency that repeats is: ", result)
			done = true
			return
		}
		// Otherwise add to the history list
		history[result] = true
	}
}
