package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

type Input struct {
	start []byte
	rules map[string]byte
}

func ReadInput(r io.ReadSeeker) Input {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	s.Scan()
	var input Input
	input.start = []byte(s.Text())
	input.rules = make(map[string]byte)
	s.Scan()
	for s.Scan() {
		s := s.Text()
		p := strings.Split(s, " -> ")
		input.rules[p[0]] = p[1][0]
	}
	return input
}

func Iterate(steps int, input Input) map[string]int {
	counts := make(map[string]int)
	singles := make(map[string]int)

	for j := 0; j < len(input.start)-1; j++ {
		key := string(input.start[j : j+2])
		if v, ok := counts[key]; ok {
			counts[key] = v + 1
		} else {
			counts[key] = 1
		}
		key = string(input.start[j])
		if v, ok := singles[key]; ok {
			singles[key] = v + 1
		} else {
			singles[key] = 1
		}
	}
	k := string(input.start[len(input.start)-1])
	if v, ok := singles[k]; ok {
		singles[k] = v + 1
	} else {
		singles[k] = 1
	}

	//fmt.Println(counts)
	//fmt.Println(singles)
	for i := 0; i < steps; i++ {
		deltas := make(map[string]int)
		for c, v := range counts {
			// c=AB
			if r, ok := input.rules[c]; ok {
				// r=C
				addDelta(deltas, string(r), v)
				addDelta(deltas, string(c[0])+string(r), v)
				addDelta(deltas, string(r)+string(c[1]), v)
				addDelta(deltas, c, -v)
			}
		}
		for k, d := range deltas {
			if len(k) == 2 {
				if c, ok := counts[k]; ok {
					counts[k] = c + d
				} else {
					counts[k] = d
				}
			} else {
				if c, ok := singles[k]; ok {
					singles[k] = c + d
				} else {
					singles[k] = d
				}
			}
		}
		//fmt.Println("counts", counts)
		//fmt.Println("singles", singles, Sum(singles))
	}
	return singles
}

func addDelta(deltas map[string]int, key string, d int) {
	if v, ok := deltas[key]; ok {
		deltas[key] = v + d
	} else {
		deltas[key] = d
	}
}

func MakeArray(m map[string]int) (r []int) {
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func Part1and2(iterations int, r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	result := Iterate(iterations, ReadInput(r))
	freq := MakeArray(result)
	sort.Ints(freq)
	return freq[len(freq)-1] - freq[0]
}

func timeTrack(start time.Time) {
	fmt.Printf("(%10s) ", time.Since(start))
}

func expect(expected, actual int, msg string) {
	if expected == actual {
		fmt.Printf("\033[1;32mOK\033[0m")
	} else {
		fmt.Printf("\033[1;31mFAIL\033[0m (expected %d, actual %d)", expected, actual)
	}
	fmt.Println(" ", msg)
}

func main() {
	test1 := strings.NewReader(`NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`)

	input, err := os.Open("2021/14/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	expect(1588, Part1and2(10, test1), "Part1 - test1")
	fmt.Println("Part1 - puzzle", Part1and2(10, input))

	expect(2188189693529, Part1and2(40, test1), "Part2 - test1")
	fmt.Println("Part2 - puzzle", Part1and2(40, input))
}
