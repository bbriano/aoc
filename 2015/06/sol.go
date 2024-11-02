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

	lines := strings.Split(string(buf), "\n")
	light1 := [1000][1000]bool{}
	light2 := [1000][1000]int{}
	for _, line := range lines {
		idx := strings.IndexAny(line, "0123456789")
		cmd := line[:idx-1]
		f := strings.Split(line[idx:], " ")
		p1 := strings.Split(f[0], ",")
		p2 := strings.Split(f[2], ",")
		p1x, p1y := atoi(p1[0]), atoi(p1[1])
		p2x, p2y := atoi(p2[0]), atoi(p2[1])
		for y := p1y; y <= p2y; y++ {
			for x := p1x; x <= p2x; x++ {
				switch cmd {
				case "turn on":
					light1[y][x] = true
					light2[y][x]++
				case "toggle":
					light1[y][x] = !light1[y][x]
					light2[y][x] += 2
				case "turn off":
					light1[y][x] = false
					light2[y][x]--
					if light2[y][x] < 0 {
						light2[y][x] = 0
					}
				default:
					panic("unhandled cmd")
				}
			}
		}
	}

	part1, part2 := 0, 0
	for y := range light1 {
		for x := range light1[y] {
			if light1[y][x] {
				part1++
			}
			part2 += light2[y][x]
		}
	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
