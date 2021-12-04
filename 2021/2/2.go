package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Part1(r io.Reader) int {
	depth, pos := 0, 0
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		s := scan.Text()
		i := strings.Split(s, " ")
		n, _ := strconv.Atoi(i[1])
		switch i[0] {
		case "forward":
			{
				pos += n
				break
			}
		case "up":
			{
				depth -= n
				break
			}
		case "down":
			{
				depth += n
				break
			}
		}
	}
	return depth * pos
}

func Part2(r io.Reader) int {
	depth, pos, aim := 0, 0, 0
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		s := scan.Text()
		i := strings.Split(s, " ")
		n, _ := strconv.Atoi(i[1])
		switch i[0] {
		case "forward":
			{
				pos += n
				depth += (aim * n)
				break
			}
		case "up":
			{
				aim -= n
				break
			}
		case "down":
			{
				aim += n
				break
			}
		}
	}
	return depth * pos
}

func main() {
	test := strings.NewReader(`forward 5
down 5
forward 8
up 3
down 8
forward 2
`)
	fmt.Println("Part1", Part2(test))
	input, err := os.Open("2021/1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1", Part2(input))
	/*
		fmt.Println("Part2", Part2(test), Part2(input))
	*/
}
