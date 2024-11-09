package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	var capacity []int
	for _, line := range strings.Split(string(input), "\n") {
		n, _ := strconv.Atoi(line)
		capacity = append(capacity, n)
	}
	count := make([]int, len(capacity)+1)
	for i := range 1 << len(capacity) {
		space := 0
		ncontainer := 0
		for j := range capacity {
			if i&(1<<j) != 0 {
				space += capacity[j]
				ncontainer++
			}
		}
		if space == 150 {
			count[ncontainer]++
		}
	}
	fmt.Println("part1:", sum(count))
	least := slices.IndexFunc(count, func(x int) bool { return x >= 1 })
	fmt.Println("part2:", count[least])
}

func sum(s []int) int {
	out := 0
	for _, x := range s {
		out += x
	}
	return out
}
