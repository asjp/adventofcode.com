package main

import (
	"strings"
	"testing"
)

func BenchmarkPart2MapExample(b *testing.B) {
	test := ParseGrid(strings.NewReader(`30373
25512
65332
33549
35390`))
	for i := 0; i < b.N; i++ {
		Part2(test)
	}
}

func BenchmarkPart2MapPuzzleInput(b *testing.B) {
	test := ParseGrid(PuzzleInput())
	for i := 0; i < b.N; i++ {
		Part2(test)
	}
}

func BenchmarkPart2SliceExample(b *testing.B) {
	test := ParseGridSlice(strings.NewReader(`30373
25512
65332
33549
35390`))
	for i := 0; i < b.N; i++ {
		Part2Slice(test)
	}
}

func BenchmarkPart2SlicePuzzleInput(b *testing.B) {
	test := ParseGridSlice(PuzzleInput())
	for i := 0; i < b.N; i++ {
		Part2Slice(test)
	}
}
