package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
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

func Part1(r io.ReadSeeker) int {
	scan := bufio.NewScanner(r)
	max := 0
	cur := 0
	for scan.Scan() {
		n, _ := strconv.Atoi(scan.Text())
		if n > 0 {
			cur = cur + n
		} else {
			if cur > max {
				max = cur
			}
			cur = 0
		}
	}
	return max
}

func Part2(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	max1, max2, max3 := 0, 0, 0
	cur := 0
	for scan.Scan() {
		n, _ := strconv.Atoi(scan.Text())
		if n > 0 {
			cur = cur + n
		} else {
			if cur > max1 {
				max3, max2, max1 = max2, max1, cur
			} else if cur > max2 {
				max3, max2 = max2, cur
			} else if cur > max3 {
				max3 = cur
			}
			cur = 0
		}
	}
	return max1 + max2 + max3
}

func main() {
	test := strings.NewReader(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

`)

	input := PuzzleInput()
	fmt.Println("Part1", Part1(test), Part1(input))
	fmt.Println("Part2", Part2(test), Part2(input))
}
