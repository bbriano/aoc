package main

import (
	"fmt"
	"io"
	"math"
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

	distance := make(map[string]map[string]int)
	for _, line := range strings.Split(string(buf), "\n") {
		f := strings.Split(line, " = ")
		loc := strings.Split(f[0], " to ")
		dist, err := strconv.Atoi(f[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if _, ok := distance[loc[0]]; !ok {
			distance[loc[0]] = make(map[string]int)
		}
		if _, ok := distance[loc[1]]; !ok {
			distance[loc[1]] = make(map[string]int)
		}
		distance[loc[0]][loc[1]] = dist
		distance[loc[1]][loc[0]] = dist
	}

	var location []string
	for key := range distance {
		location = append(location, key)
	}

	mindist, maxdist := math.MaxInt, 0
	for _, p := range permutations(len(location)) {
		dist := 0
		for i := 0; i < len(p)-1; i++ {
			from := location[p[i]]
			to := location[p[i+1]]
			dist += distance[from][to]
		}
		if dist < mindist {
			mindist = dist
		}
		if dist > maxdist {
			maxdist = dist
		}
	}
	fmt.Println("part1:", mindist)
	fmt.Println("part2:", maxdist)
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
