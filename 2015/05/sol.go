package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	var nice, nice2 int
	for _, s := range strings.Split(string(input), "\n") {
		if vowel(s) && pair(s) && !ugly(s) {
			nice++
		}
		if doublepair(s) && between(s) {
			nice2++
		}
	}
	fmt.Println("part1:", nice)
	fmt.Println("part2:", nice2)
}

func vowel(s string) bool {
	for range 3 {
		i := strings.IndexAny(s, "aeiou")
		if i == -1 {
			return false
		}
		s = s[i+1:]
	}
	return true
}

func pair(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func ugly(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		switch s[i : i+2] {
		case "ab", "cd", "pq", "xy":
			return true
		}
	}
	return false
}

func doublepair(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		for j := i + 2; j < len(s)-1; j++ {
			if s[i:i+2] == s[j:j+2] {
				return true
			}
		}
	}
	return false
}

func between(s string) bool {
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] == s[i+1] {
			return true
		}
	}
	return false
}
