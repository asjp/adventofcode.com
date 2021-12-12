package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Diff struct {
	x, y int
}

func increase(height [][]int, x, y int) int {
	diffs := []Diff{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, {-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	flashes := 0
	height[y][x]++
	if height[y][x] == 10 {
		flashes++
		for _, d := range diffs {
			x2 := x + d.x
			y2 := y + d.y
			if x2 < 0 || y2 < 0 || y2 >= len(height) || x2 >= len(height[y2]) {
				continue
			}
			flashes += increase(height, x2, y2)
		}
	}
	return flashes
}

func Calc(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	height := make([][]int, 0, 0)
	for s.Scan() {
		s := s.Text()
		line := make([]int, 0, len(s))
		for _, c := range s {
			h, _ := strconv.Atoi(string(c))
			line = append(line, h)
		}
		height = append(height, line)
	}

	flashes := 0
	for step := 0; step < 100; step++ {
		//fmt.Println(height)
		for y, yy := range height {
			for x := range yy {
				flashes += increase(height, x, y)
			}
		}
		for y, yy := range height {
			for x, xx := range yy {
				if xx >= 10 {
					height[y][x] = 0
				}
			}
		}
	}
	return flashes
}

func Calc2(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	height := make([][]int, 0, 0)
	for s.Scan() {
		s := s.Text()
		line := make([]int, 0, len(s))
		for _, c := range s {
			h, _ := strconv.Atoi(string(c))
			line = append(line, h)
		}
		height = append(height, line)
	}

	flashes := 0
	step := 1
	for {
		//fmt.Println(height)
		for y, yy := range height {
			for x := range yy {
				flashes += increase(height, x, y)
			}
		}
		allflash := true
		for y, yy := range height {
			for x, xx := range yy {
				if xx >= 10 {
					height[y][x] = 0
				} else {
					allflash = false
				}
			}
		}
		if allflash {
			break
		}
		step++
	}
	return step
}

func Part1(r io.ReadSeeker) int {
	return Calc(r)
}

func Part2(r io.ReadSeeker) int {
	return Calc2(r)
}

func main() {
	test1 := strings.NewReader(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`)
	fmt.Println("Part1 - test1", Part1(test1))
	fmt.Println("Part2 - test1", Part2(test1))

	input, err := os.Open("2021/11/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))
	fmt.Println("Part2 - puzzle", Part2(input))
}
