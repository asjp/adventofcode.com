package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var countcontain map[string][]string

func main() {
	f, _ := os.Open("7/input.txt")
	scanner := bufio.NewScanner(f)
	countcontain = make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Split(line, " bags contain ")
		if a[1] == "no other bags." {
			continue
		}
		b := strings.Split(a[1], ", ")
		for _, c := range b {
			d := strings.Split(c, " ")
			bag := fmt.Sprintf("%s %s %s", d[0], d[1], d[2])
			countcontain[a[0]] = append(countcontain[a[0]], bag)
		}
	}
	n := count("shiny gold", 0) - 1 // minus the shiny gold bag itself (doh!)
	fmt.Println(n)
}

func count(target string, depth int) int {
	n := 1
	if a, ok := countcontain[target]; ok {
		for _, b := range a {
			c := strings.SplitN(b, " ", 2)
			d, _ := strconv.Atoi(c[0])
			e := count(c[1], depth+1)
			fmt.Println("+", d, "*", e)
			n += d * e
		}
	}
	return n
}