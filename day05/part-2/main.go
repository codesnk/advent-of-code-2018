package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var inputText string

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open input.txt file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputText = scanner.Text()
	tmp := ""
	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	truth := map[rune]int{}
	for _, ch := range alphabets {
		tmp = strings.Replace(inputText, string(ch), "", -1)
		tmp = strings.Replace(tmp, string(ch+32), "", -1)
		truth[ch] = reduceText(tmp)
	}

	lowest := -1
	for a := range truth {
		if lowest == -1 {
			lowest = truth[a]
		}
		if lowest > truth[a] {
			lowest = truth[a]
		}
	}
	fmt.Println("The lowest count is: ", lowest)
}

func reduceText(input string) int {
	repeat := true
	change := false
	for repeat {
		repeat = false
		change = false
		for idx, x := range input {
			if (idx + 1) < len(input) {
				if x+32 == rune(input[idx+1]) {
					pair := string(x) + string(x+32)
					input = strings.Replace(input, pair, "", 1)
					change = true
					continue
				}
				if x-32 == rune(input[idx+1]) {
					pair := string(x) + string(x-32)
					input = strings.Replace(input, pair, "", 1)
					change = true
				}
			}
		}
		if change == true {
			repeat = true
		}

	}
	return len(input)
}
