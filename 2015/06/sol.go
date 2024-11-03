package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input")
	const N = 1000
	var grid [N][N]bool
	var grid2 [N][N]int
	re := regexp.MustCompile(`(.*) (\d+),(\d+) through (\d+),(\d+)`)
	for _, match := range re.FindAllStringSubmatch(string(input), -1) {
		x0, y0 := atoi(match[2]), atoi(match[3])
		x1, y1 := atoi(match[4]), atoi(match[5])
		for y := y0; y <= y1; y++ {
			for x := x0; x <= x1; x++ {
				switch match[1] {
				case "turn off":
					grid[y][x] = false
					grid2[y][x] = max(0, grid2[y][x]-1)
				case "turn on":
					grid[y][x] = true
					grid2[y][x]++
				case "toggle":
					grid[y][x] = !grid[y][x]
					grid2[y][x] += 2
				}
			}
		}
	}
	var c, c2 int
	for y := range N {
		for x := range N {
			if grid[y][x] {
				c++
			}
			c2 += grid2[y][x]
		}
	}
	fmt.Println("part1:", c)
	fmt.Println("part2:", c2)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
