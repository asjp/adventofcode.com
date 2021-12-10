package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

/*
func runeStr(r []rune) string {
	s := ""
	for _, i := range r {
		s += string(i)
	}
	return s
}
*/

func opposite(r rune) rune {
	switch r {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	}
	return ' '
}

func Calc(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	score := 0
	for s.Scan() {
		line := s.Text()
		m := make([]rune, 0)
		for _, c := range line {
			linecorrupt := false
			switch c {
			case '{', '[', '(', '<':
				m = append(m, c)
			case '}', ']', ')', '>':
				if m[len(m)-1] == opposite(c) {
					m = append([]rune(nil), m[:len(m)-1]...)
				} else {
					// corrupt
					linecorrupt = true
					switch c {
					case ')':
						score += 3
					case ']':
						score += 57
					case '}':
						score += 1197
					case '>':
						score += 25137
					}
				}
			}
			if linecorrupt {
				break
			}
		}
	}

	return score
}

func Calc2(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	scores := make([]int, 0)
	for s.Scan() {
		line := s.Text()
		m := make([]rune, 0)
		linecorrupt := false
		for _, c := range line {
			linecorrupt = false
			switch c {
			case '{', '[', '(', '<':
				m = append(m, c)
			case '}', ']', ')', '>':
				if m[len(m)-1] == opposite(c) {
					m = append([]rune(nil), m[:len(m)-1]...)
				} else {
					// corrupt
					linecorrupt = true
				}
			}
			if linecorrupt {
				break
			}
		}
		if !linecorrupt {
			// complete the line
			score := 0
			//fmt.Println(runeStr(m))
			for i := len(m) - 1; i >= 0; i-- {
				switch opposite(m[i]) {
				case ')':
					score = score*5 + 1
				case ']':
					score = score*5 + 2
				case '}':
					score = score*5 + 3
				case '>':
					score = score*5 + 4
				}
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func Part1(r io.ReadSeeker) int {
	return Calc(r)
}

func Part2(r io.ReadSeeker) int {
	return Calc2(r)
}

func main() {
	test1 := strings.NewReader(`[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`)
	fmt.Println("Part1 - test1", Part1(test1))
	fmt.Println("Part2 - test1", Part2(test1))

	input, err := os.Open("2021/10/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))
	fmt.Println("Part2 - puzzle", Part2(input))
}
