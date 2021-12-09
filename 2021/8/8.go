package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func Calc(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	count := 0
	for s.Scan() {
		main := strings.Split(s.Text(), "|")
		if len(main) < 2 {
			continue
		}
		output := strings.Split(main[1], " ")
		for _, o := range output {
			switch len(o) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}

func segmentContains(s, sub string) int {
	count := 0
	for _, r := range sub {
		if strings.ContainsRune(s, r) {
			count++
		}
	}
	return count
}

func Calc2(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	sum := 0
	for s.Scan() {
		main := strings.Split(s.Text(), "|")
		if len(main) < 2 {
			continue
		}
		mapping := make(map[string]int)
		inverse := make(map[int]string)
		input := strings.Split(main[0], " ")
		for _, i := range input {
			i := SortString(i)
			switch len(i) {
			case 2:
				mapping[i] = 1
				inverse[1] = i
			case 3:
				mapping[i] = 7
				inverse[7] = i
			case 4:
				mapping[i] = 4
				inverse[4] = i
			case 7:
				mapping[i] = 8
				inverse[8] = i
			}
		}
		for _, i := range input {
			i := SortString(i)
			switch len(i) {
			case 5:
				if segmentContains(i, inverse[1]) == 2 {
					mapping[i] = 3
				} else {
					if segmentContains(i, inverse[4]) == 2 {
						mapping[i] = 2
					} else {
						mapping[i] = 5
					}
				}
			case 6:
				if segmentContains(i, inverse[4]) == 4 {
					mapping[i] = 9
				} else {
					if segmentContains(i, inverse[1]) == 2 {
						mapping[i] = 0
					} else {
						mapping[i] = 6
					}
				}
			}
		}
		outputvalue := ""
		output := strings.Split(main[1], " ")
		for _, o := range output {
			o := SortString(o)
			outputvalue = fmt.Sprintf("%v%v", outputvalue, mapping[o])
		}
		outputnumber, _ := strconv.Atoi(outputvalue)
		sum += outputnumber
	}
	return sum
}

func Part1(r io.ReadSeeker) int {
	return Calc(r)
}

func Part2(r io.ReadSeeker) int {
	return Calc2(r)
}

func main() {
	test1 := strings.NewReader("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	test2 := strings.NewReader(`be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`)
	fmt.Println("Part1 - test1", Part1(test1))
	fmt.Println("Part2 - test1", Part2(test1))
	fmt.Println("Part1 - test2", Part1(test2))
	fmt.Println("Part2 - test2", Part2(test2))

	input, err := os.Open("2021/8/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))
	fmt.Println("Part2 - puzzle", Part2(input))
}
