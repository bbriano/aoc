package main

import (
	"strings"
	"fmt"
	"os"
	"strconv"
	"io"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	buf, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	target := map[string]int{
		"children": 3,
		"cats": 7,
		"samoyeds": 2,
		"pomeranians": 3,
		"akitas": 0,
		"vizslas": 0,
		"goldfish": 5,
		"trees": 3,
		"cars": 2,
		"perfumes": 1,
	}
	for _, line := range strings.Split(string(buf), "\n") {
		sue, props, ok := strings.Cut(line, ": ")
		if !ok {
			panic("bad cut")
		}
		part1, part2 := true, true
		for _, prop := range strings.Split(props, ", ") {
			key, val, ok := strings.Cut(prop, ": ")
			if !ok {
				panic("bad cut")
			}
			n, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			if n != target[key] {
				part1 = false
			}
			switch key {
			case "cats", "trees":
				if n <= target[key] {
					part2 = false
				}
			case "pomeranians", "goldfish":
				if n >= target[key] {
					part2 = false
				}
			default:
				if n != target[key] {
					part2 = false
				}
			}
		}
		if part1 {
			fmt.Println("part1:", sue)
		}
		if part2 {
			fmt.Println("part2:", sue)
		}
	}
}
