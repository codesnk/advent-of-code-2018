package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open input.txt file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	repeat := true
	scanner.Scan()
	txt := scanner.Text()
	change := false
	for repeat {
		repeat = false
		change = false
		for idx, x := range txt {
			if (idx + 1) < len(txt) {
				if x+32 == rune(txt[idx+1]) {
					pair := string(x) + string(x+32)
					txt = strings.Replace(txt, pair, "", 1)
					change = true
					continue
				}
				if x-32 == rune(txt[idx+1]) {
					pair := string(x) + string(x-32)
					txt = strings.Replace(txt, pair, "", 1)
					change = true
				}
			}
		}
		if change == true {
			repeat = true
		}

	}
	fmt.Println(len(txt))
}
