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
		var a1, a2, b1, b2 int
		fmt.Sscanf(sack, "%d-%d,%d-%d", &a1, &a2, &b1, &b2)
		if (b1 >= a1 && b2 <= a2) ||
			(a1 >= b1 && a2 <= b2) {
			score++
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
		var a1, a2, b1, b2 int
		fmt.Sscanf(sack, "%d-%d,%d-%d", &a1, &a2, &b1, &b2)
		if (a1 <= b1 && b1 <= a2) ||
			(b1 <= a1 && a1 <= b2) {
			score++
		}
	}
	return score
}

func main() {
	test := strings.NewReader(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

	input := PuzzleInput()
	fmt.Println("Part1", Part1(test), Part1(input))
	fmt.Println("Part2", Part2(test), Part2(input))
}
