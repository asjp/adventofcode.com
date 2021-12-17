package main

import (
	"bufio"
	"fmt"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"io"
	"os"
	"strings"
	"time"

	"gonum.org/v1/gonum/graph"
)

func ReadInput(r io.ReadSeeker) graph.Graph {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	result := simple.NewWeightedUndirectedGraph()
	for s.Scan() {
		s := s.Text()
		for i := range s {
			n := result.NewNode()
			if i > 0 {

			}
		}
		result = append(result, line)
	}
	return result
}

func Part1(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	result := path.DijkstraFrom(ReadInput(r))
	//fmt.Println(cost)
	return result
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
	test1 := strings.NewReader(`1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`)

	test2 := strings.NewReader(`19999
19111
11191`)

	input, err := os.Open("2021/15/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	expect(40, Part1(test1), "Part1 - test1")
	expect(14, Part1(test2), "Part1 - test2")
	fmt.Println("Part1 - puzzle", Part1(input))

}
