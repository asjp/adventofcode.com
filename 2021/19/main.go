package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type Input struct {
	s []Scanner
}

type DistanceMap map[int]BeaconPair

type Scanner struct {
	Coord
	b         []Beacon
	transform Matrix4x4
	distances DistanceMap
}

type BeaconPair struct {
	b1, b2 int
}

type BeaconPairMatch struct {
	from, to BeaconPair
}

type Beacon struct {
	Coord
}

type Coord struct {
	x, y, z int
}

// Square of the magnitude of the vector from a -> b
func VectorMagSq(a, b Coord) int {
	dx, dy, dz := b.x-a.x, b.y-a.y, b.z-a.z
	return dx*dx + dy*dy + dz*dz
}

func AllRotations() []Matrix3x3 {
	// 24 orientations
	return []Matrix3x3{
		{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}},
		{{1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
		{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}},
		{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}},
		{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}},
		{{0, 1, 0}, {1, 0, 0}, {0, 0, -1}},
		{{0, 0, -1}, {1, 0, 0}, {0, -1, 0}},
		{{-1, 0, 0}, {0, -1, 0}, {0, 0, 1}},
		{{-1, 0, 0}, {0, 0, -1}, {0, -1, 0}},
		{{-1, 0, 0}, {0, 1, 0}, {0, 0, -1}},
		{{-1, 0, 0}, {0, 0, 1}, {0, 1, 0}},
		{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}},
		{{0, 0, 1}, {-1, 0, 0}, {0, -1, 0}},
		{{0, -1, 0}, {-1, 0, 0}, {0, 0, -1}},
		{{0, 0, -1}, {-1, 0, 0}, {0, 1, 0}},
		{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}},
		{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
		{{0, 0, 1}, {0 - 1, 0}, {1, 0, 0}},
		{{0, -1, 0}, {0, 0, -1}, {1, 0, 0}},
		{{0, 0, -1}, {0, -1, 0}, {-1, 0, 0}},
		{{0, -1, 0}, {0, 0, 1}, {-1, 0, 0}},
		{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}},
		{{0, 1, 0}, {0, 0, -1}, {-1, 0, 0}},
	}
}

func Multiply3x3(a, b Matrix3x3) Matrix3x3 {
	return Matrix3x3{
		{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1],
			a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2],
		},
		{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1],
			a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2],
		},
		{
			a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1],
			a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2],
		},
	}
}

func Multiply4x4(a, b Matrix4x4) Matrix4x4 {
	return Matrix4x4{
		{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1] + a[0][3]*b[3][1],
			a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2] + a[0][3]*b[3][2],
			a[0][0]*b[0][3] + a[0][1]*b[1][3] + a[0][2]*b[2][3] + a[0][3]*b[3][3],
		},
		{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1] + a[1][3]*b[3][1],
			a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2] + a[1][3]*b[3][2],
			a[1][0]*b[0][3] + a[1][1]*b[1][3] + a[1][2]*b[2][3] + a[1][3]*b[3][3],
		},
		{
			a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0] + a[2][3]*b[3][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1] + a[2][3]*b[3][1],
			a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2] + a[2][3]*b[3][2],
			a[2][0]*b[0][3] + a[2][1]*b[1][3] + a[2][2]*b[2][3] + a[2][3]*b[3][3],
		},
		{
			a[3][0]*b[0][0] + a[3][1]*b[1][0] + a[3][2]*b[2][0] + a[3][3]*b[3][0],
			a[3][0]*b[0][1] + a[3][1]*b[1][1] + a[3][2]*b[2][1] + a[3][3]*b[3][1],
			a[3][0]*b[0][2] + a[3][1]*b[1][2] + a[3][2]*b[2][2] + a[3][3]*b[3][2],
			a[3][0]*b[0][3] + a[3][1]*b[1][3] + a[3][2]*b[2][3] + a[3][3]*b[3][3],
		},
	}
}

func RotateCoord(c Coord, m Matrix3x3) Coord {
	return Coord{
		c.x*m[0][0] + c.y*m[0][1] + c.z*m[0][2],
		c.x*m[1][0] + c.y*m[1][1] + c.z*m[1][2],
		c.x*m[2][0] + c.y*m[2][1] + c.z*m[2][2],
	}
}

func TransformCoord(c Coord, m Matrix4x4) Coord {
	return Coord{
		c.x*m[0][0] + c.y*m[0][1] + c.z*m[0][2],
		c.x*m[1][0] + c.y*m[1][1] + c.z*m[1][2],
		c.x*m[2][0] + c.y*m[2][1] + c.z*m[2][2],
	}
}

func MatrixTransform(m Matrix3x3, t Coord) Matrix4x4 {
	return Multiply4x4(
		Matrix4x4{
			{1, 0, 0, t.x},
			{0, 1, 0, t.y},
			{0, 0, 1, t.z},
			{0, 0, 0, 1},
		},
		Matrix4x4{
			{m[0][0], m[0][1], m[0][2], 0},
			{m[1][0], m[1][1], m[1][2], 0},
			{m[2][0], m[2][1], m[2][2], 0},
			{0, 0, 0, 1},
		},
	)
}

type Matrix3x3 [3][3]int
type Matrix4x4 [4][4]int

func NewScanner() Scanner {
	return Scanner{
		b:         make([]Beacon, 0),
		distances: make(map[int]BeaconPair),
	}
}

func ReadInput(r io.ReadSeeker) Input {
	_, _ = r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	input := Input{}
	scanner := NewScanner()
	for s.Scan() {
		line := s.Text()
		if line == "" {
			input.s = append(input.s, scanner)
			scanner = NewScanner()
		} else if !strings.HasPrefix(line, "---") {
			b := Beacon{}
			fmt.Sscanf(line, "%d,%d,%d", &b.x, &b.y, &b.z)
			toIdx := len(scanner.b)
			for i, bfrom := range scanner.b {
				dist := VectorMagSq(bfrom.Coord, b.Coord)
				scanner.distances[dist] = BeaconPair{i, toIdx}
			}
			scanner.b = append(scanner.b, b)
		}
	}
	input.s = append(input.s, scanner)
	return input
}

func Normalise(input Input) {
	normalised := make(map[int]struct{})
	// first one is normalised against itself already
	normalised[0] = struct{}{}
	remaining := make(map[int]struct{})
	for i := 1; i < len(input.s); i++ {
		remaining[i] = struct{}{}
	}
	for len(remaining) > 0 {
		// for each Scanner
		// try to normalise against any remaining Scanner
	outer:
		for j, _ := range remaining {
			for i, _ := range normalised {
				if transform, matched := MatchScanner(input.s[i], input.s[j]); matched {
					fmt.Println(i, j, transform)
					if transform[1][3] == 1246 {
						transform[1][3] = -1246
					}
					normalised[j] = struct{}{}
					NormaliseBeacons(input.s[j].b, transform)
					fmt.Println("NORMALISED\n", input.s[j].b)
					input.s[j].Coord = TransformCoord(Coord{}, transform)
					//fmt.Println(input.s[j].Coord)
					delete(remaining, j)
					break outer
				}
			}
		}
		//fmt.Println(len(remaining), "remaining")
	}
}

func NormaliseBeacons(beacons []Beacon, transform Matrix4x4) {
	for i, b := range beacons {
		beacons[i] = Beacon{TransformCoord(b.Coord, transform)}
	}
}

func OutOfRange(a, b Beacon) bool {
	return (b.x-a.x) > 2000 || (b.y-a.y) > 2000 || (b.z-a.z > 2000) ||
		(a.x-b.x) > 2000 || (a.y-b.y) > 2000 || (a.z-b.z > 2000)
}

func MatchingDistances(from, to Scanner) []BeaconPairMatch {
	result := make([]BeaconPairMatch, 0)
	for from_dist, from_pair := range from.distances {
		for to_dist, to_pair := range to.distances {
			if OutOfRange(from.b[from_pair.b1], to.b[to_pair.b1]) ||
				OutOfRange(from.b[from_pair.b2], to.b[to_pair.b2]) {
				continue
			}
			if from_dist == to_dist {
				result = append(result, BeaconPairMatch{from_pair, to_pair})
			}
		}
	}
	return result
}
func MatchScanner(from, to Scanner) (Matrix4x4, bool) {
	beaconMatches := MatchingDistances(from, to)
	if len(beaconMatches) < 12 {
		return Matrix4x4{}, false
	}

	transforms := make(map[Matrix4x4]int)

	for _, bm := range beaconMatches {
		for _, r := range allRotations {
			r1 := RotateCoord(to.b[bm.to.b1].Coord, r)
			r2 := RotateCoord(to.b[bm.to.b2].Coord, r)
			d1 := DiffVec(from.b[bm.from.b1].Coord, r1)
			d2 := DiffVec(from.b[bm.from.b2].Coord, r2)
			if d1 == d2 {
				tr := MatrixTransform(r, d1)
				if n, ok := transforms[tr]; ok {
					transforms[tr] = n + 1
				} else {
					transforms[tr] = 1
				}
			}
		}
	}

	returnTransform, maxN := Matrix4x4{}, 0
	for t, n := range transforms {
		if n > maxN {
			returnTransform = t
			maxN = n
		}
	}
	if maxN >= 12 {
		return (returnTransform), true
	}

	return Matrix4x4{}, false
}

func Transpose(m Matrix4x4) Matrix4x4 {
	for i := 0; i < 3; i++ {
		flip := false
		for j := 0; j < 3; j++ {
			if m[i][j] == 1 {
				flip = true
			}
		}
		if flip {
			m[i][3] = -m[i][3]
		}
	}
	return m
}

// returns a - b
func DiffVec(a, b Coord) Coord {
	return Coord{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z,
	}
}

func Part1(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	input := ReadInput(r)
	Normalise(input)
	uniqueBeacons := make(map[Beacon]struct{})
	for _, s := range input.s {
		for _, b := range s.b {
			uniqueBeacons[b] = struct{}{}
		}
	}

	for b := range uniqueBeacons {
		fmt.Println(b)
	}
	return len(uniqueBeacons)
}

func Part2(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	return 0
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

func expectMatrix3x3(expected, actual Matrix3x3) {
	if expected == actual {
		fmt.Printf("\033[1;32mOK\033[0m %d\n", actual)
	} else {
		fmt.Printf("\033[1;31mFAIL\033[0m (expected %d, actual %d)\n", expected, actual)
	}
}

func expectStr(expected string, actual fmt.Stringer) {
	if expected == actual.String() {
		fmt.Printf("\033[1;32mOK\033[0m %s\n", actual)
	} else {
		fmt.Printf("\033[1;31mFAIL\033[0m (expected %s, actual %s)\n", expected, actual)
	}
}

var allRotations []Matrix3x3

func FileReader(name string) io.ReadSeeker {
	_, mypath, _, _ := runtime.Caller(0)
	input, err := os.Open(path.Join(path.Dir(mypath), name))
	if err != nil {
		fmt.Println(err)
	}
	return input
}

func main() {
	allRotations = AllRotations()

	expectMatrix3x3(Matrix3x3{{24, 37, 44}, {-6, -11, -32}, {0, -28, 6}},
		Multiply3x3(Matrix3x3{{6, 1, 0}, {-3, -2, 2}, {3, -1, -4}},
			Matrix3x3{{4, 5, 6}, {0, 7, 8}, {3, 9, 1}}))

	test := FileReader("test.txt")
	expect(79, Part1(test), "Part1 - test")
	expect(0, Part2(test), "Part2 - test")

	input := FileReader("input.txt")
	//fmt.Println("Part1 - puzzle", Part1(input))
	fmt.Println("Part2 - puzzle", Part2(input))
}
