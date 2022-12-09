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

type Grid map[int]map[int]byte
type GridSlice [][]byte

func ParseGrid(r io.ReadSeeker) Grid {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	trees := Grid{}
	y := 0
	for scan.Scan() {
		line := scan.Text()
		trees[y] = map[int]byte{}
		for x := 0; x < len(line); x++ {
			trees[y][x] = line[x]
		}
		y++
	}
	return trees
}

func ParseGridSlice(r io.ReadSeeker) GridSlice {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	trees := GridSlice{}
	y := 0
	for scan.Scan() {
		line := scan.Text()
		trees = append(trees, []byte{})
		for x := 0; x < len(line); x++ {
			trees[y] = append(trees[y], line[x])
		}
		y++
	}
	return trees
}

func Part1(trees Grid) int {
	visible := map[int]map[int]bool{}
	for y := 0; y < len(trees); y++ {
		visible[y] = map[int]bool{}
		min := trees[y][0]
		for x := 0; x < len(trees[y]); x++ {
			if y == 0 || y == len(trees)-1 {
				visible[y][x] = true
				continue
			}
			if trees[y][x] > min {
				visible[y][x] = true
				min = trees[y][x]
			}
		}
		min = trees[y][len(trees[y])-1]
		for x := len(trees[y]) - 1; x >= 0; x-- {
			if y == 0 || y == len(trees)-1 {
				visible[y][x] = true
				continue
			}
			if trees[y][x] > min {
				visible[y][x] = true
				min = trees[y][x]
			}
		}
	}
	for x := 0; x < len(trees); x++ {
		min := trees[0][x]
		for y := 0; y < len(trees[x]); y++ {
			if x == 0 || x == len(trees)-1 {
				visible[y][x] = true
				continue
			}
			if trees[y][x] > min {
				visible[y][x] = true
				min = trees[y][x]
			}
		}
		min = trees[len(trees)-1][x]
		for y := len(trees) - 1; y >= 0; y-- {
			if x == 0 || x == len(trees)-1 {
				visible[y][x] = true
				continue
			}
			if trees[y][x] > min {
				visible[y][x] = true
				min = trees[y][x]
			}
		}
	}

	total := 0
	for _, a := range visible {
		for _, b := range a {
			if b {
				total++
			}
		}
	}

	return total
}

func Part2(trees Grid) int {
	bestscore := 0
	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees[y])-1; x++ {
			start := trees[y][x]
			var n, i int
			for n, i = 0, x+1; i < len(trees[y]); i++ {
				if trees[y][i] < start {
					n++
				} else {
					n++
					break
				}
			}
			score := n
			for n, i = 0, x-1; i >= 0; i-- {
				if trees[y][i] < start {
					n++
				} else {
					n++
					break
				}
			}
			score *= n
			for n, i = 0, y+1; i < len(trees); i++ {
				if trees[i][x] < start {
					n++
				} else {
					n++
					break
				}
			}
			score *= n
			for n, i = 0, y-1; i >= 0; i-- {
				if trees[i][x] < start {
					n++
				} else {
					n++
					break
				}
			}
			score *= n
			if score > bestscore {
				bestscore = score
			}
		}
	}

	return bestscore
}

func Part2Slice(trees GridSlice) int {
	bestscore := 0
	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees[y])-1; x++ {
			start := trees[y][x]
			var n, i int
			for n, i = 0, x+1; i < len(trees[y]); i++ {
				if trees[y][i] < start {
					n++
				} else {
					n++
					break
				}
			}
			score := n
			for n, i = 0, x-1; i >= 0; i-- {
				if trees[y][i] < start {
					n++
				} else {
					n++
					break
				}
			}
			score *= n
			for n, i = 0, y+1; i < len(trees); i++ {
				if trees[i][x] < start {
					n++
				} else {
					n++
					break
				}
			}
			score *= n
			for n, i = 0, y-1; i >= 0; i-- {
				if trees[i][x] < start {
					n++
				} else {
					n++
					break
				}
			}
			score *= n
			if score > bestscore {
				bestscore = score
			}
		}
	}

	return bestscore
}

func main() {
	test := ParseGrid(strings.NewReader(`30373
25512
65332
33549
35390`))

	input := ParseGrid(PuzzleInput())
	fmt.Println("Part1", Part1(test), Part1(input))
	fmt.Println("Part2", Part2(test), Part2(input), Part2Slice(ParseGridSlice(PuzzleInput())))
}
