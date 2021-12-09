package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func isLow(height [][]int, x, y int) bool {
	_, _, h := lowestAround(height, x, y)
	return height[y][x] <= h
}

type Diff struct {
	x, y int
}

func lowestAround(height [][]int, x, y int) (int, int, int) {
	l := height[y][x]
	lx := x
	ly := y
	diffs := []Diff{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, d := range diffs {
		x2 := x + d.x
		y2 := y + d.y
		if x2 < 0 || y2 < 0 || y2 >= len(height) || x2 >= len(height[y2]) {
			continue
		}
		if height[y2][x2] < l {
			l = height[y2][x2]
			lx = x2
			ly = y2
		}
	}
	return lx, ly, l
}

func lowestFrom(height [][]int, x, y int) (int, int, int) {
	lx, ly, l := lowestAround(height, x, y)
	if l < height[y][x] {
		lx, ly, l = lowestFrom(height, lx, ly)
	}
	return lx, ly, l
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

	lowPoints := make(map[string]int)
	for y, yy := range height {
		for x, xx := range yy {
			if xx != 9 && isLow(height, x, y) {
				lowPoints[fmt.Sprintf("%d,%d", x, y)] = xx
			}
		}
	}
	//fmt.Println(lowPoints)

	sum := 0
	for _, h := range lowPoints {
		sum += (h + 1)
	}
	return sum
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

	basins := make(map[string]int)
	for y, yy := range height {
		for x, xx := range yy {
			if xx == 9 {
				continue
			}
			bx, by, _ := lowestFrom(height, x, y)
			//fmt.Printf("%d,%d -> %d,%d\n", y, x, by, bx)
			basinkey := fmt.Sprintf("%d,%d", by, bx)
			if bp, ok := basins[basinkey]; !ok {
				basins[basinkey] = 1
			} else {
				basins[basinkey] = bp + 1
			}
		}
	}
	//fmt.Println(basins)
	var topbasins []int
	for _, sz := range basins {
		topbasins = append(topbasins, sz)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(topbasins)))

	return topbasins[0] * topbasins[1] * topbasins[2]
}

func Part1(r io.ReadSeeker) int {
	return Calc(r)
}

func Part2(r io.ReadSeeker) int {
	return Calc2(r)
}

func main() {
	test1 := strings.NewReader(`2199943210
3987894921
9856789892
8767896789
9899965678
`)
	fmt.Println("Part1 - test1", Part1(test1))
	fmt.Println("Part2 - test1", Part2(test1))

	input, err := os.Open("2021/9/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))
	fmt.Println("Part2 - puzzle", Part2(input))
}
