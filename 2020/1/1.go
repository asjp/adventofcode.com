package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("1/1.input.txt")
	scanner := bufio.NewScanner(f)
	numbers := make([]int, 0)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, n)
	}

	fmt.Printf("%v\n", numbers)

	// part 1
	for _, a := range numbers {
		for _, b := range numbers {
			if a + b == 2020 {
				fmt.Printf("%d\n", a*b)
			}
		}
	}

	// part 2
	for _, a := range numbers {
		for _, b := range numbers {
			for _, c := range numbers {
				if a+b+c == 2020 {
					fmt.Printf("%d\n", a*b*c)
				}
			}
		}
	}

}
