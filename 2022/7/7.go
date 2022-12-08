package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func dirSize(scan *bufio.Scanner, msgSize int, total *int) int {
	size := 0
	for scan.Scan() {
		line := scan.Text()
		if line == "$ cd .." {
			return size
		} else if line[0:4] == "$ cd" {
			d := dirSize(scan, msgSize, total)
			size += d
			if d <= msgSize {
				*total += d
			}
		}
		if line[0:4] == "$ ls" {
			continue
		}
		if line[0:4] != "dir " {
			var item int
			fmt.Sscanf(line, "%d ", &item)
			size += item
		}
	}
	return size
}

func minDir(scan *bufio.Scanner, requiredMin int, min *int) int {
	size := 0
	for scan.Scan() {
		line := scan.Text()
		if line == "$ cd .." {
			return size
		} else if line[0:4] == "$ cd" {
			d := minDir(scan, requiredMin, min)
			size += d
			if d >= requiredMin && d < *min {
				*min = d
			}
		}
		if line[0:4] == "$ ls" {
			continue
		}
		if line[0:4] != "dir " {
			var item int
			fmt.Sscanf(line, "%d ", &item)
			size += item
		}
	}
	return size
}

func sumSize(scan *bufio.Scanner, total *int) {
	for scan.Scan() {
		line := scan.Text()
		if line[0:1] == "$" {
			continue
		}
		if line[0:4] != "dir " {
			var item int
			fmt.Sscanf(line, "%d ", &item)
			*total += item
		}
	}
}

func Part1(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	total := 0
	top := dirSize(scan, 100000, &total)
	if top <= 100000 {
		total += top
	}
	return total
}

func Part2(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	total := 0
	sumSize(scan, &total)
	limit := 30000000 - (70000000 - total)
	_, _ = r.Seek(0, io.SeekStart)
	scan = bufio.NewScanner(r)
	min := math.MaxInt
	minDir(scan, limit, &min)
	return min
}

func main() {
	test := strings.NewReader(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`)

	input := PuzzleInput()
	fmt.Println("Part1", Part1(test), Part1(input))
	fmt.Println("Part2", Part2(test), Part2(input))
}
