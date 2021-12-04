package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const invalid = 18272118

func main() {
	f, _ := os.Open("9/input.txt")
	scanner := bufio.NewScanner(f)
	buf := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		a, _ := strconv.Atoi(line)

		buf = append(buf, a)

		sum := 0
		for _, b := range buf {
			sum += b
		}

		for sum > invalid {
			sum -= buf[0]
			buf = buf[1:]
		}

		if len(buf) > 1 && sum == invalid {
			// found it!
			// add together smallest and largest in buf
			max := 0
			min := math.MaxInt64
			for _, b := range buf {
				if b < min {
					min = b
				}
				if b > max {
					max = b
				}
			}
			fmt.Println(min+max)
			break
		}
	}
}
