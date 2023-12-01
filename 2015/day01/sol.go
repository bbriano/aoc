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
	part1(input)
	//part2(input)
}

func part1(input []byte) {
	floor := 0
	for _, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		} else {
			fmt.Fprintln(os.Stderr, "not ( or )\n")
			os.Exit(1)
		}
	}
	fmt.Println(floor)
}

func part2(input []byte) {
	floor := 0
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
			fmt.Println(i + 1)
			return
		}
	}
}
