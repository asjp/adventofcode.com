package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("13/input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	start, _ := strconv.Atoi(scanner.Text())
	fmt.Println(start)
	scanner.Scan()
	busdata := strings.Split(scanner.Text(), ",")
	fmt.Println(busdata)
	var buses []int
	for _, b := range busdata {
		if b != "x" {
			n, _ := strconv.Atoi(b)
			buses = append(buses, n)
		}
	}

	bus, earliest := 0, start
	for i := start; earliest == start; i++ {
		for _, b := range buses {
			fmt.Println(i, b, i % b)
			if i % b == 0 {
				earliest = i
				bus = b
				break
			}
		}
	}

	fmt.Println(earliest, bus)
	fmt.Println((earliest-start) * bus)
}