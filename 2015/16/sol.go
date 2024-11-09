package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var gift = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func main() {
	input, _ := os.ReadFile("input")
	for _, line := range strings.Split(string(input), "\n") {
		part1, part2 := true, true
		sue, line, _ := strings.Cut(line, ": ")
		for _, props := range strings.Split(line, ", ") {
			key, v, _ := strings.Cut(props, ": ")
			value, _ := strconv.Atoi(v)
			if value != gift[key] {
				part1 = false
			}
			switch key {
			case "cats", "trees":
				if value <= gift[key] {
					part2 = false
				}
			case "pomeranians", "goldfish":
				if value >= gift[key] {
					part2 = false
				}
			default:
				if value != gift[key] {
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
