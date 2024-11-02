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
	buf, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	input := string(buf)
	part1, part2 := 0, 0
	for _, name := range strings.Split(input, "\n") {
		if vowel3(name) && repeat(name) && nobadseq(name) {
			part1++
		}
		if repeatpair(name) && repeatgap(name) {
			part2++
		}
	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func vowel3(name string) bool {
	vc := 0
	for _, c := range name {
		if strings.ContainsAny(string(c), "aeiou") {
			vc++
		}
	}
	return vc >= 3
}
func repeat(name string) bool {
	for i := 0; i < len(name)-1; i++ {
		if name[i] == name[i+1] {
			return true
		}
	}
	return false
}
func nobadseq(name string) bool {
	for i := 0; i < len(name)-1; i++ {
		switch name[i : i+2] {
		case "ab", "cd", "pq", "xy":
			return false
		}
	}
	return true
}

func repeatpair(name string) bool {
	for i := 0; i < len(name)-2; i++ {
		for j := i + 2; j < len(name)-1; j++ {
			if name[i] == name[j] && name[i+1] == name[j+1] {
				return true
			}
		}
	}
	return false
}
func repeatgap(name string) bool {
	for i := 0; i < len(name)-2; i++ {
		if name[i] == name[i+2] {
			return true
		}
	}
	return false
}
