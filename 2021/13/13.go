package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Fold struct {
	axis string
	pos  int
}

func GetFold(s string) Fold {
	p := strings.Split(s[11:], "=")
	d, _ := strconv.Atoi(p[1])
	return Fold{
		axis: p[0],
		pos:  d,
	}
}

type DotMap map[int]map[int]bool

func ReadInput(r io.ReadSeeker) (DotMap, []Fold) {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	result, folds := make(DotMap, 0), make([]Fold, 0)

	for s.Scan() {
		s := s.Text()
		if strings.HasPrefix(s, "fold along ") {
			folds = append(folds, GetFold(s))
		} else if s != "" {
			parts := strings.Split(s, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			if _, ok := result[y]; !ok {
				result[y] = make(map[int]bool)
			}
			result[y][x] = true
		}
	}
	return result, folds
}

func DoFolds(n int, dots DotMap, folds []Fold) (int, DotMap) {
	for i := 0; i < n; i++ {
		f := folds[i]
		newdots := make(DotMap)
		for y, yy := range dots {
			for x, xx := range yy {
				if f.axis == "x" {
					if _, ok := newdots[y]; !ok {
						newdots[y] = make(map[int]bool)
					}
					if xx && x > f.pos {
						newdots[y][2*f.pos-x] = true
					} else {
						newdots[y][x] = true
					}
				} else {
					if xx && y > f.pos {
						if _, ok := newdots[2*f.pos-y]; !ok && f.axis == "y" {
							newdots[2*f.pos-y] = make(map[int]bool)
						}
						newdots[2*f.pos-y][x] = true
					} else {
						if _, ok := newdots[y]; !ok {
							newdots[y] = make(map[int]bool)
						}
						newdots[y][x] = true
					}
				}
			}
		}
		dots = newdots
	}

	count := 0
	for _, yy := range dots {
		for _, xx := range yy {
			if xx {
				count++
			}
		}
	}

	return count, dots
}

func Part1(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	a, b := ReadInput(r)
	n, _ := DoFolds(1, a, b)
	return n
}

func Part2(r io.ReadSeeker) string {
	defer timeTrack(time.Now())
	a, b := ReadInput(r)
	_, m := DoFolds(len(b), a, b)
	return DrawDotMap(m)
}

func DrawDotMap(m DotMap) string {
	maxy := 0
	maxx := 0
	for y, yy := range m {
		if y > maxy {
			maxy = y
		}
		for x := range yy {
			if x > maxx {
				maxx = x
			}
		}
	}

	lines := make([]string, maxy+1, maxy+1)

	for y, yy := range m {
		line := make([]rune, maxx+1, maxx+1)
		for i := range line {
			line[i] = '.'
		}
		for x := range yy {
			line[x] = '█'
		}
		lines[y] = string(line)
	}
	return "\n" + strings.Join(lines, "\n")
}

func timeTrack(start time.Time) {
	fmt.Printf("(%10s) ", time.Since(start))
}

func expect(expected, actual int, msg string) {
	if expected == actual {
		fmt.Printf("\033[1;32mOK\033[0m")
	} else {
		fmt.Printf("\033[1;31mFAIL\033[0m (expected %d, actual %d)", expected, actual)
	}
	fmt.Println(" ", msg)
}

func main() {
	test1 := strings.NewReader(`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
`)
	input, err := os.Open("2021/13/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	expect(17, Part1(test1), "Part1 - test1")
	fmt.Println("Part1 - puzzle", Part1(input))

	fmt.Println("Part2 - test1", Part2(test1))
	fmt.Println("Part2 - puzzle", Part2(input))
}
