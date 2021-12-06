package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Calc(r io.ReadSeeker, iterations int) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	s.Scan()
	start := strings.Split(s.Text(), ",")
	days := make([]int, 9, 9)
	for _, s := range start {
		i, _ := strconv.Atoi(s)
		days[i] += 1
	}
	var t int
	for i := 0; i < iterations; i++ {
		t = days[0]
		for j := 0; j < 8; j++ {
			days[j] = days[j+1]
		}
		days[6] += t
		days[8] = t
	}
	count := 0
	for _, i := range days {
		count += i
	}
	return count
}

func Part1(r io.ReadSeeker) int {
	return Calc(r, 80)
}

func Part2(r io.ReadSeeker) int {
	return Calc(r, 256)
}

func main() {
	test := strings.NewReader("3,4,3,1,2")
	fmt.Println("Part1 - test", Part1(test))
	fmt.Println("Part2 - test", Part2(test))

	input, err := os.Open("2021/6/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))
	fmt.Println("Part2 - puzzle", Part2(input))
}
