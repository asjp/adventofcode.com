package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strings"
)

type Coord struct {
	x, y int
}

func Parse(s string) ([][]byte, Coord, Coord) {
	scan := bufio.NewScanner(strings.NewReader(s))
	mymap := [][]byte{}
	y := 0
	var mystart, myend Coord
	for scan.Scan() {
		line := scan.Text()
		if startIndex := strings.Index(line, "S"); startIndex >= 0 {
			mystart = Coord{startIndex, y}
			line = strings.Replace(line, "S", "a", 1)
		}
		if endIndex := strings.Index(line, "E"); endIndex >= 0 {
			myend = Coord{endIndex, y}
			line = strings.Replace(line, "E", "z", 1)
		}
		mymap = append(mymap, []byte(line))
		y++
	}
	return mymap, mystart, myend
}

func Part1(mymap [][]byte, mystart, myend Coord) int {
	distance := map[Coord]int{myend: 0}
	ShortestPath(distance, mymap, myend)
	return distance[mystart]
}

func Part2(mymap [][]byte, mystart, myend Coord) int {
	distance := map[Coord]int{myend: 0}
	ShortestPath(distance, mymap, myend)
	min := math.MaxInt
	for c, d := range distance {
		if mymap[c.y][c.x] == 'a' && d < min {
			min = d
		}
	}
	return min
}

func ShortestPath(distance map[Coord]int, mymap [][]byte, curr Coord) {
	next := []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, n := range next {
		c := Coord{curr.x + n.x, curr.y + n.y}
		if InBounds(mymap, c) && CanWalk(mymap, c, curr) {
			if d, ok := distance[c]; !ok || d > distance[curr]+1 {
				distance[c] = distance[curr] + 1
				ShortestPath(distance, mymap, c)
			}
		}
	}
}

func InBounds(mymap [][]byte, c Coord) bool {
	return c.x >= 0 && c.y >= 0 && c.y < len(mymap) && c.x < len(mymap[0])
}

func Adjacent(a, b Coord) bool {
	dx := Abs(a.x - b.x)
	dy := Abs(a.y - b.y)
	return (dx == 1 && dy == 0) || (dx == 0 && dy == 1)
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func CanWalk(mymap [][]byte, a, b Coord) bool {
	if !Adjacent(a, b) {
		return false
	}
	return mymap[b.y][b.x] < mymap[a.y][a.x]+2
}

//go:embed input.txt
var puzzleinput string

func main() {
	example, ea, eb := Parse(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`)

	input, ia, ib := Parse(puzzleinput)
	fmt.Println("Part1", Part1(example, ea, eb), Part1(input, ia, ib))
	fmt.Println("Part2", Part2(example, ea, eb), Part2(input, ia, ib))
}
