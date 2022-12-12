package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	Add     = "+ %d"
	Times   = "* %d"
	Squared = "* old"
)

type Jungle struct {
	monkeys       []Monkey
	commonDivisor int
}

type Monkey struct {
	items      []int
	operation  string
	operand    int
	divisor    int
	trueThrow  int
	falseThrow int
}

func Parse(s string) Jungle {
	scan := bufio.NewScanner(strings.NewReader(s))
	monkeys := []Monkey{}
	for scan.Scan() {
		scan.Scan()
		m := Monkey{}
		line := scan.Text()
		items := strings.Split(line[18:], ", ")
		for _, i := range items {
			j, _ := strconv.Atoi(i)
			m.items = append(m.items, j)
		}
		scan.Scan()
		opstr := scan.Text()
		opstr = opstr[23:]
		if opstr == Squared {
			m.operation = Squared
		} else {
			m.operation = Add
			_, err := fmt.Sscanf(opstr, Add, &m.operand)
			if err != nil {
				m.operation = Times
				fmt.Sscanf(opstr, Times, &m.operand)
			}
		}

		scan.Scan()
		fmt.Sscanf(scan.Text(), "  Test: divisible by %d", &m.divisor)

		scan.Scan()
		fmt.Sscanf(scan.Text(), "    If true: throw to monkey %d", &m.trueThrow)

		scan.Scan()
		fmt.Sscanf(scan.Text(), "    If false: throw to monkey %d", &m.falseThrow)

		scan.Scan()
		monkeys = append(monkeys, m)
	}
	j := Jungle{monkeys, 1}
	for _, n := range monkeys {
		j.commonDivisor *= n.divisor
	}
	return j
}

func Part(jungle Jungle, iterations, stressDivisor int) int {
	inspections := make([]int, len(jungle.monkeys), len(jungle.monkeys))
	for i := 0; i < iterations; i++ {
		for n, m := range jungle.monkeys {
			for _, j := range m.items {
				inspections[n]++
				jnew := j
				if m.operation == Add {
					jnew += m.operand
				} else if m.operation == Times {
					jnew *= m.operand
				} else {
					jnew *= j
				}
				jnew /= stressDivisor
				jnew %= jungle.commonDivisor

				throwTo := m.falseThrow
				if jnew%m.divisor == 0 {
					throwTo = m.trueThrow
				}
				jungle.monkeys[throwTo].items = append(jungle.monkeys[throwTo].items, jnew)
			}
			jungle.monkeys[n].items = nil
		}
	}
	sort.Ints(inspections)
	n := len(inspections)
	return inspections[n-2] * inspections[n-1]
}

func Part1(jungle Jungle) int {
	return Part(jungle, 20, 3)
}

func Part2(jungle Jungle) int {
	return Part(jungle, 10000, 1)
}

//go:embed input.txt
var puzzleinput string

//go:embed example.txt
var examplestr string

func main() {
	example := Parse(examplestr)
	input := Parse(puzzleinput)

	fmt.Println("Part1", Part1(example), Part1(input))

	// need to reparse to reset to initial state
	example = Parse(examplestr)
	input = Parse(puzzleinput)
	fmt.Println("Part2", Part2(example), Part2(input))
}
