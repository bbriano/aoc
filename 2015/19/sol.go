package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	var repls [][2]string
	replacements, medicine, _ := strings.Cut(string(input), "\n\n")
	for _, line := range strings.Split(replacements, "\n") {
		from, to, _ := strings.Cut(line, " => ")
		repls = append(repls, [2]string{from, to})
	}
	next := make(map[string]bool)
	for i := range medicine {
		for _, rep := range repls {
			from, to := rep[0], rep[1]
			if strings.HasPrefix(medicine[i:], from) {
				next[medicine[:i]+to+medicine[i+len(from):]] = true
			}
		}
	}
	fmt.Println("part1:", len(next))
	steps := 0
	for medicine != "e" {
		for _, rep := range repls {
			from, to := rep[0], rep[1]
			if strings.Contains(medicine, to) {
				medicine = strings.Replace(medicine, to, from, 1)
				steps++
			}
		}
	}
	fmt.Println("part2:", steps)
}
