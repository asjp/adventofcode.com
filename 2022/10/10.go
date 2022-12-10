package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strings"
)

func Parse(s string) []int {
	scan := bufio.NewScanner(strings.NewReader(s))
	reg, x := []int{}, 1
	for scan.Scan() {
		reg = append(reg, x)
		if scan.Text() == "noop" {
			continue
		}
		var n int
		fmt.Sscanf(scan.Text(), "addx %d", &n)
		reg = append(reg, x)
		x += n
	}
	return reg
}

func Part1(reg []int) int {
	sum := 0
	for i := 20; i <= 220 && i < len(reg); i += 40 {
		sum += i * reg[i-1]
	}
	return sum
}

func Part2(reg []int) string {
	msg := make([]rune, 240)
	for i := 0; i < len(reg); i++ {
		if math.Abs(float64(reg[i]-i%40)) < 2 {
			msg[i] = '\u2588'
		} else {
			msg[i] = ' '
		}
	}
	msgstr := ""
	for i := 0; i < 6; i++ {
		msgstr += string(msg[i*40:(i+1)*40]) + "\n"
	}
	return msgstr
}

//go:embed input.txt
var puzzleinput string

//go:embed example2.txt
var example2str string

func main() {
	example := Parse(`noop
addx 3
addx -5`)
	example2 := Parse(example2str)
	input := Parse(puzzleinput)

	fmt.Println("Part1", Part1(example), Part1(example2), Part1(input))
	fmt.Println("Part2")
	fmt.Println(Part2(example2))
	fmt.Println(Part2(input))
}
