package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	lines := strings.Split(string(buf), "\n")
	transform := make(map[string][]string)
	for _, line := range lines[:len(lines)-2] {
		f := strings.Split(line, " => ")
		from, to := f[0], f[1]
		transform[from] = append(transform[from], to)
	}
	medicine := lines[len(lines)-1]

	count := make(map[string]struct{})
	for from, tos := range transform {
		for i := 0; i <= len(medicine)-len(from); i++ {
			if medicine[i:i+len(from)] == from {
				for _, to := range tos {
					str := medicine[:i] + to + medicine[i+len(from):]
					count[str] = struct{}{}
				}
			}
		}
	}
	fmt.Println("part1:", len(count))

	steps := 0
	for medicine != "e" {
		for from, tos := range transform {
			for _, to := range tos {
				if strings.Contains(medicine, to) {
					medicine = strings.Replace(medicine, to, from, 1)
					steps++
				}
			}
		}
	}
	fmt.Println("part2:", steps)
}
