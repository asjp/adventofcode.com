package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

type Input struct {
	x1, y1, x2, y2 int
}

func ReadInput(r io.ReadSeeker) Input {
	_, _ = r.Seek(0, io.SeekStart)
	var input Input
	fmt.Fscanf(r, "target area: x=%d..%d, y=%d..%d", &input.x1,
		&input.x2, &input.y1, &input.y2)
	fmt.Printf("%+v\n", input)
	return input
}

func Between(a, x, b int) bool {
	return a <= x && x <= b
}

func Iterate(input Input) (int, int) {
	// d = vt + 1/2 at^2
	// solve for x
	//   a = -1
	//   v = (d + 1/2 t^2)/t => v = (d+1/2)/t
	//   if at t0: v = v0, then when v = 0, t = v0
	//   => for d at v=0, v0 = (d + 1/2)/v0
	//   => v0 = sqrt(d + 1/2)
	Vx_lower := int(math.Sqrt(float64(input.x1) + 0.5))
	// this is wrong for the upper bound
	// because it assumes constant acceleration
	// but here we have discrete steps rather than continuous acc
	// so apply a fudge factor
	//Vx_upper := int(math.Sqrt(float64(input.x2)+0.5)) + 10
	Vx_upper := input.x2 + 1

	Dy_max := 0 // max height
	Ay := -1    // gravity
	Ax := -1    // drag
	n := 0
	for Vx := Vx_lower; Vx <= Vx_upper; Vx++ {
		for Vy := input.y1 - 1; Vy <= 100; Vy++ {
			x, y := 0, 0
			tVx, tVy := Vx, Vy // initial velocity
			tDy_max := 0
			for y > input.y1 && x < input.x2 {
				y += tVy
				tVy += Ay
				x += tVx
				if tVx > 0 {
					tVx += Ax
				}
				if y > tDy_max {
					tDy_max = y
				}
				if Between(input.y1, y, input.y2) && Between(input.x1, x, input.x2) {
					n++
					if tDy_max > Dy_max {
						Dy_max = tDy_max
					}
					//fmt.Println(Vx, Vy, tDy_max, n)
					break
				}
			}
		}
	}

	return Dy_max, n
}

func Part1(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	result, _ := Iterate(ReadInput(r))
	return result
}

func Part2(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	_, result := Iterate(ReadInput(r))
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
	expect(45, Part1(strings.NewReader("target area: x=20..30, y=-10..-5")), "Part1 - test1")
	input, err := os.Open("2021/17/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(input)
	fmt.Println("Part1 - puzzle", Part1(input))

	expect(112, Part2(strings.NewReader("target area: x=20..30, y=-10..-5")), "Part2 - test1")
	fmt.Println("Part2 - puzzle", Part2(input))
}
