package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	type reindeer struct{ speed, uptime, downtime int }
	var reindeers []reindeer
	input, _ := os.ReadFile("input")
	for _, line := range strings.Split(string(input), "\n") {
		f := strings.Fields(line)
		r := reindeer{
			speed:    atoi(f[3]),
			uptime:   atoi(f[6]),
			downtime: atoi(f[13]),
		}
		reindeers = append(reindeers, r)
	}
	dist := make([]int, len(reindeers))
	points := make([]int, len(reindeers))
	for t := range 2503 {
		for i, r := range reindeers {
			if t%(r.uptime+r.downtime) < r.uptime {
				dist[i] += r.speed
			}
		}
		maxd := max(dist)
		for i := range dist {
			if dist[i] == maxd {
				points[i]++
			}
		}
	}
	fmt.Println("part1:", max(dist))
	fmt.Println("part2:", max(points))
}

func max(s []int) int {
	out := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > out {
			out = s[i]
		}
	}
	return out
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
