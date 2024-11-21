package main

import "fmt"

func main() {
	row, col := 2978, 3083
	level := row + col - 1
	n := level*(level-1)/2 + level - row
	x := 20151125
	for range n {
		x = x * 252533 % 33554393
	}
	fmt.Println("part1:", x)
}
