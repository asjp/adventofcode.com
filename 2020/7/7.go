package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var contain map[string][]string

func main() {
	f, _ := os.Open("7/input.txt")
	scanner := bufio.NewScanner(f)
	contain = make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Split(line, " bags contain ")
		if a[1] == "no other bags." {
			continue
		}
		b := strings.Split(a[1], ", ")
		for _, c := range b {
			d := strings.Split(c, " ")
			bag := fmt.Sprintf("%s %s", d[1], d[2])
			contain[a[0]] = append(contain[a[0]], bag)
		}
	}
	n := 0
	for a, b := range contain {
		//fmt.Println(a, len(b), b)
		if contains("shiny gold", a, b) {
			n++
		}
	}
	fmt.Println(n)
}

func contains(target, a string, b []string) bool {
	for _, c := range b {
		if c == target {
			fmt.Println(a, "contains", target)
			return true
		} else {
			if d, ok := contain[c]; ok {
				//fmt.Println(a, "->", c)
				found := contains(target, fmt.Sprintf("%s -> %s", a, c), d)
				if found {
					return true
				}
			}
		}
	}
	return false
}