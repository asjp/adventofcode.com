package main

import (
	"bufio"
	"fmt"
	"os"
)

type Slope struct {
	x, trees, right, down int
}

func main() {
	f, _ := os.Open("3/input.txt")
	scanner := bufio.NewScanner(f)

	slopes := []Slope{{right: 1, down: 1}, {right: 3, down: 1}, {right:5, down: 1}, {right:7, down: 1}, {right:1, down: 2}}

	n := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, _ := range slopes {
			s := &slopes[i]
			if n % s.down == 0 {
				if line[s.x%len(line)] == uint8('#') {
					s.trees++
				}
				s.x += s.right
			}
		}
		n++
	}
	fmt.Println(slopes)
	n = 1
	for _, s := range slopes {
		n = n * s.trees
	}
	fmt.Println(n)
}
