package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

type reindeer struct {
	name     string
	speed    int
	uptime   int
	downtime int
}

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

	var reindeers []reindeer
	re := regexp.MustCompile("([A-Za-z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.")
	for _, match := range re.FindAllStringSubmatch(string(buf), -1) {
		r := reindeer{
			name:     match[1],
			speed:    atoi(match[2]),
			uptime:   atoi(match[3]),
			downtime: atoi(match[4]),
		}
		reindeers = append(reindeers, r)
	}

	dist := make([]int, len(reindeers))
	points := make([]int, len(reindeers))
	state := make([]int, len(reindeers))
	for i := 0; i < 2503; i++ {
		for j, r := range reindeers {
			if state[j] >= 0 {
				dist[j] += r.speed
			}
			state[j]++
			if state[j] >= r.uptime {
				state[j] = -r.downtime
			}
		}
		maxdist := 0
		for i := range dist {
			if dist[i] > maxdist {
				maxdist = dist[i]
			}
		}
		for i := range dist {
			if dist[i] == maxdist {
				points[i]++
			}
		}
	}

	fmt.Println("part1:", max(dist))
	fmt.Println("part2:", max(points))
}

func max(a []int) int {
	out := math.MinInt
	for _, x := range a {
		if x > out {
			out = x
		}
	}
	return out
}

func atoi(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return out
}
