package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input")
	floor, pos := 0, 0
	for i, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
		if pos == 0 && floor == -1 {
			pos = i + 1
		}
	}
	fmt.Println("part1:", floor)
	fmt.Println("part2:", pos)
}
