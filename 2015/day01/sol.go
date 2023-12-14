package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	input, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	floor, pos := 0, []int{}
	for i, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		} else {
			fmt.Fprintln(os.Stderr, "not ( or )\n")
			os.Exit(1)
		}
		if floor == -1 {
			pos = append(pos, i+1)
		}
	}
	fmt.Println("part1:", floor)
	fmt.Println("part2:", pos[0])
}
