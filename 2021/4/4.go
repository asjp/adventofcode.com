package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	v      int
	marked bool
}

func sumUnmarked(board [][]Cell) int {
	sum := 0
	for _, ii := range board {
		for _, jj := range ii {
			if !jj.marked {
				sum += jj.v
			}
		}
	}
	return sum
}

func checkWin(board [][]Cell) int {
	//check rows
	for _, ii := range board {
		for j, jj := range ii {
			if !jj.marked {
				break
			}
			if j == len(ii)-1 {
				return sumUnmarked(board)
			}
		}
	}
	// check cols
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			c := board[j][i]
			if !c.marked {
				break
			}
			if j == 4 {
				return sumUnmarked(board)
			}
		}
	}
	return 0
}

func Calc(r io.ReadSeeker, whichBoard int) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	s.Scan()
	drawStrs := strings.Split(s.Text(), ",")
	draw := make([]int, len(drawStrs))
	for i, d := range drawStrs {
		draw[i], _ = strconv.Atoi(d)
	}
	s.Scan()
	var boards [][][]Cell
	var b [][]Cell
	for s.Scan() {
		if s.Text() == "" {
			boards = append(boards, b)
			b = make([][]Cell, 0)
			continue
		}
		var l []Cell
		x := strings.Split(s.Text(), " ")
		for _, y := range x {
			if y == "" {
				continue
			}
			z, _ := strconv.Atoi(y)
			l = append(l, Cell{z, false})
		}
		b = append(b, l)
	}
	winningBoards := make(map[int]struct{}, 0)
	if whichBoard == -1 {
		whichBoard = len(boards)
	}
	for _, d := range draw {
		for bNum, b := range boards {
			for _, xx := range b {
				for y, yy := range xx {
					if yy.v == d {
						xx[y].marked = true
						//fmt.Println(d)
						//fmt.Println(b)
						sum := checkWin(b)
						if sum > 0 {
							//fmt.Println(bNum, winningBoards)
							winningBoards[bNum] = struct{}{}
							if len(winningBoards) == whichBoard {
								return sum * d
							}
						}
					}
				}
			}
		}
	}
	return 0
}

func Part1(r io.ReadSeeker) int {
	return Calc(r, 1)
}

func Part2(r io.ReadSeeker) int {
	return Calc(r, -1)
}

func main() {
	test := strings.NewReader(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7

`)
	fmt.Println("Part1 - test", Part1(test))
	fmt.Println("Part2 - test", Part2(test))

	input, err := os.Open("2021/4/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))
	fmt.Println("Part2 - puzzle", Part2(input))
}
