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
	transform Transform
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
		{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, // 1
		{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}},
		{{1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
		{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}},
		{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}}, // 5
		{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}},
		{{0, 1, 0}, {1, 0, 0}, {0, 0, -1}},
		{{0, 0, -1}, {1, 0, 0}, {0, -1, 0}},
		{{-1, 0, 0}, {0, -1, 0}, {0, 0, 1}}, // 9
		{{-1, 0, 0}, {0, 0, -1}, {0, -1, 0}},
		{{-1, 0, 0}, {0, 1, 0}, {0, 0, -1}},
		{{-1, 0, 0}, {0, 0, 1}, {0, 1, 0}},
		{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}}, // 13
		{{0, 0, 1}, {-1, 0, 0}, {0, -1, 0}},
		{{0, -1, 0}, {-1, 0, 0}, {0, 0, -1}},
		{{0, 0, -1}, {-1, 0, 0}, {0, 1, 0}},
		{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}}, // 17
		{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
		{{0, 0, 1}, {0, -1, 0}, {1, 0, 0}},
		{{0, -1, 0}, {0, 0, -1}, {1, 0, 0}},
		{{0, 0, -1}, {0, -1, 0}, {-1, 0, 0}}, // 21
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

func RotateCoord(c Coord, m Matrix3x3) Coord {
	return Coord{
		c.x*m[0][0] + c.y*m[0][1] + c.z*m[0][2],
		c.x*m[1][0] + c.y*m[1][1] + c.z*m[1][2],
		c.x*m[2][0] + c.y*m[2][1] + c.z*m[2][2],
	}
}

type Matrix3x3 [3][3]int

type Transform struct {
	rotate    Matrix3x3
	translate Coord
}

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

func Keys(m map[int]struct{}) string {
	s := []string{}
	for k, _ := range m {
		s = append(s, fmt.Sprintf("%d", k))
	}
	return strings.Join(s, ",")
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
		found := false
		for j, _ := range remaining {
			for i, _ := range normalised {
				if transform, matched := MatchScanner(input.s[i], input.s[j]); matched {
					//fmt.Println(i, "->", j, "T =", transform)
					normalised[j] = struct{}{}
					input.s[j].transform = transform
					NormaliseBeacons(input.s[j].b, transform)
					delete(remaining, j)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			//fmt.Println("Normalised", Keys(normalised))
			fmt.Println("Remaining", Keys(remaining))
		} else {
			fmt.Println("Unable to normalise", Keys(remaining))
			break
		}
	}
}

func NormaliseBeacons(beacons []Beacon, transform Transform) {
	for i, b := range beacons {
		c := RotateCoord(b.Coord, transform.rotate)
		c.x += transform.translate.x
		c.y += transform.translate.y
		c.z += transform.translate.z
		beacons[i] = Beacon{c}
	}
}

func MatchingDistances(from, to Scanner) []BeaconPairMatch {
	result := make([]BeaconPairMatch, 0)
	for from_dist, from_pair := range from.distances {
		for to_dist, to_pair := range to.distances {
			if from_dist == to_dist {
				result = append(result, BeaconPairMatch{from_pair, to_pair})
			}
		}
	}
	return result
}

func MatchScanner(from, to Scanner) (Transform, bool) {
	beaconMatches := MatchingDistances(from, to)
	if len(beaconMatches) < 12 {
		return Transform{}, false
	}

	transforms := make(map[Transform]int)

	for _, bm := range beaconMatches {
		for _, r := range allRotations {
			r1 := RotateCoord(to.b[bm.to.b1].Coord, r)
			r2 := RotateCoord(to.b[bm.to.b2].Coord, r)
			d1 := DiffVec(from.b[bm.from.b1].Coord, r1)
			d2 := DiffVec(from.b[bm.from.b2].Coord, r2)
			if d1 == d2 {
				tr := Transform{r, d1}
				if n, ok := transforms[tr]; ok {
					transforms[tr] = n + 1
				} else {
					transforms[tr] = 1
				}
			}
		}
	}

	returnTransform, maxN := Transform{}, 0
	for t, n := range transforms {
		if n > maxN {
			returnTransform = t
			maxN = n
		}
	}
	if maxN >= 12 {
		return returnTransform, true
	}

	return Transform{}, false
}

// returns a - b
func DiffVec(a, b Coord) Coord {
	return Coord{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z,
	}
}

func Part1(input Input) int {
	defer timeTrack(time.Now())
	uniqueBeacons := make(map[Beacon]struct{})
	for _, s := range input.s {
		for _, b := range s.b {
			uniqueBeacons[b] = struct{}{}
		}
	}
	return len(uniqueBeacons)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Parse(r io.ReadSeeker) Input {
	defer timeTrack(time.Now())
	input := ReadInput(r)
	Normalise(input)
	return input
}

func Part2(input Input) int {
	defer timeTrack(time.Now())
	max := 0
	for i, si := range input.s {
		for j, sj := range input.s {
			if i == j {
				continue
			}
			d := Abs(si.transform.translate.x-sj.transform.translate.x) +
				Abs(si.transform.translate.y-sj.transform.translate.y) +
				Abs(si.transform.translate.z-sj.transform.translate.z)
			if d > max {
				max = d
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

	expectMatrix3x3(Matrix3x3{{130, 120, 240}, {51, 47, 73}, {35, 33, 45}},
		Multiply3x3(Matrix3x3{{10, 20, 10}, {4, 5, 6}, {2, 3, 5}},
			Matrix3x3{{3, 2, 4}, {3, 3, 9}, {4, 4, 2}}))

	test := Parse(FileReader("test.txt"))
	expect(79, Part1(test), "Part1 - test")

	input := Parse(FileReader("input.txt"))
	fmt.Println("Part1 - puzzle", Part1(input))

	expect(3621, Part2(test), "Part2 - test")
	fmt.Println("Part2 - puzzle", Part2(input))
}
