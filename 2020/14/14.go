package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Part1(input io.Reader) int64 {
	mem := make(map[int64]int64)

	scanner := bufio.NewScanner(input)
	andMask := int64(0)
	orMask := int64(0)
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), " ")
		left := lineParts[0]
		right := lineParts[2]
		fmt.Println(left, " ", right)
		if left == "mask" {
			andMask, _ = strconv.ParseInt(strings.ReplaceAll(right, "X", "1"), 2, 64)
			orMask, _ = strconv.ParseInt(strings.ReplaceAll(right, "X", "0"), 2, 64)
			fmt.Printf("%#.36b %#.36b\n", andMask, orMask)
		} else {
			address, _ := strconv.ParseInt(left[4:len(left)-1], 10, 64)
			value, _ := strconv.ParseInt(right, 10, 64)
			mem[address] = (value & andMask) | orMask
			fmt.Println(address, "=", mem[address])
		}
	}
	sum := int64(0)
	for _, v := range mem {
		sum += v
	}
	return sum
}

func Part2(input io.Reader) int64 {
	mem := make(map[int64]int64)

	scanner := bufio.NewScanner(input)
	orMask := int64(0)
	floatingMask := int64(0)
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), " ")
		left := lineParts[0]
		right := lineParts[2]
		fmt.Println(left, " ", right)
		if left == "mask" {
			orMask, _ = strconv.ParseInt(strings.ReplaceAll(right, "X", "0"), 2, 64)
			floatingMask, _ = strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(right, "1", "0"), "X", "1"), 2, 64)
			fmt.Printf("%#.36b %#.36b\n", orMask, floatingMask)
		} else {
			address, _ := strconv.ParseInt(left[4:len(left)-1], 10, 64)
			value, _ := strconv.ParseInt(right, 10, 64)
			address = address | orMask
			updateMem(mem, address, value, floatingMask, 0)
		}
	}
	sum := int64(0)
	for _, v := range mem {
		sum += v
	}
	return sum
}

func updateMem(mem map[int64]int64, address, value, floatingMask, bitNumber int64) {
	if bitNumber > 35 {
		return
	}
	if floatingMask & (1<<bitNumber) != 0 {
		fmt.Println("updating ", bitNumber)
		a := address & ^(1<<bitNumber)
		mem[a] = value
		fmt.Printf("mem[%v] = %v\n", a, value)
		updateMem(mem, a, value, floatingMask, bitNumber+1)
		a = address | (1<<bitNumber)
		mem[a] = value
		fmt.Printf("mem[%v] = %v\n", a, value)
		updateMem(mem, a, value, floatingMask, bitNumber+1)
	} else {
		updateMem(mem, address, value, floatingMask, bitNumber+1)
	}
}

func main() {
	fmt.Println("Part1")
	f, _ := os.Open("14/input.txt")
	result := Part1(f)
	fmt.Println(result)

	fmt.Println("Part2")
	f, _ = os.Open("14/input.txt")
	result = Part2(f)
	fmt.Println(result)
}