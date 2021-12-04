package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("10/input.txt")
	scanner := bufio.NewScanner(f)
	buf := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			a, _ := strconv.Atoi(line)
			buf = append(buf, a)
		}
	}
	sort.Ints(buf)
	fmt.Println(buf)
	j := 0
	d := make(map[int]int)
	for _, v := range buf {
		d[v - j]++
		j = v
	}
	// and the built-in adapter
	j += 3
	d[3]++
	fmt.Println(j)
	fmt.Println(d)
	fmt.Println(d[1] * d[3])
}
