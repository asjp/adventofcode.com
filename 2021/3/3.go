package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Part1(r io.Reader) int64 {
	var b []int
	n := 0
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		s := scan.Text()
		if b == nil {
			b = make([]int, len(s), len(s))
		}
		for i, a := range s {
			if a == '1' {
				b[i] += 1
			}
		}
		n++
	}
	gstr, estr := "", ""
	for _, bb := range b {
		if bb >= n/2 {
			gstr = gstr + "1"
			estr = estr + "0"
		} else {
			gstr = gstr + "0"
			estr = estr + "1"
		}
	}
	gamma, _ := strconv.ParseInt(gstr, 2, 64)
	epsilon, _ := strconv.ParseInt(estr, 2, 64)
	return gamma * epsilon
}

func mostFreqBit(lines []string, bitNum int) uint8 {
	n := 0
	for _, a := range lines {
		if a[bitNum] == '1' {
			n++
		}
	}
	if float64(n) >= float64(len(lines))/2. {
		return '1'
	} else {
		return '0'
	}
}

func leastFreqBit(lines []string, bitNum int) uint8 {
	n := 0
	for _, a := range lines {
		if a[bitNum] == '1' {
			n++
		}
	}
	if float64(n) >= float64(len(lines))/2. {
		return '0'
	} else {
		return '1'
	}
}

func Part2(r io.Reader) int64 {
	var lines []string
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		s := scan.Text()
		lines = append(lines, s)
	}

	oxygen := make([]string, len(lines))
	copy(oxygen, lines)
	bitNum := 0
	for {
		b := mostFreqBit(oxygen, bitNum)

		var keep []string
		for _, line := range oxygen {
			if line[bitNum] == b {
				keep = append(keep, line)
			}
		}

		oxygen = make([]string, len(keep))
		copy(oxygen, keep)
		if len(oxygen) < 2 {
			break
		}

		bitNum++
		if bitNum >= len(oxygen[0]) {
			break
		}
	}

	co2 := make([]string, len(lines))
	copy(co2, lines)
	bitNum = 0
	for {
		b := leastFreqBit(co2, bitNum)
		var keep []string
		for _, line := range co2 {
			if line[bitNum] == b {
				keep = append(keep, line)
			}
		}

		co2 = make([]string, len(keep))
		copy(co2, keep)
		if len(co2) < 2 {
			break
		}

		bitNum++
		if bitNum >= len(co2[0]) {
			break
		}
	}

	orating, _ := strconv.ParseInt(oxygen[0], 2, 64)
	crating, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Println(orating, crating)
	return orating * crating
}

func main() {
	test := strings.NewReader(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`)
	fmt.Println("Part2", Part2(test))
	input, err := os.Open("2021/3/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part2", Part2(input))

}
