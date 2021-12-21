package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"
)

type Input struct {
	algo  map[int]struct{}
	image Image
}

type Image struct {
	pixels      map[Coord]struct{}
	size        int // assume images are always square
	infinityLit bool
}

type Coord struct {
	x, y int
}

func ReadInput(r io.ReadSeeker) Input {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	input := Input{
		make(map[int]struct{}),
		Image{
			make(map[Coord]struct{}),
			0,
			false,
		},
	}

	s.Scan()
	for i, c := range s.Text() {
		if c == '#' {
			input.algo[i] = struct{}{}
		}
	}
	s.Scan()
	for s.Scan() {
		for x, c := range s.Text() {
			if c == '#' {
				input.image.pixels[Coord{x, input.image.size}] = struct{}{}
			}
		}
		input.image.size++
	}
	return input
}

func InfiniteLit(image Image, x, y int) bool {
	if x < 0 || y < 0 || x >= image.size || y >= image.size {
		return image.infinityLit
	}
	return false
}

func CalcKey(image Image, x, y int) int {
	k := 0
	for ay := y - 1; ay < y+2; ay++ {
		for ax := x - 1; ax < x+2; ax++ {
			k <<= 1
			if _, isSet := image.pixels[Coord{ax, ay}]; isSet || InfiniteLit(image, ax, ay) {
				k++
			}
		}
	}
	return k
}

func Iterate(n int, input Input) int {
	im := input.image
	for i := 0; i < n; i++ {
		newImage := Image{
			make(map[Coord]struct{}),
			im.size + 2, // grow 1 in each direction
			im.infinityLit,
		}
		for x := -1; x < im.size+1; x++ {
			for y := -1; y < im.size+1; y++ {
				key := CalcKey(im, x, y)
				if _, isSet := input.algo[key]; isSet {
					newImage.pixels[Coord{x + 1, y + 1}] = struct{}{}
				}
			}
		}
		if im.infinityLit {
			if _, isSet := input.algo[511]; !isSet {
				newImage.infinityLit = false
			}
		} else {
			if _, isSet := input.algo[0]; isSet {
				newImage.infinityLit = true
			}
		}
		im = newImage
	}

	return len(im.pixels)
}

func Part1(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	return Iterate(2, ReadInput(r))
}

func Part2(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	return Iterate(50, ReadInput(r))
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
	testinput := FileReader("test.txt")
	expect(35, Part1(testinput), "Part1 - test")

	input := FileReader("input.txt")
	fmt.Println("Part1 - puzzle", Part1(input))

	expect(3351, Part2(testinput), "Part2 - test")
	fmt.Println("Part2 - puzzle", Part2(input))
}
