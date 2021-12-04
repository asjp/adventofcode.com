package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op string
	d int
}

func main() {
	f, _ := os.Open("8/input.txt")
	scanner := bufio.NewScanner(f)
	lines := make([]Instruction, 0)
	exec := make(map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Split(line, " ")
		d, _ := strconv.Atoi(a[1])
		lines = append(lines, Instruction{op: a[0], d: d})
	}

	for i := 0; i < len(lines); i++ {
		old := lines[i].op
		if lines[i].op == "jmp" {
			lines[i].op = "nop"
		} else if lines[i].op == "nop" {
			lines[i].op = "jmp"
		} else {
			continue
		}

		acc := 0
		cur := 0
		for {
			// if moved off end we're done
			if cur == len(lines) {
				fmt.Println(acc)
				os.Exit(0)
			}

			// if gone too far, abandon
			if cur > len(lines) {
				break
			}

			// are we looping?
			if _, ok := exec[cur]; ok {
				break
			}

			// mark current line visited
			exec[cur] = true
			// execute current line
			do := lines[cur]
			switch (do.op) {
			case "acc":
				acc += do.d
				cur++
			case "nop":
				cur++
			case "jmp":
				cur += do.d
			}
		}

		// reset for next iteration
		lines[i].op = old
		exec = make(map[int]bool)
	}

}
