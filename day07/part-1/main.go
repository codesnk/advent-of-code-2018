package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open input.txt :", err)
	}
	defer file.Close()

}
