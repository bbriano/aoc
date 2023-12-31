package main

import (
	"fmt"
	"io"
	"os"
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
	var part1, part2 [][]bool
	for _, line := range strings.Split(string(buf), "\n") {
		var row1, row2 []bool
		for _, c := range line {
			switch c {
			case '.':
				row1 = append(row1, false)
				row2 = append(row2, false)
			case '#':
				row1 = append(row1, true)
				row2 = append(row2, true)
			default:
				panic("unhandled character")
			}
		}
		part1 = append(part1, row1)
		part2 = append(part2, row2)
	}
	part2[0][0] = true
	part2[0][len(part2[0])-1] = true
	part2[len(part2)-1][0] = true
	part2[len(part2)-1][len(part2[0])-1] = true
	for i := 0; i < 100; i++ {
		step(part1)
		step(part2)
		part2[0][0] = true
		part2[0][len(part2[0])-1] = true
		part2[len(part2)-1][0] = true
		part2[len(part2)-1][len(part2[0])-1] = true
	}
	fmt.Println("part1:", count(part1))
	fmt.Println("part2:", count(part2))
}

func step(grid [][]bool) {
	var off, on [][2]int
	for i := range grid {
		for j := range grid[i] {
			nc := 0
			neighbours := [][2]int{
				{i - 1, j - 1},
				{i - 1, j},
				{i - 1, j + 1},
				{i, j - 1},
				{i, j + 1},
				{i + 1, j - 1},
				{i + 1, j},
				{i + 1, j + 1},
			}
			for _, idx := range neighbours {
				if idx[0] < 0 || idx[0] >= len(grid) ||
					idx[1] < 0 || idx[1] >= len(grid[i]) {
					continue
				}
				if grid[idx[0]][idx[1]] {
					nc++
				}
			}
			if grid[i][j] && nc != 2 && nc != 3 {
				off = append(off, [2]int{i, j})
			} else if !grid[i][j] && nc == 3 {
				on = append(on, [2]int{i, j})
			}
		}
	}
	for _, idx := range off {
		grid[idx[0]][idx[1]] = false
	}
	for _, idx := range on {
		grid[idx[0]][idx[1]] = true
	}
}

func count(grid [][]bool) int {
	out := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				out++
			}
		}
	}
	return out
}
