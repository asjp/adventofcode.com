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

	x, y := 0, 0
	heading := 90

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			command := string(line[0])
			amount, _ := strconv.Atoi(line[1:])
			switch command {
			case "L":
				heading = (360+heading - amount) % 360
			case "R":
				heading = (heading + amount) % 360
			case "F":
				switch heading {
				case 0:
					y -= amount
				case 90:
					x += amount
				case 180:
					y += amount
				case 270:
					x -= amount
				}
			case "N":
				y -= amount
			case "S":
				y += amount
			case "E":
				x += amount
			case "W":
				x -= amount
			}
		}
		fmt.Println(line, heading, x, y)
	}

	md := math.Abs(float64(x)) + math.Abs(float64(y))
	fmt.Printf("%.1f\n", md)
}