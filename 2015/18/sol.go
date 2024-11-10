package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	var grid [][]bool
	for _, line := range strings.Split(string(input), "\n") {
		var row []bool
		for _, c := range line {
			row = append(row, c == '#')
		}
		grid = append(grid, row)
	}
	m, n := len(grid), len(grid[0])
	tmp := make([][]bool, m)
	grid2 := make([][]bool, m)
	for i := range m {
		tmp[i] = make([]bool, n)
		grid2[i] = make([]bool, n)
		copy(grid2[i], grid[i])
	}
	grid2[0][0], grid2[0][n-1] = true, true
	grid2[m-1][0], grid2[m-1][n-1] = true, true
	for range 100 {
		conway(tmp, grid)
		grid, tmp = tmp, grid
		conway(tmp, grid2)
		grid2, tmp = tmp, grid2
		grid2[0][0], grid2[0][n-1] = true, true
		grid2[m-1][0], grid2[m-1][n-1] = true, true
	}
	fmt.Println("part1:", count(grid))
	fmt.Println("part2:", count(grid2))
}

func conway(dst, src [][]bool) {
	m, n := len(src), len(src[0])
	for i := range m {
		for j := range n {
			neighbours := 0
			for p := max(0, i-1); p <= min(m-1, i+1); p++ {
				for q := max(0, j-1); q <= min(n-1, j+1); q++ {
					if p == i && q == j {
						continue
					}
					if src[p][q] {
						neighbours++
					}
				}
			}
			if src[i][j] {
				switch neighbours {
				case 2, 3:
					dst[i][j] = true
				default:
					dst[i][j] = false
				}
			} else {
				switch neighbours {
				case 3:
					dst[i][j] = true
				default:
					dst[i][j] = false
				}
			}
		}
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
