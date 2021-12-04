package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("13/test.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	scanner.Scan()
	fmt.Println(part2(scanner.Text()))
}

func part2(input string) int {
	busdata := strings.Split(input, ",")
	var buses []int
	for _, b := range busdata {
		n, _ := strconv.Atoi(b)
		buses = append(buses, n)
	}
	fmt.Println(buses)

	mods := make(map[int]int)
	for i, b := range buses {
		if b > 0 {
			mods[b] = i % b
		}
	}

	earliest := 0
	i := 1
	for b, mod := range mods {
		fmt.Println(earliest, i, b, mod)
		for ; earliest % b != mod; earliest += i {
		}
		i *= b
	}

	return earliest
}