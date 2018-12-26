package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Runes []rune

func (s Runes) Len() int {
	return len(s)
}
func (s Runes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Runes) Less(i, j int) bool {
	return s[i] < s[j]
}

func getInput() (input string) {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	input = strings.TrimSpace(string(bytes))
	return
}

func main() {
	input := getInput()
	fmt.Printf("Solution to part 1 is: %v\n", findAnswer(input))
	//fmt.Printf("Solution to part 1 is: %v\n", aoc02(input))
}

func findAnswer(input string) string {
	records := make(map[string][]string)
	parents := make(map[string]int)
	candidateStack := make([]string, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var k, v string
		fmt.Sscanf(scanner.Text(), "Step %v must be finished before step %v can begin.", &k, &v)
		records[k] = append(records[k], v) // compiles list of children for each rune
		parents[v] = parents[v] + 1        // keep track of parents
	}

	for k := range records {
		if parents[k] == 0 { // If a node has no parents, its a root
			candidateStack = append(candidateStack, k)
		}
	}
	var ans string

	for len(candidateStack) > 0 {
		tmp := candidateStack                      // RJ
		sort.Strings(tmp)                          // JR
		par := tmp[0]                              // J
		for i := 0; i < len(candidateStack); i++ { // because candidateStack is not sorted, we find idx of par and delete it
			if candidateStack[i] == par {
				candidateStack = append(candidateStack[:i], candidateStack[i+1:]...)
			}
		}
		ans = ans + par
		for _, v := range records[par] { // For each record for par, remove it from its children's record
			parents[v] = parents[v] - 1
			if parents[v] == 0 {
				candidateStack = append(candidateStack, v) // If all parents are done, put on candidate stack/finished stack
			}
		}
	}

	return ans
}

func aoc02(input string) int {
	instructions := make(map[rune][]rune)
	nodes := make(map[rune]int)
	sl := strings.Split(input, "\n")
	finished, readyTasks := make([]rune, 0), make([]rune, 0)
	workersTasks := []rune{'.', '.', '.', '.', '.'}
	timeLeft := []int{0, 0, 0, 0, 0}
	t, working := 0, 1

	for _, k := range sl {
		key, value := rune(k[5]), rune(k[36])
		instructions[key] = append(instructions[key], value)
		nodes[value] = nodes[value] + 1
	}

	for k := range instructions {
		if nodes[k] == 0 {
			readyTasks = append(readyTasks, k)
		}
	}

	for ; working > 0; t++ {
		working = 0
		for n := range timeLeft {
			if timeLeft[n] != 0 {
				timeLeft[n] = timeLeft[n] - 1
				working = working + 1
			} else {
				if workersTasks[n] != '.' {
					finishedTask := workersTasks[n]
					workersTasks[n] = '.'
					for _, v := range instructions[finishedTask] {
						nodes[v] = nodes[v] - 1
						if nodes[v] == 0 {
							readyTasks = append(readyTasks, v)
						}
					}
				}
			}
		}

		for len(readyTasks) > 0 && working < len(timeLeft) {
			temp := readyTasks
			sort.Sort(Runes(temp))
			x := temp[0]
			for i := 0; i < len(readyTasks); i++ {
				if readyTasks[i] == x {
					readyTasks = append(readyTasks[:i], readyTasks[i+1:]...)
				}
			}
			finished = append(finished, x)
			for n := range timeLeft {
				if workersTasks[n] == '.' {
					workersTasks[n] = x
					timeLeft[n] = int(x) - 5
					working = working + 1
					break
				}
			}
		}
	}
	return t - 1
}
