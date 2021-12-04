package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("6/input.txt")
	scanner := bufio.NewScanner(f)
	yes := make(map[uint8]bool)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// count the number of questions to which anyone answered yes
			sum += len(yes)
			yes = make(map[uint8]bool)
		}
		for i := 0; i < len(line); i++ {
			yes[line[i]] = true
		}
	}
	sum += len(yes)
	fmt.Println(sum)
}
