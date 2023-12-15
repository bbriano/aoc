package main

import (
	"fmt"
	"io"
	"os"
)

type point struct{ x, y int }

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	input, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	part1 := map[string]int{"0,0": 1}
	part2 := map[string]int{"0,0": 1}
	santa := point{0, 0}
	santa2 := point{0, 0}
	robot := point{0, 0}
	for i, c := range input {
		switch c {
		case '^':
			santa.y--
		case 'v':
			santa.y++
		case '<':
			santa.x--
		case '>':
			santa.x++
		}
		part1[fmt.Sprintf("%d,%d", santa.x, santa.y)]++

		if i%2 == 0 {
			switch c {
			case '^':
				santa2.y--
			case 'v':
				santa2.y++
			case '<':
				santa2.x--
			case '>':
				santa2.x++
			}
			part2[fmt.Sprintf("%d,%d", santa2.x, santa2.y)]++
		} else {
			switch c {
			case '^':
				robot.y--
			case 'v':
				robot.y++
			case '<':
				robot.x--
			case '>':
				robot.x++
			}
			part2[fmt.Sprintf("%d,%d", robot.x, robot.y)]++
		}
	}
	fmt.Println("part1:", len(part1))
	fmt.Println("part2:", len(part2))
}
