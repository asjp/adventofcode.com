package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Slope struct {
	x, trees, right, down int
}

func main() {
	f, _ := os.Open("4/input.txt")
	scanner := bufio.NewScanner(f)

	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := 0

	var record string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			found := 0
			for _, k := range keys {
				if strings.Contains(record, k+":") {
					found++
				}
			}
			if found == len(keys) {
				valid++
			}
			record = ""
		} else {
			record += line
		}
	}
	fmt.Println(valid)
}
