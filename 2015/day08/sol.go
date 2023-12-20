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
	part1, part2 := 0, 0
	for _, line := range strings.Split(string(buf), "\n") {
		dec := 0
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '"':
			case '\\':
				i++
				switch line[i] {
				case '\\':
					dec++
				case '"':
					dec++
				case 'x':
					dec++
					i += 2
				default:
					panic("unhandled escape character")
				}
			default:
				dec++
			}
		}
		enc := 2
		for _, c := range line {
			switch c {
			case '"', '\\':
				enc += 2
			default:
				enc++
			}
		}
		part1 += len(line) - dec
		part2 += enc - len(line)
	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}
