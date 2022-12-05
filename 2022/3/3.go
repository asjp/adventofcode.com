package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
)

func PuzzleInput() io.ReadSeeker {
	_, filename, _, _ := runtime.Caller(0)
	today := path.Dir(filename)
	f, err := os.Open(path.Join(today, "input.txt"))
	if err != nil {
		fmt.Println(err)
	}
	return f
}

func Part1(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	score := 0
	for scan.Scan() {
		sack := scan.Text()
		c1 := sack[:len(sack)/2]
		c2 := sack[len(sack)/2:]
		n := map[rune]int{}
		for _, r := range c1 {
			n[r] = 1
		}
		for _, r := range c2 {
			if _, ok := n[r]; ok {
				if r >= 'a' {
					score += 1 + int(r-'a')
				} else {
					score += 27 + int(r-'A')
				}
				break
			}
		}
	}
	return score
}

func Part2(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	score := 0
	for scan.Scan() {
		sack := scan.Text()
		n := map[rune]int{}
		for _, r := range sack {
			n[r] = 1
		}
		scan.Scan()
		sack = scan.Text()
		for _, r := range sack {
			if n[r] == 1 {
				n[r] = 2
			}
		}
		scan.Scan()
		sack = scan.Text()
		for _, r := range sack {
			if n[r] == 2 {
				if r >= 'a' {
					score += 1 + int(r-'a')
				} else {
					score += 27 + int(r-'A')
				}
				break
			}
		}
	}
	return score
}

func main() {
	test := strings.NewReader(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`)

	input := PuzzleInput()
	fmt.Println("Part1", Part1(test), Part1(input))
	fmt.Println("Part2", Part2(test), Part2(input))
}
