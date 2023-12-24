package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	var root interface{}
	if err := json.Unmarshal(buf, &root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("part1:", part1(root))
	fmt.Println("part2:", part2(root))
}

func part1(tree interface{}) int {
	out := 0
	switch tree := tree.(type) {
	case []interface{}:
		for _, sub := range tree {
			out += part1(sub)
		}
	case map[string]interface{}:
		for _, sub := range tree {
			out += part1(sub)
		}
	case float64:
		out = int(tree)
	case string:
	default:
		fmt.Fprintf(os.Stderr, "unhandled type: %T\n", tree)
		os.Exit(1)
	}
	return out
}

func part2(tree interface{}) int {
	out := 0
	switch tree := tree.(type) {
	case []interface{}:
		for _, sub := range tree {
			out += part2(sub)
		}
	case map[string]interface{}:
		for _, sub := range tree {
			if s, ok := sub.(string); ok && s == "red" {
				return 0
			}
			out += part2(sub)
		}
	case float64:
		out = int(tree)
	case string:
	default:
		fmt.Fprintf(os.Stderr, "unhandled type: %T\n", tree)
		os.Exit(1)
	}
	return out
}
