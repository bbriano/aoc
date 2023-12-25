package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
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

	happiness := make(map[string]map[string]int)
	re := regexp.MustCompile("([A-Za-z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([A-Za-z]+)\\.")
	for _, f := range re.FindAllStringSubmatch(string(buf), -1) {
		n, err := strconv.Atoi(f[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if f[2] == "lose" {
			n = -n
		}
		if happiness[f[1]] == nil {
			happiness[f[1]] = make(map[string]int)
		}
		happiness[f[1]][f[4]] = n
	}

	fmt.Println("part1:", maxHappy(happiness))
	happiness["$"] = nil
	fmt.Println("part2:", maxHappy(happiness))
}

func maxHappy(happiness map[string]map[string]int) int {
	var people []string
	for person := range happiness {
		people = append(people, person)
	}
	maxhappy := math.MinInt
	for _, p := range permutations(len(people)) {
		happy := 0
		for i := range p {
			first := people[p[i]]
			second := people[p[(i+1)%len(p)]]
			happy += happiness[first][second]
			happy += happiness[second][first]
		}
		if happy > maxhappy {
			maxhappy = happy
		}
	}
	return maxhappy
}

func permutations(n int) [][]int {
	if n == 1 {
		return [][]int{{0}}
	}
	var out [][]int
	for _, p := range permutations(n - 1) {
		for i := 0; i < n; i++ {
			var q []int
			q = append(q, p[:i]...)
			q = append(q, n-1)
			q = append(q, p[i:]...)
			out = append(out, q)
		}
	}
	return out
}
