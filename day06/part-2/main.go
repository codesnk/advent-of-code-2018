package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type record struct {
	X, Y float64
}

var instances []record

func main() {
	readFile()
	findAnswer()
}

func readFile() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not read input.txt file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	n := 0
	for scanner.Scan() {
		line := scanner.Text()
		rec := record{}

		if _, err := fmt.Sscanf(line, "%b, %b", &rec.X, &rec.Y); err != nil {
			log.Fatal("Could not parse coordinates: ", err)
		}
		n++
		instances = append(instances, rec)
	}
}

func findAnswer() {
	maxHeight, maxWidth := float64(0), float64(0)
	for _, input := range instances {
		if input.Y > maxHeight {
			maxHeight = input.Y
		}
		if input.X > maxWidth {
			maxWidth = input.X
		}
	}

	maxSum := float64(10000)
	regions := 0
	truthMap := make(map[record]bool)
	tmp := make(map[record]int)
	for y := float64(0); y < maxHeight; y++ {
		for x := float64(0); x < maxWidth; x++ { // start at 0,0
			mc := record{0, 0}
			min := float64(-1)
			totalDistance := float64(0)
			for _, c := range instances {
				dist := math.Abs(x-c.X) + math.Abs(y-c.Y)
				if dist < min || min == -1 {
					min = dist
					mc = c
				} else if dist == min {
					mc = record{-1, -1}
				}
				totalDistance += math.Abs(x-c.X) + math.Abs(y-c.Y)
			}

			if x == 0 || y == 0 || x == maxWidth || y == maxHeight {
				truthMap[mc] = true
			}

			tmp[mc]++

			if totalDistance < maxSum {
				regions++
			}
		}
	}

	answer := 0
	for k, v := range tmp {
		if _, found := truthMap[k]; v > answer && !found {
			answer = v
		}
	}
	fmt.Println("Answer is: ", answer, regions)
}
