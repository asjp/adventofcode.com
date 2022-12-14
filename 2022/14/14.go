package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

type Cave struct {
	space          [][]bool
	sandx, bottomy int
	sand           Coord
	lasty          int
}

type Coord struct {
	x, y int
}

func Parse(s string, space int) Cave {
	scan := bufio.NewScanner(strings.NewReader(s))

	lines := [][]Coord{}
	minx := math.MaxInt
	miny := math.MaxInt
	maxx := 0
	maxy := 0

	for scan.Scan() {
		coords := strings.Split(scan.Text(), " -> ")
		line := []Coord{}
		for _, ctext := range coords {
			next := Coord{}
			fmt.Sscanf(ctext, "%d,%d", &next.x, &next.y)
			line = append(line, next)
			if next.x < minx {
				minx = next.x
			}
			if next.x > maxx {
				maxx = next.x
			}
			if next.y < miny {
				miny = next.y
			}
			if next.y > maxy {
				maxy = next.y
			}
		}
		lines = append(lines, line)
	}
	miny = 0
	minx -= space
	maxx += 2 * space
	maxy++
	cave := Cave{}
	cave.space = make([][]bool, (maxx-minx)+1)
	cave.sandx = 500 - minx
	cave.bottomy = maxy

	for i := 0; i < len(cave.space); i++ {
		cave.space[i] = make([]bool, (maxy-miny)+1)
	}

	for _, line := range lines {
		for i := 1; i < len(line); i++ {
			from := Coord{line[i-1].x - minx, line[i-1].y - miny}
			to := Coord{line[i].x - minx, line[i].y - miny}
			Fill(cave, from, to)
		}
	}

	return cave
}

func Fill(cave Cave, from, to Coord) {
	if from.x == to.x {
		if from.y < to.y {
			for y := from.y; y <= to.y; y++ {
				cave.space[from.x][y] = true
			}
		} else {
			for y := to.y; y <= from.y; y++ {
				cave.space[from.x][y] = true
			}
		}
	} else {
		if from.x < to.x {
			for x := from.x; x <= to.x; x++ {
				cave.space[x][from.y] = true
			}
		} else {
			for x := to.x; x <= from.x; x++ {
				cave.space[x][from.y] = true
			}
		}
	}
}

func (c Cave) String() string {
	// move cursor to 0,0 first
	all := "\033[0;0H"
	for x := 0; x < c.sandx; x++ {
		all += " "
	}
	all += "+\n"
	maxlines := len(c.space[0])
	for y := c.lasty - animateLines; y < maxlines && y < c.lasty+2; y++ {
		if y < 0 {
			y = 0
		}
		line := ""
		for x := 0; x < len(c.space); x++ {
			if c.sand.x == x && c.sand.y == y {
				line += "o"
			} else if c.space[x][y] {
				line += "#"
			} else {
				line += "."
			}
		}
		all += line + "\n"
	}
	return all
}

func AddSand(cave *Cave) {
	cave.sand = Coord{cave.sandx, 0}
}

func MoveSand(cave *Cave, fixedBottom bool) bool {
	// falling through?
	if cave.sand.y == cave.bottomy {
		if !fixedBottom {
			cave.sand.y++
		}
		return false
	}
	// move down?
	if !cave.space[cave.sand.x][cave.sand.y+1] {
		cave.sand.y++
		return true
	}
	// move down-left?
	if !cave.space[cave.sand.x-1][cave.sand.y+1] {
		cave.sand.x--
		cave.sand.y++
		return true
	}
	// move down-r0ght?
	if !cave.space[cave.sand.x+1][cave.sand.y+1] {
		cave.sand.x++
		cave.sand.y++
		return true
	}
	// stuck
	return false
}

var (
	animate = true
)

const (
	animateLines = 45
	animateEvery = 1
)

func Part(cave Cave, fixedBottom bool) int {
	countsand := 0
	if animate {
		fmt.Print("\033[H\033[2J")
	}
	for {
		AddSand(&cave)
		if animate && countsand%animateEvery == 0 {
			fmt.Println(cave)
			time.Sleep(20 * time.Millisecond)
		}
		for MoveSand(&cave, fixedBottom) {
		}
		if cave.sand.y > cave.lasty {
			cave.lasty = cave.sand.y
		}
		if cave.sand.y > cave.bottomy {
			// lost to the abyss (part 1)
			break
		}
		// sand becomes part of the cave
		cave.space[cave.sand.x][cave.sand.y] = true
		countsand++
		if cave.sand.y == 0 {
			// full to the top (part 2)
			break
		}
	}
	if animate {
		fmt.Println(cave)
	}
	return countsand
}

func Part1(input string) int {
	return Part(Parse(input, 1), false)
}

func Part2(input string) int {
	return Part(Parse(input, 100), true)
}

//go:embed input.txt
var puzzleinput string

//go:embed example.txt
var examplestr string

func main() {
	p1example := Part1(examplestr)
	time.Sleep(2 * time.Second)
	p1puzzle := Part1(puzzleinput)
	fmt.Println("Part1", p1example, p1puzzle)
	animate = false
	fmt.Println("Part2", Part2(examplestr), Part2(puzzleinput))
}
