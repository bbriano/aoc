package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input")
	part1 := make(map[[2]int]int)
	part2 := make(map[[2]int]int)
	var santa, santa2, robot [2]int
	part1[santa]++
	part2[santa2]++
	part2[robot]++
	for i, c := range input {
		switch c {
		case '<':
			santa[1]--
		case '>':
			santa[1]++
		case '^':
			santa[0]--
		case 'v':
			santa[0]++
		}
		part1[santa]++
		switch i % 2 {
		case 0:
			switch c {
			case '<':
				santa2[1]--
			case '>':
				santa2[1]++
			case '^':
				santa2[0]--
			case 'v':
				santa2[0]++
			}
			part2[santa2]++
		case 1:
			switch c {
			case '<':
				robot[1]--
			case '>':
				robot[1]++
			case '^':
				robot[0]--
			case 'v':
				robot[0]++
			}
			part2[robot]++
		}
	}
	fmt.Println("part1:", len(part1))
	fmt.Println("part2:", len(part2))
}
