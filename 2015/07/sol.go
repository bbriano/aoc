package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	wire := make(map[string]func() int)
	cache := make(map[string]int)
	value := func(s string) int {
		if n, err := strconv.Atoi(s); err == nil {
			return n
		}
		if n, ok := cache[s]; ok {
			return n
		}
		cache[s] = wire[s]()
		return cache[s]
	}
	for _, line := range strings.Split(string(input), "\n") {
		f := strings.Fields(line)
		switch len(f) {
		case 3: // x -> y
			wire[f[2]] = func() int { return value(f[0]) }
		case 4: // NOT x -> y
			wire[f[3]] = func() int { return ^value(f[1]) }
		case 5: // x op y -> z
			switch f[1] {
			case "AND":
				wire[f[4]] = func() int { return value(f[0]) & value(f[2]) }
			case "OR":
				wire[f[4]] = func() int { return value(f[0]) | value(f[2]) }
			case "LSHIFT":
				wire[f[4]] = func() int { return value(f[0]) << value(f[2]) }
			case "RSHIFT":
				wire[f[4]] = func() int { return value(f[0]) >> value(f[2]) }
			}
		}
	}
	a := wire["a"]()
	fmt.Println("part1:", a)
	wire["b"] = func() int { return a }
	cache = make(map[string]int)
	fmt.Println("part2:", wire["a"]())
}
