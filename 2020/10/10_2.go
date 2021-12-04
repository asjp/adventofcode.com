package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var buf []int
var num map[int]int

func main() {
	f, _ := os.Open("10/input.txt")
	scanner := bufio.NewScanner(f)
	buf = make([]int, 1, 1)
	buf[0] = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			a, _ := strconv.Atoi(line)
			buf = append(buf, a)
		}
	}
	sort.Ints(buf)
	buf = append(buf, buf[len(buf)-1] + 3)
	fmt.Println(buf)

	num = make(map[int]int)
	w := ways(0, "")
	fmt.Println(w)
}

func ways(a int, acc string) int {
	if a == len(buf)-1 {
		return 1
	}
	w := 0
	j := buf[a]
	//fmt.Println(a, j)
	for i := a + 1; i < len(buf) && buf[i] < j + 4; i++ {
		if n, ok := num[i]; ok {
			w += n
		} else {
			c := ways(i, fmt.Sprintf("%s %d", acc, buf[i]))
			num[i] = c
			w += c
		}
	}
	fmt.Println(w, ":", acc)
	return w
}