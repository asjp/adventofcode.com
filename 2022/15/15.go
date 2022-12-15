package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

type Area struct {
	sensor, beacon Coord
	distance       int
}

type Cave struct {
	areas    []Area
	coverage map[Coord]struct{}
	min, max Coord
}

type Coord struct {
	x, y int
}

func Parse(s string) Cave {
	scan := bufio.NewScanner(strings.NewReader(s))
	cave := Cave{}
	for scan.Scan() {
		var area Area
		fmt.Sscanf(scan.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &area.sensor.x, &area.sensor.y, &area.beacon.x, &area.beacon.y)
		area.distance = Manhattan(area.sensor, area.beacon)
		cave.areas = append(cave.areas, area)
	}
	return cave
}

func Manhattan(a, b Coord) int {
	return Abs(a.x-b.x) + Abs(a.y-b.y)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (cave *Cave) BuildCoverage() {
	cave.coverage = make(map[Coord]struct{})
	for _, a := range cave.areas {
		for dx := -a.distance; dx <= a.distance; dx++ {
			for dy := -a.distance; dy <= a.distance; dy++ {
				if Abs(dx)+Abs(dy) <= a.distance {
					cave.coverage[Coord{a.sensor.x + dx, a.sensor.y + dy}] = struct{}{}
				}
			}
		}
	}
	for c, _ := range cave.coverage {
		if c.x < cave.min.x {
			cave.min.x = c.x
		}
		if c.y < cave.min.y {
			cave.min.y = c.y
		}
		if c.x > cave.max.x {
			cave.max.x = c.x
		}
		if c.y > cave.max.y {
			cave.max.y = c.y
		}
	}
}

func (cave *Cave) BuildExtents() {
	for _, a := range cave.areas {
		if a.sensor.x-a.distance < cave.min.x {
			cave.min.x = a.sensor.x - a.distance
		}
		if a.sensor.y-a.distance < cave.min.y {
			cave.min.y = a.sensor.y - a.distance
		}
		if a.sensor.x+a.distance > cave.max.x {
			cave.max.x = a.sensor.x + a.distance
		}
		if a.sensor.y+a.distance > cave.max.y {
			cave.max.y = a.sensor.y + a.distance
		}
	}
}

func (cave Cave) CoverageY(y int) int {
	count := 0
	c := Coord{0, y}
	for c.x = cave.min.x; c.x <= cave.max.x; c.x++ {
		for _, a := range cave.areas {
			if a.beacon == c {
				continue
			}
			if a.Contains(c) {
				count++
				break
			}
		}
	}
	return count
}

func (a Area) Contains(c Coord) bool {
	return Manhattan(a.sensor, c) <= a.distance
}

func (c Coord) String() string {
	return fmt.Sprintf("[%d,%d]", c.x, c.y)
}

func (cave Cave) String() string {
	result := fmt.Sprintf("%v -> %v\n%v\n", cave.min, cave.max, cave.areas)
	for y := cave.min.y; y <= cave.max.y; y++ {
		result += fmt.Sprintf("%3d ", y)
		for x := cave.min.x; x <= cave.max.x; x++ {
			if _, ok := cave.coverage[Coord{x, y}]; ok {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}

func (c Coord) Within(limit int) bool {
	return c.x >= 0 && c.x <= limit && c.y >= 0 && c.y <= limit
}

func (a Area) Extend(size, limit int) []Coord {
	extent := a.distance + size
	points := make([]Coord, 0)
	for dx := 0; dx <= extent; dx++ {
		dy := extent - dx
		c := Coord{a.sensor.x + dx, a.sensor.y + dy}
		if c.Within(limit) {
			points = append(points, c)
		}
		if dx > 0 {
			c := Coord{a.sensor.x - dx, a.sensor.y + dy}
			if c.Within(limit) {
				points = append(points, c)
			}
		}
		if dy > 0 {
			c := Coord{a.sensor.x + dx, a.sensor.y - dy}
			if c.Within(limit) {
				points = append(points, c)
			}
		}
		if dx > 0 && dy > 0 {
			c := Coord{a.sensor.x - dx, a.sensor.y - dy}
			if c.Within(limit) {
				points = append(points, c)
			}
		}
	}
	return points
}

func (cave Cave) Search(limit int) (Coord, error) {
	// collect all the points around the sensor areas
	boundaries := map[Coord]struct{}{}
	for _, a := range cave.areas {
		coords := a.Extend(1, limit)
		for _, c := range coords {
			boundaries[c] = struct{}{}
		}
	}
	// for each boundary point, check if any sensor covers it
	for c, _ := range boundaries {
		found := false
		for _, a := range cave.areas {
			if a.Contains(c) {
				found = true
				break
			}
		}
		if !found {
			// not covered by any sensor
			return c, nil
		}
	}
	return Coord{}, fmt.Errorf("not found!")
}

func Part1(input string, y int) int {
	cave := Parse(input)
	cave.BuildExtents()
	return cave.CoverageY(y)
}

func Part2(input string, limit int) int {
	cave := Parse(input)
	coord, err := cave.Search(limit)
	if err != nil {
		panic(err)
	}
	return coord.x*4000000 + coord.y
}

//go:embed input.txt
var puzzleinput string

//go:embed example.txt
var examplestr string

func main() {
	fmt.Println("Part1", Part1(examplestr, 10), Part1(puzzleinput, 2000000))
	fmt.Println("Part2", Part2(examplestr, 20), Part2(puzzleinput, 4000000))
}
