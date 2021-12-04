package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Part1(input io.Reader) int {
	scan := bufio.NewScanner(input)
	scan.Scan()
	start := strings.Split(scan.Text(), ",")
	nums := make([]int, len(start))
	turns := 1
	for _, ns := range start {
		n, _ := strconv.Atoi(ns)
		nums = append(nums, n)
		turns++
	}

	for turns < 2021 {
		last := nums[len(nums)-1]
		next := 0
		for i := len(nums) - 2; i >= 0; i-- {
			if nums[i] == last {
				next = (len(nums) - 1) - i
				break
			}
		}
		//fmt.Println(turns, " ", next)
		nums = append(nums, next)
		turns++
	}

	return nums[len(nums)-1]
}

func Part2(input io.Reader) int {
	scan := bufio.NewScanner(input)
	scan.Scan()
	start := strings.Split(scan.Text(), ",")
	nums := make(map[int]int)
	turns := 1
	for _, ns := range start {
		n, _ := strconv.Atoi(ns)
		nums[n] = turns
		turns++
	}
	next := 0

	fmt.Println(nums)
	for turns < 30000000 {
		//fmt.Println(next, " ", nums, " turns = ", turns)
		last := next
		if lastPos, ok := nums[next]; ok {
			spoken := turns - lastPos
			next = spoken
		} else {
			next = 0
		}
		nums[last] = turns
		turns++
	}

	return next
}

func main() {
	fmt.Println("Part1")
	fmt.Println(Part1(strings.NewReader("0,3,6")))
	fmt.Println(Part1(strings.NewReader("1,3,2")))
	fmt.Println(Part1(strings.NewReader("2,1,3")))
	fmt.Println(Part1(strings.NewReader("1,2,3")))
	fmt.Println(Part1(strings.NewReader("2,0,1,9,5,19")))

	fmt.Println("Part2")
	fmt.Println(Part2(strings.NewReader("0,3,6")))
	fmt.Println(Part2(strings.NewReader("1,3,2")))
	fmt.Println(Part2(strings.NewReader("2,1,3")))
	fmt.Println(Part2(strings.NewReader("1,2,3")))
	fmt.Println(Part2(strings.NewReader("2,0,1,9,5,19")))
}
