package main

import (
	"fmt"
	"maps"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	dist := make(map[string]map[string]int)
	for _, line := range strings.Split(string(input), "\n") {
		f := strings.Split(line, " ")
		src, dst := f[0], f[2]
		d, _ := strconv.Atoi(f[4])
		if _, ok := dist[src]; !ok {
			dist[src] = make(map[string]int)
		}
		dist[src][dst] = d
		if _, ok := dist[dst]; !ok {
			dist[dst] = make(map[string]int)
		}
		dist[dst][src] = d
	}
	var walk func(string) (int, int)
	walk = func(u string) (int, int) {
		if len(dist) == 0 {
			return 0, 0
		}
		m, M := math.MaxInt, 0
		for v, dv := range maps.Clone(dist) {
			delete(dist, v)
			d, D := walk(v)
			m = min(m, dv[u]+d)
			M = max(M, dv[u]+D)
			dist[v] = dv
		}
		return m, M
	}
	d, D := walk("")
	fmt.Println("part1:", d)
	fmt.Println("part2:", D)
}
