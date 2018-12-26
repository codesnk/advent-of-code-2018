package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

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
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Could not open input.txt file: ", err)
	}

	records := make(map[string][]string) // For each node, a slice of its children
	parents := make(map[string]int)      // Records how many parents each node has
	candidateStack := make([]string, 0)  // Holds all available nodes that are to be considered
	var answer string                    // Holds the final answer
	text := strings.TrimSpace(string(file))
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		var k, v string
		fmt.Sscanf(scanner.Text(), "Step %v must be finished before step %v can begin.", &k, &v)
		records[k] = append(records[k], v) // compiles list of children for each parent
		parents[v] = parents[v] + 1        // keep track of parents of each node
	}

	for k := range records {
		if parents[k] == 0 { // If a node has no parents, its a root and should be on the candidateStack at start
			candidateStack = append(candidateStack, k)
		}
	}

	for len(candidateStack) > 0 {
		tmp := candidateStack
		sort.Strings(tmp)
		par := tmp[0]
		for i := 0; i < len(candidateStack); i++ { // because candidateStack is not sorted, we find idx of par and delete it
			if candidateStack[i] == par {
				candidateStack = append(candidateStack[:i], candidateStack[i+1:]...)
			}
		}
		answer = answer + par
		for _, v := range records[par] { // For each record for par, remove it from its children's record
			parents[v] = parents[v] - 1
			if parents[v] == 0 {
				candidateStack = append(candidateStack, v) // If all parents are done, put on candidate stack/finished stack
			}
		}
	}
	fmt.Printf(answer)
}
