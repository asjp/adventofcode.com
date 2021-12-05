package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

func makeLine(x, y, x2, y2 int) []Coord {
	r := make([]Coord, 0)
	if x <= x2 {
		if y <= y2 {
			for {
				r = append(r, Coord{x, y})
				if x == x2 && y == y2 {
					break
				}
				if x < x2 {
					x++
				}
				if y < y2 {
					y++
				}
			}
		} else {
			for {
				r = append(r, Coord{x, y})
				if x == x2 && y == y2 {
					break
				}
				if x < x2 {
					x++
				}
				if y > y2 {
					y--
				}
			}
		}
	} else {
		if y <= y2 {
			for {
				r = append(r, Coord{x, y})
				if x == x2 && y == y2 {
					break
				}
				if x > x2 {
					x--
				}
				if y < y2 {
					y++
				}
			}
		} else {
			for {
				r = append(r, Coord{x, y})
				if x == x2 && y == y2 {
					break
				}
				if x > x2 {
					x--
				}
				if y > y2 {
					y--
				}
			}
		}
	}
	return r
}

func Calc(r io.ReadSeeker, height int, allowDiag bool) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	maps := make(map[int]map[int]int, 0)
	var (
		x1, y1, x2, y2 int
	)
	for s.Scan() {
		_, _ = fmt.Sscanf(s.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if allowDiag || (x1 == x2 || y1 == y2) {
			line := makeLine(x1, y1, x2, y2)
			for _, p := range line {
				if _, ok := maps[p.x]; !ok {
					maps[p.x] = make(map[int]int, 0)
				}
				if oldv, ok := maps[p.x][p.y]; !ok {
					maps[p.x][p.y] = 1
				} else {
					maps[p.x][p.y] = oldv + 1
				}
				//fmt.Println(p.x, p.y, maps[p.x][p.y])
			}
		}
	}
	count := 0
	for _, x := range maps {
		for _, y := range x {
			if y >= height {
				count++
			}
		}
	}
	return count
}

func Part1(r io.ReadSeeker) int {
	return Calc(r, 2, false)
}

func Part2(r io.ReadSeeker) int {
	return Calc(r, 2, true)
}

func main() {
	test := strings.NewReader(`0,9 -> 5,9
	8,0 -> 0,8
	9,4 -> 3,4
	2,2 -> 2,1
	7,0 -> 7,4
	6,4 -> 2,0
	0,9 -> 2,9
	3,4 -> 1,4
	0,0 -> 8,8
	5,5 -> 8,2
`)
	fmt.Println("Part1 - test", Part1(test))
	fmt.Println("Part2 - test", Part2(test))

	input, err := os.Open("2021/5/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))
	fmt.Println("Part2 - puzzle", Part2(input))
}
