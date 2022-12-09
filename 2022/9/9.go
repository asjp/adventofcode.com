package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
)

func PuzzleInput() io.ReadSeeker {
	_, filename, _, _ := runtime.Caller(0)
	today := path.Dir(filename)
	f, err := os.Open(path.Join(today, "input.txt"))
	if err != nil {
		fmt.Println(err)
	}
	return f
}

type Coord struct {
	x, y int
}

func Part1(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	visited := map[Coord]struct{}{{0, 0}: {}}
	var head, tail Coord
	dir, steps := "", 0
	for scan.Scan() {
		fmt.Sscanf(scan.Text(), "%s %d", &dir, &steps)
		for i := 0; i < steps; i++ {
			switch dir {
			case "U":
				head.y++
				break
			case "D":
				head.y--
				break
			case "R":
				head.x++
				break
			case "L":
				head.x--
				break
			}
			tail = UpdateTail(head, tail)
			visited[tail] = struct{}{}
		}
	}
	return len(visited)
}

func UpdateTail(head, tail Coord) Coord {
	if head.x == tail.x {
		if tail.y < head.y-1 {
			tail.y++
		} else if tail.y > head.y+1 {
			tail.y--
		}
	} else if head.y == tail.y {
		if tail.x < head.x-1 {
			tail.x++
		} else if tail.x > head.x+1 {
			tail.x--
		}
	} else if head.x > tail.x+1 {
		tail.x++
		if head.y > tail.y {
			tail.y++
		} else {
			tail.y--
		}
	} else if head.x < tail.x-1 {
		tail.x--
		if head.y > tail.y {
			tail.y++
		} else {
			tail.y--
		}
	} else if head.y > tail.y+1 {
		tail.y++
		if head.x > tail.x {
			tail.x++
		} else {
			tail.x--
		}
	} else if head.y < tail.y-1 {
		tail.y--
		if head.x > tail.x {
			tail.x++
		} else {
			tail.x--
		}
	}

	return tail
}

func Part2(r io.ReadSeeker) int {
	_, _ = r.Seek(0, io.SeekStart)
	scan := bufio.NewScanner(r)
	visited := map[Coord]struct{}{{0, 0}: {}}
	rope := [10]Coord{}
	dir, steps := "", 0
	for scan.Scan() {
		fmt.Sscanf(scan.Text(), "%s %d", &dir, &steps)
		for i := 0; i < steps; i++ {
			switch dir {
			case "U":
				rope[0] = Coord{rope[0].x, rope[0].y + 1}
				break
			case "D":
				rope[0] = Coord{rope[0].x, rope[0].y - 1}
				break
			case "R":
				rope[0] = Coord{rope[0].x + 1, rope[0].y}
				break
			case "L":
				rope[0] = Coord{rope[0].x - 1, rope[0].y}
				break
			}
			for k := 0; k < 9; k++ {
				rope[k+1] = UpdateTail(rope[k], rope[k+1])
			}
			visited[rope[9]] = struct{}{}
		}
	}
	return len(visited)
}

func main() {
	example := strings.NewReader(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)

	example2 := strings.NewReader(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`)

	input := PuzzleInput()
	fmt.Println("Part1", Part1(example), Part1(input))
	fmt.Println("Part2", Part2(example2), Part2(input))
}
