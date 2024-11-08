package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input")
	d := json.NewDecoder(f)
	d.UseNumber()
	var data any
	d.Decode(&data)
	fmt.Println("part1:", part1(data))
	fmt.Println("part2:", part2(data))
}

func part1(node any) int {
	out := 0
	switch node := node.(type) {
	case string:
	case json.Number:
		n, _ := node.Int64()
		out += int(n)
	case []any:
		for _, v := range node {
			out += part1(v)
		}
	case map[string]any:
		for _, v := range node {
			out += part1(v)
		}
	}
	return out
}

func part2(node any) int {
	out := 0
	switch node := node.(type) {
	case string:
	case json.Number:
		n, _ := node.Int64()
		out += int(n)
	case []any:
		for _, v := range node {
			out += part2(v)
		}
	case map[string]any:
		for _, v := range node {
			if v == "red" {
				return 0
			}
		}
		for _, v := range node {
			out += part2(v)
		}
	}
	return out
}
