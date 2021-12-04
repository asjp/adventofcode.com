package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("9/input.txt")
	scanner := bufio.NewScanner(f)
	buf := make([]int, 25)
	n := 1
	for scanner.Scan() {
		line := scanner.Text()
		a, _ := strconv.Atoi(line)

		fmt.Println(n, a)
		if n > 25 {
			// is a the sum of two numbers in buf?
			found := false
			for i, b := range buf {
				for j, c := range buf {
					if i == j {
						continue
					}
					if b + c == a {
						found = true
						break
					}
				}
			}
			if !found {
				fmt.Println(a)
				break
			}
			buf = buf[1:]
		}
		buf = append(buf, a)
		n++
	}
}
