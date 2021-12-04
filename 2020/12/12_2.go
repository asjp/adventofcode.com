package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("12/input.txt")

	wx, wy := 10, -1
	sx, sy := 0, 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			command := string(line[0])
			amount, _ := strconv.Atoi(line[1:])
			switch command {
			case "L":
				switch amount {
				case 90:
					wx, wy = wy, -wx
				case 180:
					wx, wy = -wx, -wy
				case 270:
					wx, wy = -wy, wx
				}
			case "R":
				switch amount {
				case 90:
					wx, wy = -wy, wx
				case 180:
					wx, wy = -wx, -wy
				case 270:
					wx, wy = wy, -wx
				}

			case "F":
				sx += amount * wx
				sy += amount * wy
			case "N":
				wy -= amount
			case "S":
				wy += amount
			case "E":
				wx += amount
			case "W":
				wx -= amount
			}
		}
	}

	md := math.Abs(float64(sx)) + math.Abs(float64(sy))
	fmt.Printf("%.1f\n", md)
}