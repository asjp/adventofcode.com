package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Input struct {
	fish []*Snailfish
}

type Snailfish struct {
	first, second, parent *Snailfish
	num                   int
}

func ReadInput(r io.ReadSeeker) Input {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	input := Input{}
	for s.Scan() {
		line := s.Text()
		fish := ParseFish0(line)
		input.fish = append(input.fish, fish)
	}
	return input
}

func ParseFish0(s string) *Snailfish {
	return ParseFish(s, 0, nil)
}

func ParseFish(s string, depth int, parent *Snailfish) *Snailfish {
	//fmt.Println(s)
	fish := &Snailfish{}
	fish.parent = parent
	if s[0] == '[' {
		end := findEnd(s)
		if end < len(s) {
			fish.first = ParseFish(s[0:end], depth+1, fish)
			fish.second = ParseFish(s[end+1:], depth+1, fish)
		} else {
			fish = ParseFish(s[1:len(s)-1], depth+1, parent)
		}
	} else {
		part := strings.SplitN(s, ",", 2)
		if len(part) == 2 {
			fish.first = ParseFish(part[0], depth+1, fish)
			fish.second = ParseFish(part[1], depth+1, fish)
		} else {
			fish.num, _ = strconv.Atoi(s)
		}
	}
	return fish
}

func findEnd(s string) int {
	e := 0
	for i, c := range s {
		if c == '[' {
			e++
		}
		if c == ']' {
			e--
			if e == 0 {
				return i + 1
			}
		}
	}
	return len(s)
}

func (s *Snailfish) String() string {
	if s.HasChildren() {
		return fmt.Sprintf("[%s,%s]", s.first.String(), s.second.String())
	}
	return fmt.Sprint(s.num)
}

func FindNumLeft(fish *Snailfish) *Snailfish {
	find := fish
	for find.parent != nil {
		if find.parent.first != find {
			find = find.parent.first
			for find.HasChildren() {
				find = find.second
			}
			return find
		}
		find = find.parent
	}
	return nil
}

func FindNumRight(fish *Snailfish) *Snailfish {
	find := fish
	for find.parent != nil {
		if find.parent.second != find {
			find = find.parent.second
			for find.HasChildren() {
				find = find.first
			}
			return find
		}
		find = find.parent
	}
	return nil
}

func (fish *Snailfish) HasChildren() bool {
	return fish.first != nil
}

func (fish *Snailfish) Zero() {
	// keep parent unchanged
	fish.first, fish.second, fish.num = nil, nil, 0
}

func Reduce(fish *Snailfish, depth int) bool {
	if Explode(fish, depth) {
		return true
	}
	return Split(fish)
}

func Explode(fish *Snailfish, depth int) bool {
	//fmt.Println(fish)
	if depth == 3 && fish.HasChildren() {
		if fish.first.HasChildren() {
			p := FindNumLeft(fish.first)
			if p != nil {
				p.num += fish.first.first.num
			}
			p = FindNumRight(fish.first)
			if p != nil {
				p.num += fish.first.second.num
			}
			fish.first.Zero()
			return true
		} else if fish.second.HasChildren() {
			p := FindNumLeft(fish.second)
			if p != nil {
				p.num += fish.second.first.num
			}
			p = FindNumRight(fish.second)
			if p != nil {
				p.num += fish.second.second.num
			}
			fish.second.Zero()
			return true
		}
	}
	if fish.HasChildren() {
		if Explode(fish.first, depth+1) {
			return true
		}
		if Explode(fish.second, depth+1) {
			return true
		}
	}
	return false
}

func Split(fish *Snailfish) bool {
	if !fish.HasChildren() {
		if fish.num >= 10 {
			fish.first = &Snailfish{nil, nil, fish, fish.num / 2}
			fish.second = &Snailfish{nil, nil, fish, int(math.Ceil(float64(fish.num) / 2.))}
			fish.num = 0
			return true
		}
	} else {
		if Split(fish.first) {
			return true
		}
		if Split(fish.second) {
			return true
		}
	}
	return false
}
func Reduce0(fish *Snailfish) *Snailfish {
	for Reduce(fish, 0) {
	}
	return fish
}

func FishSum(r io.ReadSeeker) *Snailfish {
	input := ReadInput(r)
	return FishSlice(input.fish)
}

func FishSlice(fishes []*Snailfish) *Snailfish {
	var sum *Snailfish
	for _, f := range fishes {
		if sum == nil {
			sum = f
		} else {
			sum = &Snailfish{
				first:  sum,
				second: f,
			}
			sum.first.parent = sum
			sum.second.parent = sum
			Reduce0(sum)
		}
	}
	return sum
}

func Magnitude(sum *Snailfish) int {
	if sum.HasChildren() {
		return 3*Magnitude(sum.first) + 2*Magnitude(sum.second)
	}
	return sum.num
}

func Part1(r io.ReadSeeker) int {
	return Magnitude(FishSum(r))
}

func Part2(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	input := ReadInput(r).fish
	max := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			// reparse each pair we try as summing modifies the input
			f1, f2 := ParseFish0(input[i].String()), ParseFish0(input[j].String())
			mag := Magnitude(FishSlice([]*Snailfish{f1, f2}))
			if mag > max {
				max = mag
			}
		}
	}
	return max
}

func timeTrack(start time.Time) {
	fmt.Printf("(%10s) ", time.Since(start))
}

func expect(expected, actual int, msg string) {
	if expected == actual {
		fmt.Printf("\033[1;32mOK\033[0m %d", actual)
	} else {
		fmt.Printf("\033[1;31mFAIL\033[0m (expected %d, actual %d)", expected, actual)
	}
	fmt.Println(" ", msg)
}

func expectStr(expected string, actual fmt.Stringer) {
	if expected == actual.String() {
		fmt.Printf("\033[1;32mOK\033[0m %s\n", actual)
	} else {
		fmt.Printf("\033[1;31mFAIL\033[0m (expected %s, actual %s)\n", expected, actual)
	}
}

func main() {
	expectStr("[[[[0,9],2],3],4]", Reduce0(ParseFish0("[[[[[9,8],1],2],3],4]")))
	expectStr("[7,[6,[5,[7,0]]]]", Reduce0(ParseFish0("[7,[6,[5,[4,[3,2]]]]]")))
	expectStr("[[6,[5,[7,0]]],3]", Reduce0(ParseFish0("[[6,[5,[4,[3,2]]]],1]")))
	expectStr("[[3,[2,[8,0]]],[9,[5,[7,0]]]]", Reduce0(ParseFish0("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")))

	expectStr("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", Reduce0(ParseFish0("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")))

	expectStr("[[[[5,0],[7,4]],[5,5]],[6,6]]", FishSum(strings.NewReader(`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]`)))

	expectStr("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", FishSum(strings.NewReader(`[[[[4,3],4],4],[7,[[8,4],9]]]
[1,1]`)))

	expectStr("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		FishSum(strings.NewReader(`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`)))

	expect(29, Magnitude(ParseFish0("[9,1]")), "[9,1]")
	expect(3488, Magnitude(ParseFish0("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")), "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")

	finalExample := strings.NewReader(`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`)
	expect(4140, Part1(finalExample), "Part1 - final example")

	input, err := os.Open("2021/18/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))

	expect(3993, Part2(finalExample), "Part2 - final example")
	fmt.Println("Part2 - puzzle", Part2(input))
}
