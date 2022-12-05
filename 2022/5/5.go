package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
)

func PuzzleInput() io.ReadSeeker {
	_, filename, _, _ := runtime.Caller(0)
	today := path.Dir(filename)
	f, err := os.Open(path.Join(today, "input.txt"))
	if err != nil {
		fmt.Println(err)
	}
	return f
}

func Part1(r io.ReadSeeker) string {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	var stacks [9]string
	for scan.Scan() {
		line := scan.Text()
		if !strings.Contains(line, "[") {
			break
		}
		n := 0
		for i := 1; i < len(line); i = i + 4 {
			if line[i] != ' ' {
				stacks[n] += string(line[i])
			}
			n++
		}
	}
	scan.Scan()
	for scan.Scan() {
		var n, a, b int
		fmt.Sscanf(scan.Text(), "move %d from %d to %d", &n, &a, &b)
		a--
		b--
		if n > len(stacks[a]) {
			n = len(stacks[a])
		}
		for i := 0; i < n; i++ {
			stacks[b] = stacks[a][:1] + stacks[b]
			stacks[a] = stacks[a][1:]
		}
	}
	msg := ""
	for _, s := range stacks {
		if len(s) > 0 {
			msg += string(s[0:1])
		}
	}
	return msg
}

func Part2(r io.ReadSeeker) string {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	var stacks [9]string
	for scan.Scan() {
		line := scan.Text()
		if !strings.Contains(line, "[") {
			break
		}
		n := 0
		for i := 1; i < len(line); i = i + 4 {
			if line[i] != ' ' {
				stacks[n] += string(line[i])
			}
			n++
		}
	}
	scan.Scan()
	for scan.Scan() {
		var n, a, b int
		fmt.Sscanf(scan.Text(), "move %d from %d to %d", &n, &a, &b)
		a--
		b--
		if n > len(stacks[a]) {
			n = len(stacks[a])
		}
		stacks[b] = stacks[a][:n] + stacks[b]
		stacks[a] = stacks[a][n:]
	}
	msg := ""
	for _, s := range stacks {
		if len(s) > 0 {
			msg += string(s[0:1])
		}
	}
	return msg
}

func main() {
	test := strings.NewReader(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

	input := PuzzleInput()
	fmt.Println("Part1", Part1(test), Part1(input))
	fmt.Println("Part2", Part2(test), Part2(input))
}
