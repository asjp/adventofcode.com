package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Part1(r io.Reader) int {
	incr := 0
	scan := bufio.NewScanner(r)
	scan.Scan()
	last, _ := strconv.Atoi(scan.Text())
	for scan.Scan() {
		n, _ := strconv.Atoi(scan.Text())
		if n > last {
			incr++
		}
		last = n
	}
	return incr
}

func Part2(r io.Reader) int {
	incr := 0
	scan := bufio.NewScanner(r)
	w := [4]int{}
	count := 0
	for scan.Scan() {
		n, _ := strconv.Atoi(scan.Text())
		if count < 4 {
			w[count] = n
		} else {
			win1 := w[0] + w[1] + w[2]
			win2 := w[3] + w[1] + w[2]
			if win2 > win1 {
				incr++
			}
			w[0], w[1], w[2], w[3] = w[1], w[2], w[3], n
		}
		count++
	}
	win1 := w[0] + w[1] + w[2]
	win2 := w[3] + w[1] + w[2]
	if win2 > win1 {
		incr++
	}
	return incr
}

func main() {
	test := strings.NewReader(`199
200
208
210
200
207
240
269
260
263
`)

	input, err := os.Open("2021/1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Part1", Part1(test), Part1(input))
	fmt.Println("Part2", Part2(test), Part2(input))
}
