package main

import "fmt"

func main() {
	const N = 33100000 / 10
	part1 := make([]int, N)
	part2 := make([]int, N)
	gifted := make([]int, N)
	for elf := 1; elf < N; elf++ {
		for house := elf; house < N; house += elf {
			part1[house] += 10 * elf
			if gifted[elf] < 50 {
				part2[house] += 11 * elf
				gifted[elf]++
			}
		}
	}
	for house := range N {
		if part1[house] >= 33100000 {
			fmt.Println("part1:", house)
			break
		}
	}
	for house := range N {
		if part2[house] >= 33100000 {
			fmt.Println("part2:", house)
			break
		}
	}
}
