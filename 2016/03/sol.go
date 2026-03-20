package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	lines := strings.Split(string(input), "\n")
	count1, count2 := 0, 0
	for _, line := range lines {
		f := strings.Fields(line)
		count1 += triangle(f[0], f[1], f[2])
	}
	for i := range len(lines) / 3 {
		top := strings.Fields(lines[3*i+0])
		mid := strings.Fields(lines[3*i+1])
		bot := strings.Fields(lines[3*i+2])
		count2 += triangle(top[0], mid[0], bot[0])
		count2 += triangle(top[1], mid[1], bot[1])
		count2 += triangle(top[2], mid[2], bot[2])
	}
	fmt.Println("part1:", count1)
	fmt.Println("part2:", count2)
}

func triangle(a, b, c string) int {
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	z, _ := strconv.Atoi(c)
	if x+y+z-max(x, y, z) > max(x, y, z) {
		return 1
	}
	return 0
}
