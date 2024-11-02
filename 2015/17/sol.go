package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
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
	var containers []int
	for _, line := range strings.Split(string(buf), "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		containers = append(containers, n)
	}
	count := make(map[int]int)
	for i := 0; i < 1<<len(containers); i++ {
		sum := 0
		for j := range containers {
			sum += containers[j] * (i >> j & 1)
		}
		if sum == 150 {
			count[popcount(i)]++
		}
	}
	part1, part2 := 0, 0
	minnc := math.MaxInt
	for nc := range count {
		part1 += count[nc]
		if nc < minnc {
			minnc = nc
			part2 = count[nc]
		}
	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func popcount(n int) int {
	out := 0
	for n > 0 {
		out += n & 1
		n >>= 1
	}
	return out
}
