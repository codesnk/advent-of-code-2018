package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open the input.txt file: ", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
}
