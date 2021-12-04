package main

import (
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Min, Max int
	Letter rune
	Password string
	LetterStr string
}

func main() {
	f, _ := os.Open("2/2.input.txt")
	valid := 0
	for {
		var item Item
		n, _ := fmt.Fscanf(f, "%d-%d %c: %s\n", &item.Min, &item.Max, &item.Letter, &item.Password)
		if n == 0 {
			break
		}
		item.LetterStr = string(item.Letter)
		c := strings.Count(item.Password, item.LetterStr)
		if c >= item.Min && c <= item.Max {
			valid++
			fmt.Println("VALID ", item)
		} else {
			fmt.Println("INVALID ", item)
		}
	}
	fmt.Println(valid)
}
