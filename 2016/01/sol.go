package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	North uint = iota
	West
	South
	East
)

type location [2]int

func main() {
	input, _ := os.ReadFile("input")
	x, y, dir := 0, 0, North
	x2, y2 := math.MaxInt, math.MaxInt
	visited := map[location]bool{{0, 0}: true}
	for _, inst := range strings.Split(string(input), ", ") {
		switch inst[0] {
		case 'L':
			dir = (dir + 1) % 4
		case 'R':
			dir = (dir - 1) % 4
		}
		n, _ := strconv.Atoi(inst[1:])
		for range n {
			switch dir {
			case North:
				y++
			case West:
				x--
			case South:
				y--
			case East:
				x++
			}
			if visited[location{x, y}] && x2 == math.MaxInt {
				x2, y2 = x, y
			}
			visited[location{x, y}] = true
		}
	}
	part1 := abs(x) + abs(y)
	part2 := abs(x2) + abs(y2)
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
