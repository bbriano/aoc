package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"1113122113"}
	for i := 0; i < 50; i++ {
		var next []string
		j, count := 1, 1
		for j < len(s[len(s)-1]) {
			if s[len(s)-1][j] == s[len(s)-1][j-1] {
				count++
			} else {
				next = append(next, fmt.Sprintf("%d%c", count, s[len(s)-1][j-1]))
				count = 1
			}
			j++
		}
		next = append(next, fmt.Sprintf("%d%c", count, s[len(s)-1][j-1]))
		s = append(s, strings.Join(next, ""))
	}
	fmt.Println("part1:", len(s[40]))
	fmt.Println("part2:", len(s[50]))
}
