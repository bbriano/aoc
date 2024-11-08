package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	var ingredient [][]int
	for _, line := range strings.Split(string(input), "\n") {
		line = strings.Split(line, ":")[1]
		var prop []int
		for _, clause := range strings.Split(line, ",") {
			f := strings.Fields(clause)
			n, _ := strconv.Atoi(f[1])
			prop = append(prop, n)
		}
		ingredient = append(ingredient, prop)
	}
	maxscore, maxscore2 := math.MinInt, math.MinInt
	for _, teaspoon := range partition(100, len(ingredient)) {
		prop := make([]int, len(ingredient[0]))
		for i := range prop {
			for j := range ingredient {
				prop[i] += teaspoon[j] * ingredient[j][i]
			}
		}
		maxscore = max(maxscore, score(prop))
		if prop[4] == 500 {
			maxscore2 = max(maxscore2, score(prop))
		}
	}
	fmt.Println("part1:", maxscore)
	fmt.Println("part2:", maxscore2)
}

func partition(n, k int) [][]int {
	if k == 1 {
		return [][]int{{n}}
	}
	var out [][]int
	for i := range n + 1 {
		for _, sub := range partition(n-i, k-1) {
			s := []int{i}
			s = append(s, sub...)
			out = append(out, s)
		}
	}
	return out
}

func score(prop []int) int {
	out := 1
	for _, p := range prop[:4] {
		out *= max(0, p)
	}
	return out
}
