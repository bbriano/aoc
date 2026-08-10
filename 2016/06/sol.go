package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	var message []string
	for _, line := range strings.Split(string(input), "\n") {
		for i, r := range line {
			if i >= len(message) {
				message = append(message, "")
			}
			message[i] += string(r)
		}
	}
	var part1, part2 string
	for _, s := range message {
		min, max := minmax(s)
		part1 += string(max)
		part2 += string(min)
	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func minmax(s string) (rune, rune) {
	count := make(map[rune]int)
	for _, r := range s {
		count[r]++
	}
	var min, max rune
	mincount, maxcount := math.MaxInt, 0
	for r, c := range count {
		if c < mincount {
			min = r
			mincount = c
		}
		if c > maxcount {
			max = r
			maxcount = c
		}
	}
	return min, max
}
