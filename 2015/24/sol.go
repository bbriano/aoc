package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var weights []int
	input, _ := os.ReadFile("input")
	for _, line := range strings.Split(string(input), "\n") {
		w, _ := strconv.Atoi(line)
		weights = append(weights, w)
	}
	fmt.Println("parts1:", entanglement(weights, 3))
	fmt.Println("parts2:", entanglement(weights, 4))
}

func entanglement(weights []int, ngroup int) int {
	target := 0
	for _, w := range weights {
		target += w
	}
	target /= ngroup
	minn, minent := math.MaxInt, math.MaxInt
	for partition := range 1 << len(weights) {
		weight, n, ent := 0, 0, 1
		for i, w := range weights {
			if partition&(1<<i) != 0 {
				weight += w
				n++
				ent *= w
			}
		}
		if weight != target {
			continue
		}
		if n < minn {
			minn, minent = n, ent
		} else if n == minn {
			minent = min(minent, ent)
		}
	}
	return minent
}
