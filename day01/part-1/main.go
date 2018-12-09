package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open the input file: ", err)
	}
	defer inputFile.Close()

	// Read input and calcualte result
	lineScanner := bufio.NewScanner(inputFile)

	// Result will store the final result
	var result int

	// Convert string input to int and add to result
	for lineScanner.Scan() {
		num, err := strconv.Atoi(lineScanner.Text())
		if err != nil {
			log.Fatal("Failed to read the input as integer: ", err)
		}
		result += num
	}

	// Check if we encountered any errors with scanning
	if err := lineScanner.Err(); err != nil {
		log.Fatal("Error in reading the file: ", err)
	}
	// Print the final result
	fmt.Println("The resulting frequency is: ", result)
}
