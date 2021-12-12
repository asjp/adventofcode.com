package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Cave struct {
	connections []string
}

func isSmall(name string) bool {
	return name[0] >= 97
}

func addLink(caves map[string]Cave, from, to string) {
	if _, ok := caves[from]; !ok {
		caves[from] = Cave{}
	}
	caves[from] = Cave{
		append(caves[from].connections, to),
	}
}

func copyMap(from map[string]struct{}) map[string]struct{} {
	c := make(map[string]struct{})
	for a, b := range from {
		c[a] = b
	}
	return c
}

func CountPaths(part int, caves map[string]Cave, smallVisited map[string]struct{}, from string) int {
	if from == "end" {
		return 1
	}
	if isSmall(from) {
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
		paths += CountPaths(part, caves, sv, l)
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
	defer timeTrack(time.Now())
	return CountPaths(1, ReadCaves(r), map[string]struct{}{}, "start")
}

func Part2(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	return CountPaths(2, ReadCaves(r), map[string]struct{}{}, "start")
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

	expect(10, Part1(test1), "Part1 - test1")
	expect(19, Part1(test2), "Part1 - test2")
	expect(226, Part1(test3), "Part1 - test3")
	fmt.Println("Part1 - puzzle", Part1(input))

	expect(36, Part2(test1), "Part2 - test1")
	expect(103, Part2(test2), "Part2 - test2")
	expect(3509, Part2(test3), "Part2 - test3")
	fmt.Println("Part2 - puzzle", Part2(input))
}
