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
	happiness := make(map[string]map[string]int)
	for _, line := range strings.Split(string(input), "\n") {
		f := strings.Fields(line)
		u := f[0]
		v, _ := strings.CutSuffix(f[10], ".")
		h, _ := strconv.Atoi(f[3])
		if f[2] == "lose" {
			h = -h
		}
		if _, ok := happiness[u]; !ok {
			happiness[u] = make(map[string]int)
		}
		happiness[u][v] = h
	}
	fmt.Println("part1:", maxHappy(happiness))
	happiness["me"] = nil
	fmt.Println("part2:", maxHappy(happiness))
}

func maxHappy(happiness map[string]map[string]int) int {
	var people []string
	for u := range happiness {
		people = append(people, u)
	}
	maxh := math.MinInt
	for _, p := range permutations(len(people)) {
		h := 0
		for i := range len(p) {
			u := people[p[i]]
			v := people[p[(i+1)%len(p)]]
			h += happiness[u][v]
			h += happiness[v][u]
		}
		maxh = max(maxh, h)
	}
	return maxh
}

func permutations(n int) [][]int {
	if n == 1 {
		return [][]int{{0}}
	}
	var out [][]int
	for _, p := range permutations(n - 1) {
		for i := range n {
			var q []int
			q = append(q, p[:i]...)
			q = append(q, n-1)
			q = append(q, p[i:]...)
			out = append(out, q)
		}
	}
	return out
}
