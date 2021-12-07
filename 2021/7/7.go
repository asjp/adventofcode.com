package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Calc(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	s.Scan()
	start := strings.Split(s.Text(), ",")
	positions := make([]int, 0)
	for _, s := range start {
		i, _ := strconv.Atoi(s)
		positions = append(positions, i)
	}
	sort.Ints(positions)
	median := positions[len(positions)/2-1]
	fmt.Println("median", median)
	fuel := 0
	for _, p := range positions {
		if p < median {
			fuel += median - p
		} else if p > median {
			fuel += p - median
		}
	}
	return fuel
}

func fuel(distance int) int {
	sum := 0
	for i := 1; i <= distance; i++ {
		sum += i
	}
	return sum
}

func Calc2(r io.ReadSeeker, rangeMedian int) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	s.Scan()
	start := strings.Split(s.Text(), ",")
	positions := make([]int, 0)
	for _, s := range start {
		i, _ := strconv.Atoi(s)
		positions = append(positions, i)
	}
	sort.Ints(positions)
	median := positions[len(positions)/2-1] - rangeMedian/2
	min := math.MaxInt64
	for offset := 0; offset < rangeMedian; offset++ {
		median++
		f := 0
		for _, p := range positions {
			if p < median {
				f += fuel(median - p)
			} else if p > median {
				f += fuel(p - median)
			}
		}
		if f < min {
			min = f
		}
	}
	return min
}

func Part1(r io.ReadSeeker) int {
	return Calc(r)
}

func Part2(r io.ReadSeeker) int {
	return Calc2(r, 1000)
}

func main() {
	test := strings.NewReader("16,1,2,0,4,2,7,1,2,14")
	fmt.Println("Part1 - test", Part1(test))
	fmt.Println("Part2 - test", Part2(test))

	input, err := os.Open("2021/7/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))
	fmt.Println("Part2 - puzzle", Part2(input))
}
