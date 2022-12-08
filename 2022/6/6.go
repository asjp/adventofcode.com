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

func Part1and2(r io.ReadSeeker, msgSize int) int {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	scan.Scan()
	line := scan.Text()
	var pos int
	for pos = 1; pos <= len(line); pos++ {
		if pos < msgSize {
			continue
		}
		freq := map[byte]int{}
		for i := 1; i <= msgSize; i++ {
			freq[line[pos-i]]++
		}
		if len(freq) == msgSize {
			break
		}
	}
	return pos
}

func Part1(r io.ReadSeeker) int {
	return Part1and2(r, 4)
}

func Part2(r io.ReadSeeker) int {
	return Part1and2(r, 14)
}

func main() {
	test := strings.NewReader(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`)

	input := PuzzleInput()
	fmt.Println("Part1", Part1(test), Part1(input))
	fmt.Println("Part2", Part2(test), Part2(input))
}
