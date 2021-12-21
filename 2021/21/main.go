package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"
)

type Input struct {
	Player1, Player2           int
	Player1Score, Player2Score int
}

func ReadInput(r io.ReadSeeker) Input {
	_, _ = r.Seek(0, io.SeekStart)
	input := Input{}
	fmt.Fscanf(r, "Player 1 starting position: %d\n", &input.Player1)
	fmt.Fscanf(r, "Player 2 starting position: %d\n", &input.Player2)
	return input
}

type Dice struct {
	next int
}

func (d *Dice) RollN(num int) int {
	score := 0
	for i := 0; i < num; i++ {
		score += d.Roll()
	}
	return score
}

func (d *Dice) Roll() int {
	r := d.next%100 + 1
	d.next++
	return r
}

func Advance(p, amount int) int {
	return ((p+amount)-1)%10 + 1
}

func Iterate(input Input) int {
	dice := Dice{}
	for {
		input.Player1 = Advance(input.Player1, dice.RollN(3))
		input.Player1Score += input.Player1
		if input.Player1Score >= 1000 {
			return input.Player2Score * dice.next
		}

		input.Player2 = Advance(input.Player2, dice.RollN(3))
		input.Player2Score += input.Player2
		if input.Player2Score >= 1000 {
			return input.Player1Score * dice.next
		}
	}
}

type State struct {
	p1, p2           int
	p1score, p2score int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IterateDirac(input Input) int {
	states := map[State]int{
		{input.Player1, input.Player2, 0, 0}: 1,
	}

	// make nPr() for dirac dice
	// i.e. frequency distribution of possible sum of 3 rolls
	allRolls := map[int]int{}
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				allRolls[r1+r2+r3]++
			}
		}
	}

	var p1wins, p2wins int

	for len(states) > 0 {
		for turn := 1; turn <= 2; turn++ {
			nextStates := map[State]int{}
			for state, count := range states {
				for roll, rCount := range allRolls {
					p1, p2 := state.p1, state.p2
					p1score, p2score := state.p1score, state.p2score
					n := count * rCount

					if turn == 1 {
						p1 = Advance(p1, roll)
						p1score += p1
						if p1score >= 21 {
							p1wins += n
						} else {
							nextStates[State{p1, p2, p1score, p2score}] += n
						}
					} else {
						p2 = Advance(p2, roll)
						p2score += p2
						if p2score >= 21 {
							p2wins += n
						} else {
							nextStates[State{p1, p2, p1score, p2score}] += n
						}
					}
				}
			}
			states = nextStates
		}
	}
	return Max(p1wins, p2wins)
}

func Part1(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	return Iterate(ReadInput(r))
}

func Part2(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	// Note: running this on a 32-bit arch
	// will give the wrong result.
	// Could use int64, but CBA :)
	return IterateDirac(ReadInput(r))
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

func FileReader(name string) io.ReadSeeker {
	_, mypath, _, _ := runtime.Caller(0)
	input, err := os.Open(path.Join(path.Dir(mypath), name))
	if err != nil {
		fmt.Println(err)
	}
	return input
}

func main() {
	testdice := Dice{}
	expect(6, testdice.RollN(3), "test dice rolls")

	test := FileReader("test.txt")
	expect(739785, Part1(test), "Part1 - test")

	input := FileReader("input.txt")
	fmt.Println("Part1 - puzzle", Part1(input))

	expect(444356092776315, Part2(test), "Part2 - test")
	fmt.Println("Part2 - puzzle", Part2(input))
}
