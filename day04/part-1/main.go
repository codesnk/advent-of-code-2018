package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// Record stores each record in input.txt
type record struct {
	timestamp  time.Time
	guardID    int
	guardState string
}

// Records implements the sort interface
type records []record

func (a records) Len() int {
	return len(a)
}
func (a records) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a records) Less(i, j int) bool {
	return a[i].timestamp.Before(a[j].timestamp)
}

// instances will store all records in sorted form
var instances records

func main() {
	sortRecords() // Sort all records and store in instances
	findAnswer()  // Find the answer (guard id * sleep time)
}

func sortRecords() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open file.txt :", err)
	}
	defer file.Close()
	// Read the file line by line
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		var rec record // Create a record from each line

		var y, m, d, hr, mn int // Get a timestamp from the line
		n, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d]", &y, &m, &d, &hr, &mn)
		if n < 5 || err != nil {
			log.Fatal("Could not parse the timestamp: ", err)
		}
		rec.timestamp = time.Date(y, time.Month(m), d, hr, mn, 0, 0, time.UTC) // Store the timestamp

		idx := strings.Index(line, "] ") // Set index for further parsing beyond timestamp
		if idx == -1 {
			log.Fatal("Could not parse beyond timestamp: ", err)
		}
		line = line[idx+2:] // Trim the line to Id and state section
		n, _ = fmt.Sscanf(line, "Guard #%d begins shift", &rec.guardID)

		switch {
		case n == 1:
			rec.guardState = "begins shift"
		case line == "falls asleep":
			rec.guardState = "falls asleep"
		case line == "wakes up":
			rec.guardState = "wakes up"
		default:
			log.Fatal("Could not parse guard action: ", err)
		}
		instances = append(instances, rec)
	}
	// Sort all the records by their timestamp
	sort.Sort(instances)
}

func findAnswer() {
	var lazyGuardID, currentGuard, sleepStartTime int
	sleepTimes := map[int]int{}

	// Calculate the total time slept for all Guards
	for _, g := range instances {
		switch g.guardState {
		case "begins shift":
			currentGuard = g.guardID
		case "falls asleep":
			sleepStartTime = g.timestamp.Minute()
		case "wakes up":
			thisNap := g.timestamp.Minute() - sleepStartTime
			sleepTimes[currentGuard] += thisNap
			if sleepTimes[currentGuard] > sleepTimes[lazyGuardID] {
				lazyGuardID = currentGuard
			}
		}
	}

	// Find the most occurent minutes in each of the lazyGuard's naps
	minutes := [60]int{}
	currentGuard = -1
	var sleepyminute int
	for _, g := range instances {
		if g.guardState == "begins shift" { // Get the id from this record
			currentGuard = g.guardID
			continue
		}
		if currentGuard != lazyGuardID { // We are not interested in this guard
			continue
		}
		switch g.guardState {
		case "falls asleep":
			sleepStartTime = g.timestamp.Minute()
		case "wakes up":
			to := g.timestamp.Minute()
			for i := sleepStartTime; i < to; i++ {
				minutes[i]++
				if minutes[i] > minutes[sleepyminute] {
					sleepyminute = i
				}
			}
		}
	}

	fmt.Printf("Answer: %d", lazyGuardID*sleepyminute)
}
