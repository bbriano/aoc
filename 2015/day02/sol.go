package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	paper, ribbon := 0, 0

	file, err := os.Open("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()
	var buf bytes.Buffer
	io.Copy(&buf, file)
	for _, line := range strings.Split(buf.String(), "\n") {
		dim := []int{0, 0, 0}
		for i := range dim {
			d, err := strconv.Atoi(strings.Split(line, "x")[i])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			dim[i] = d
		}

		a := dim[0] * dim[1]
		b := dim[1] * dim[2]
		c := dim[2] * dim[0]
		paper += 2*a + 2*b + 2*c + min(a, b, c)

		x := 2 * (dim[0] + dim[1])
		y := 2 * (dim[1] + dim[2])
		z := 2 * (dim[2] + dim[0])
		ribbon += min(x, y, z) + dim[0]*dim[1]*dim[2]
	}

	fmt.Println("part1:", paper)
	fmt.Println("part2:", ribbon)
}

func min(a ...int) int {
	switch len(a) {
	case 0:
		panic("the list cannot be empty")
	case 1:
		return a[0]
	default:
		x := a[0]
		y := min(a[1:]...)
		if x < y {
			return x
		}
		return y
	}
}
