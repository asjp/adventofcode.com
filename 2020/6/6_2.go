package main

import (
	"bufio"
	"fmt"
	"os"
)

func allYes(yes map[uint8]int, groupsize int) int {
	// count the number of questions to which everyone answered yes
	n := 0
	for _, y := range yes {
		if y == groupsize {
			n++
		}
	}
	return n
}

func main() {
	f, _ := os.Open("6/input.txt")
	scanner := bufio.NewScanner(f)
	yes := make(map[uint8]int)
	sum := 0
	groupsize := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sum += allYes(yes, groupsize)
			yes = make(map[uint8]int)
			groupsize = 0
		} else {
			groupsize++
		}
		for i := 0; i < len(line); i++ {
			if old, ok := yes[line[i]]; ok {
				yes[line[i]] = old + 1
			} else {
				yes[line[i]] = 1
			}
		}
	}
	sum += allYes(yes, groupsize)
	fmt.Println(sum)
}
