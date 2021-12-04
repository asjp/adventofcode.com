package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	SEAT = "L"
	OCCUPIED = "#"
	FLOOR = "."
)

var (
	area [][]string
	h, w int
)

func main() {
	f, _ := os.Open("11/input.txt")
	scanner := bufio.NewScanner(f)
	area = make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			row := make([]string, 0)
			for _, a := range line {
				row = append(row, string(a))
			}
			area = append(area, row)
		}
	}
	h = len(area)
	w = len(area[0])

	fmt.Println(area)

	part := 2

	// part 1
	if (part == 1) {
		n := 0
		for {
			changes := iterate()
			fmt.Println(changes)
			if changes == 0 {
				// count occupied
				c := 0
				for _, x := range area {
					for _, y := range x {
						if y == OCCUPIED {
							c++
						}
					}
				}
				fmt.Println("part 1", c)
				break
			}
			n++
			if n > 999 {
				break
			}
		}
	} else {
		// part 2
		n := 0
		for {
			changes := iterate2()
			fmt.Println(changes)
			if changes == 0 {
				// count occupied
				c := 0
				for _, x := range area {
					for _, y := range x {
						if y == OCCUPIED {
							c++
						}
					}
				}
				fmt.Println(c)
				break
			}
			n++
			if n > 999 {
				break
			}
		}
	}
}

type Change struct {
	x, y int
	to string
}

func iterate() int {
	changes := make([]Change, 0)
	for i, x := range area {
		for j, y := range x {
			if y == SEAT && clearAround(i, j) {
				changes = append(changes, Change{x: i, y: j, to: OCCUPIED})
			}
			if y == OCCUPIED && crowdedAround(i, j) {
				changes = append(changes, Change{x: i, y: j, to: SEAT})
			}
		}
	}
	// apply changes
	for _, c := range changes {
		area[c.x][c.y] = c.to
	}
	return len(changes)
}

func iterate2() int {
	changes := make([]Change, 0)
	for i, x := range area {
		for j, y := range x {
			if y == SEAT && clearAround2(i, j) {
				changes = append(changes, Change{x: i, y: j, to: OCCUPIED})
			}
			if y == OCCUPIED && crowdedAround2(i, j) {
				changes = append(changes, Change{x: i, y: j, to: SEAT})
			}
		}
	}
	// apply changes
	for _, c := range changes {
		area[c.x][c.y] = c.to
	}
	return len(changes)
}

func clearAround(x, y int) bool {
	for i := x-1; i<x+2; i++ {
		for j := y-1; j<y+2; j++ {
			if occupied(i, j) {
				return false
			}
		}
	}
	return true
}


func clearAround2(x, y int) bool {
	for i := -1; i<2; i++ {
		for j := -1; j<2; j++ {
			f := 1
			for isFloor(x+i*f, y+j*f) {
				f++
			}
			if occupied(x+i*f, y+j*f) {
				return false
			}
		}
	}
	return true
}

func occupied(x, y int) bool {
	if x < 0 || x >= h || y < 0 || y >= w {
		return false
	}
	return area[x][y] == OCCUPIED
}

func isFloor(x, y int) bool {
	if x < 0 || x >= h || y < 0 || y >= w {
		return false
	}
	return area[x][y] == FLOOR
}

func crowdedAround(x, y int) bool {
	c := 0
	for i := x-1; i<x+2; i++ {
		for j := y-1; j<y+2; j++ {
			if occupied(i, j) {
				c++
			}
		}
	}
	return c >= 5 // 4 plus the occupied seat
}

func crowdedAround2(x, y int) bool {
	c := 0
	for i := -1; i<2; i++ {
		for j := -1; j<2; j++ {
			f := 1
			for isFloor(x+i*f, y+j*f) {
				f++
			}
			if occupied(x+i*f, y+j*f) {
				c++
			}
		}
	}
	return c >= 6 // 5 plus the occupied seat
}

