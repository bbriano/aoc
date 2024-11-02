package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	var paper, ribbon int
	for {
		var x, y, z int
		if _, err := fmt.Fscanf(f, "%dx%dx%d\n", &x, &y, &z); err != nil {
			break
		}
		paper += 2*(x*y+y*z+z*x) + x*y*z/max(x, y, z)
		ribbon += 2*(x+y+z-max(x, y, z)) + x*y*z
	}
	fmt.Println("part1:", paper)
	fmt.Println("part2:", ribbon)
}
