package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func P1(r io.Reader) int {
	i := 0
	s := bufio.NewScanner(r)
	s.Scan()
	l, _ := strconv.Atoi(s.Text())
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		if n > l {
			i++
		}
		l = n
	}
	return i
}

func I(w [4]int) int {
	if w[3]+w[1]+w[2] > w[0]+w[1]+w[2] {
		return 1
	} else {
		return 0
	}
}

func P2(r io.Reader) int {
	i := 0
	s := bufio.NewScanner(r)
	w := [4]int{}
	c := 0
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		if c < 4 {
			w[c] = n
		} else {
			i += I(w)
			w[0], w[1], w[2], w[3] = w[1], w[2], w[3], n
		}
		c++
	}
	i += I(w)
	return i
}

func main() {
	f, _ := os.Open("2021/1/input.txt")
	//fmt.Println(P1(f))
	fmt.Println(P2(f))
}
