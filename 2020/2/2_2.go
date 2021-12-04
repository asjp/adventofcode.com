package main

import (
	"fmt"
	"os"
)

type Item2 struct {
	First, Second int
	Letter rune
	Password string
	LetterStr string
}

func main() {
	f, _ := os.Open("2/2.input.txt")
	valid := 0
	for {
		var item Item2
		n, _ := fmt.Fscanf(f, "%d-%d %c: %s\n", &item.First, &item.Second, &item.Letter, &item.Password)
		if n == 0 {
			break
		}
		item.LetterStr = string(item.Letter)

		a := 0
		if item.Password[item.First-1] == uint8(item.Letter) {
			a = 1
		}

		b := 0
		if item.Password[item.Second-1] == uint8(item.Letter) {
			b = 1
		}

		if a + b == 1 {
			valid++
			fmt.Println("VALID ", item)
		} else {
			fmt.Println("INVALID ", item)
		}
	}
	fmt.Println(valid)
}
