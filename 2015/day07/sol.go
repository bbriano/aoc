package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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

	instr := make(map[string][]string)
	for _, line := range strings.Split(string(buf), "\n") {
		fromto := strings.Split(line, " -> ")
		expr, key := fromto[0], fromto[1]
		instr[key] = strings.Split(expr, " ")
	}
	cache := make(map[string]int)

	var eval func(key string) int
	eval = func(key string) int {
		if val, ok := cache[key]; ok {
			return val
		}
		if n, err := strconv.Atoi(key); err == nil {
			cache[key] = n
			return n
		}
		out := 0
		cmd := instr[key]
		switch len(cmd) {
		case 1:
			out = eval(cmd[0])
		case 2:
			out = ^eval(cmd[1])
		case 3:
			switch cmd[1] {
			case "AND":
				out = eval(cmd[0]) & eval(cmd[2])
			case "OR":
				out = eval(cmd[0]) | eval(cmd[2])
			case "LSHIFT":
				out = eval(cmd[0]) << eval(cmd[2])
			case "RSHIFT":
				out = eval(cmd[0]) >> eval(cmd[2])
			default:
				panic("unhandled operator")
			}
		default:
			panic("unhandled command")
		}
		cache[key] = out
		return out
	}

	part1 := eval("a")
	cache = make(map[string]int)
	cache["b"] = part1
	part2 := eval("a")
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}
