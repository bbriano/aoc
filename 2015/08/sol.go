package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	var mem, lit, enc int
	for _, s := range strings.Split(string(input), "\n") {
		for i := 1; i < len(s)-1; {
			switch {
			case s[i] != '\\':
				i++
			case s[i+1] == 'x':
				i += 4
			default:
				i += 2
			}
			mem++
		}
		lit += len(s)
		enc += 2 // ""
		for _, c := range s {
			switch c {
			case '"', '\\':
				enc += 2
			default:
				enc += 1
			}
		}
	}
	fmt.Println("part1:", lit-mem)
	fmt.Println("part2:", enc-lit)
}
