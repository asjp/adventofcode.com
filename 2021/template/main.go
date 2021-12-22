package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"
)

type Input struct {
}

func ReadInput(r io.ReadSeeker) Input {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	input := Input{}
	for s.Scan() {
		line := s.Text()
		fmt.Print(line)
		// parse...
	}
	return input
}

func Iterate(input Input) int {
	return 0
}

func Part1(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	return Iterate(ReadInput(r))
}

func Part2(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	return Iterate(ReadInput(r))
}

func timeTrack(start time.Time) {
	fmt.Printf("(%10s) ", time.Since(start))
}

func expect(expected, actual int, msg string) {
	if expected == actual {
		fmt.Printf("\033[1;32mOK\033[0m %d", actual)
	} else {
		fmt.Printf("\033[1;31mFAIL\033[0m (expected %d, actual %d)", expected, actual)
	}
	fmt.Println(" ", msg)
}

func expectStr(expected string, actual fmt.Stringer) {
	if expected == actual.String() {
		fmt.Printf("\033[1;32mOK\033[0m %s\n", actual)
	} else {
		fmt.Printf("\033[1;31mFAIL\033[0m (expected %s, actual %s)\n", expected, actual)
	}
}

func FileReader(name string) io.ReadSeeker {
	_, mypath, _, _ := runtime.Caller(0)
	input, err := os.Open(path.Join(path.Dir(mypath), name))
	if err != nil {
		fmt.Println(err)
	}
	return input
}

func main() {
	test := FileReader("test.txt")
	expect(0, Part1(test), "Part1 - test")

	input := FileReader("input.txt")
	fmt.Println("Part1 - puzzle", Part1(input))

	expect(0, Part2(test), "Part2 - test")
	fmt.Println("Part2 - puzzle", Part2(input))
}
