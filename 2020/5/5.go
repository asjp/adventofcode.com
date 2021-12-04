package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("5/input.txt")
	scanner := bufio.NewScanner(f)
	max := uint64(0)
	min := uint64(99999999)
	seats := make([]bool, 1000, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "B", "1")
		line = strings.ReplaceAll(line, "L", "0")
		line = strings.ReplaceAll(line, "R", "1")
		fmt.Println(line)
		i, _ := strconv.ParseUint(line, 2, 16)
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
		seats[i] = true
	}
	fmt.Println(min)
	fmt.Println(max)
	for i, b := range seats {
		if uint64(i) >= min && uint64(i) <= max {
			if !b {
				fmt.Println(i)
			}
		}
	}
}
