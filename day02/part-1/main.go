package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputfile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open input file: ", err)
	}
	defer inputfile.Close()

	var twos, threes int // Will keep track of how may two and threes were found for final result
	scanner := bufio.NewScanner(inputfile)

	for scanner.Scan() { // Repeat for each line of input
		temp := map[rune]int{}
		for _, ch := range scanner.Text() {
			temp[ch]++
		}
		var foundTwo, foundThree bool
		for _, x := range temp {
			if x == 2 {
				foundTwo = true
			}
			if x == 3 {
				foundThree = true
			}
		}
		if foundTwo {
			twos++
		}
		if foundThree {
			threes++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Errors encountered while reading the file: ", err)
	}
	fmt.Println("The calculated checksum is: ", twos*threes)
}
