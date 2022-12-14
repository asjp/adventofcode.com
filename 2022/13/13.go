package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type List []interface{}

type Pair struct {
	left, right List
}

func Parse(s string) []Pair {
	scan := bufio.NewScanner(strings.NewReader(s))
	pairs := []Pair{}
	for scan.Scan() {
		p := Pair{}
		b := scan.Bytes()
		p.left = ParseList(bytes.NewReader(b[1:]))
		scan.Scan()
		b = scan.Bytes()
		p.right = ParseList(bytes.NewReader(b[1:]))
		pairs = append(pairs, p)
		scan.Scan()
	}
	return pairs
}

func ParseList(r io.RuneReader) List {
	var (
		result List
		last   string
	)
	for {
		rune, _, err := r.ReadRune()
		if err != nil {
			break
		}
		switch rune {
		case '[':
			result = append(result, ParseList(r))
			break
		case ',':
			if len(last) > 0 {
				result = append(result, Int(last))
				last = ""
			}
			break
		case ']':
			if len(last) > 0 {
				result = append(result, Int(last))
			}
			return result
		case ' ':
			break
		default:
			last += string(rune)
		}
	}
	return result
}

func Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Part1(pairs []Pair) int {
	n := []int{}
	for i, p := range pairs {
		if Compare(p.left, p.right) > 0 {
			n = append(n, i+1)
		}
	}
	return SumSlice(n)
}

func IsList(a interface{}) bool {
	_, ok := a.(List)
	return ok
}

func ToList(a interface{}) List {
	return List{a}
}

func Compare(left, right interface{}) int {
	if IsList(left) {
		if IsList(right) {
			for i, v := range left.(List) {
				if i >= len(right.(List)) {
					return -1
				}
				c := Compare(v, right.(List)[i])
				if c != 0 {
					return c
				}
			}
			if len(left.(List)) < len(right.(List)) {
				return 1
			}
		} else {
			return Compare(left, ToList(right))
		}
	} else {
		if IsList(right) {
			return Compare(ToList(left), right)
		} else {
			a := left.(int)
			b := right.(int)
			if a == b {
				return 0
			} else if a > b {
				return -1
			}
			return 1
		}
	}
	return 0
}

func SumSlice(s []int) int {
	n := 0
	for _, a := range s {
		n += a
	}
	return n
}

func Equal(a, b List) bool {
	if IsList(a) {
		if IsList(b) {
			return Compare(a, b) == 0
		} else {
			return false
		}
	} else {
		if IsList(b) {
			return false
		}
	}
	return true
}

type Lists []List

func (l Lists) Len() int {
	return len(l)
}

func (l Lists) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l Lists) Less(i, j int) bool {
	return Compare(l[i], l[j]) >= 0
}

func Part2(pairs []Pair) int {
	lists := Lists{{List{2}}, {List{6}}}
	for _, p := range pairs {
		lists = append(lists, p.left, p.right)
	}
	sort.Sort(lists)
	n := 1
	for i, l := range lists {
		if Equal(l, List{List{2}}) || Equal(l, List{List{6}}) {
			n *= i + 1
		}
	}
	return n
}

//go:embed input.txt
var puzzleinput string

//go:embed example.txt
var examplestr string

func main() {
	example := Parse(examplestr)
	input := Parse(puzzleinput)

	fmt.Println("Part1", Part1(example), Part1(input))
	fmt.Println("Part2", Part2(example), Part2(input))
}
