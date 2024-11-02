package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input")
	seen := make(map[[2]int]bool)
	seen2 := make(map[[2]int]bool)
	var santa, santa2, robot [2]int
	seen[santa], seen2[santa2] = true, true
	for i, c := range input {
		who := &santa2
		if i%2 == 1 {
			who = &robot
		}
		switch c {
		case '^':
			santa[0]--
			(*who)[0]--
		case 'v':
			santa[0]++
			(*who)[0]++
		case '<':
			santa[1]--
			(*who)[1]--
		case '>':
			santa[1]++
			(*who)[1]++
		}
		seen[santa] = true
		seen2[*who] = true
	}
	fmt.Println("part1:", len(seen))
	fmt.Println("part2:", len(seen2))
}
