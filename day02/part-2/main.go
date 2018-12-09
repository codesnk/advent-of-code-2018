package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the file
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open the input file: ", err)
	}
	defer input.Close()

	// Set a scanner
	scanner := bufio.NewScanner(input)
	var tuples []string

	for scanner.Scan() {
		tuples = append(tuples, scanner.Text())
	}

	var word1, word2 string
	for i, a := range tuples {
		for j, b := range tuples {
			dif := 0
			for x := 0; x < len(tuples[i]); x++ {
				if a[x] != b[x] {
					dif++
				}
			}
			if dif == 1 {
				word1, word2 = tuples[i], tuples[j]
				answer := ""
				for z := 0; z < len(word1); z++ {
					if word1[z] == word2[z] {
						answer += string(word1[z])
					}
				}
				fmt.Println("The word is: ", answer)
				os.Exit(0)
			}
		}
	}
}
