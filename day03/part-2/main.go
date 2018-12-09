package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Failed to open input.txt file: ", err)
	}
	defer file.Close()

	// Read the input with a scanner
	scanner := bufio.NewScanner(file)
	type coord struct {
		x int
		y int
	}

	fabric := map[coord]int{} // Will keep a track of every coordinate

	for scanner.Scan() { // Fill the coordinates
		var id, xx, yy, w, h int
		_, err := fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &xx, &yy, &w, &h)
		if err != nil {
			log.Fatal("Failed to parse input: ", err)
		}
		for a := 0; a < w; a++ {
			for b := 0; b < h; b++ {
				tmp := coord{x: (a + xx), y: (b + yy)}
				fabric[tmp]++
			}
		}
	}

	file.Seek(0, 0) // Rest file prointer
	scanner2 := bufio.NewScanner(file)
	for scanner2.Scan() {
		var id, xx, yy, w, h int
		_, err := fmt.Sscanf(scanner2.Text(), "#%d @ %d,%d: %dx%d", &id, &xx, &yy, &w, &h)
		if err != nil {
			log.Fatal("Failed to parse input: ", err)
		}
		found := true
		for a := 0; a < w; a++ {
			for b := 0; b < h; b++ {
				tmp := coord{x: (a + xx), y: (b + yy)}
				if fabric[tmp] > 1 {
					found = false
				}
			}
		}

		if found == true {
			fmt.Println("The id is: ", id)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error in reading the input file. ", err)
	}
}
