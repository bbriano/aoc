package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	var ingredients [][]int
	for _, line := range strings.Split(string(buf), "\n") {
		f := strings.Split(line, ": ")
		var ing []int
		for _, prop := range strings.Split(f[1], ", ") {
			f := strings.Split(prop, " ")
			val, err := strconv.Atoi(f[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			ing = append(ing, val)
		}
		ingredients = append(ingredients, ing)
	}
	part1, part2 := 0, 0
	for _, teaspoon := range share(100, 5) {
		props := make([]int, len(ingredients[0]))
		for i := range ingredients {
			for j := 0; j < len(ingredients[i]); j++ {
				props[j] += ingredients[i][j] * teaspoon[i]
			}
		}
		score := 1
		for i := 0; i < len(props)-1; i++ {
			if props[i] < 0 {
				score = 0
			}
			score *= props[i]
		}
		if score > part1 {
			part1 = score
		}
		if props[4] == 500 && score > part2 {
			part2 = score
		}
	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func share(n, p int) [][]int {
	var out [][]int
	if p == 1 {
		return [][]int{{n}}
	}
	for i := 0; i <= n; i++ {
		for _, sub := range share(n-i, p-1) {
			el := []int{i}
			for _, s := range sub {
				el = append(el, s)
			}
			out = append(out, el)
		}
	}
	return out
}
