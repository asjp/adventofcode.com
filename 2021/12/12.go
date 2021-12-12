package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Cave struct {
	small       bool
	connections []string
}

func isSmall(name string) bool {
	return name[0] >= 97
}

func addLink(caves map[string]Cave, from, to string) {
	var (
		c  Cave
		ok bool
	)
	if c, ok = caves[from]; !ok {
		conn := make([]string, 1, 1)
		conn[0] = to
		caves[from] = Cave{
			isSmall(from),
			conn,
		}
	} else {
		c.connections = append(c.connections, to)
		caves[from] = Cave{
			c.small,
			c.connections,
		}
	}
}

func copyMap(from map[string]struct{}) map[string]struct{} {
	c := make(map[string]struct{})
	for a, b := range from {
		c[a] = b
	}
	return c
}

func countPaths(part int, path []string, caves map[string]Cave, smallVisited map[string]struct{}, from string) int {
	path = append(path, from)
	if from == "end" {
		//fmt.Println(path)
		return 1
	}
	c := caves[from]
	if c.small {
		if _, ok := smallVisited[from]; ok {
			if part == 1 {
				return 0
			}
			part = 1
		} else {
			smallVisited[from] = struct{}{}
		}
	}
	links := caves[from].connections
	paths := 0
	for _, l := range links {
		if l == "start" {
			continue
		}
		sv := copyMap(smallVisited)
		p := make([]string, len(path)+1)
		copy(p, path)
		paths += countPaths(part, p, caves, sv, l)
	}
	return paths
}

func ReadCaves(r io.ReadSeeker) map[string]Cave {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	caves := make(map[string]Cave)
	for s.Scan() {
		s := s.Text()
		link := strings.Split(s, "-")
		addLink(caves, link[0], link[1])
		addLink(caves, link[1], link[0])
	}
	return caves
}

func Part1(r io.ReadSeeker) int {
	caves := ReadCaves(r)
	paths := countPaths(1, []string{}, caves, map[string]struct{}{}, "start")
	return paths
}

func Part2(r io.ReadSeeker) int {
	caves := ReadCaves(r)
	paths := countPaths(2, []string{}, caves, map[string]struct{}{}, "start")
	return paths
}

func main() {
	test1 := strings.NewReader(`start-A
start-b
A-c
A-b
b-d
A-end
b-end
`)

	test2 := strings.NewReader(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`)

	test3 := strings.NewReader(`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`)

	input, err := os.Open("2021/12/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("part1 - test1", Part1(test1))
	fmt.Println("part1 - test2", Part1(test2))
	fmt.Println("part1 - test3", Part1(test3))
	fmt.Println("part1 - puzzle", Part1(input))

	fmt.Println("part2 - test1", Part2(test1))
	fmt.Println("part2 - test2", Part2(test2))
	fmt.Println("part2 - test3", Part2(test3))
	fmt.Println("part2 - puzzle", Part2(input))
}
