package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type runes []rune

func (s runes) Len() int {
	return len(s)
}
func (s runes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s runes) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	input := strings.TrimSpace(string(bytes))
	fmt.Printf("%v", findAnswer(input))
}

func findAnswer(input string) int {
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
			sort.Sort(runes(temp))
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
