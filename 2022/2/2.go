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

func Part1and2(r io.ReadSeeker, rules map[string]int) int {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	score := 0
	for scan.Scan() {
		score += rules[scan.Text()]
	}
	return score
}

func main() {
	rules1 := map[string]int{
		// wins
		"C X": 7,
		"A Y": 8,
		"B Z": 9,
		// loses
		"B X": 1,
		"C Y": 2,
		"A Z": 3,
		// draws
		"A X": 4,
		"B Y": 5,
		"C Z": 6,
	}

	rules2 := map[string]int{
		// wins
		"A Z": 8,
		"B Z": 9,
		"C Z": 7,
		// loses
		"A X": 3,
		"B X": 1,
		"C X": 2,
		// draws
		"A Y": 4,
		"B Y": 5,
		"C Y": 6,
	}

	test := strings.NewReader(`A Y
B X
C Z`)

	input := PuzzleInput()
	fmt.Println("Part1", Part1and2(test, rules1), Part1and2(input, rules1))
	fmt.Println("Part2", Part1and2(test, rules2), Part1and2(input, rules2))
}
