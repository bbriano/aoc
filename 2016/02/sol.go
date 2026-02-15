package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	label string
	next  map[rune]*node
}

var map1 = [][]*node{
	{{"1", nil}, {"2", nil}, {"3", nil}},
	{{"4", nil}, {"5", nil}, {"6", nil}},
	{{"7", nil}, {"8", nil}, {"9", nil}},
}

var map2 = [][]*node{
	{{".", nil}, {".", nil}, {"1", nil}, {".", nil}, {".", nil}},
	{{".", nil}, {"2", nil}, {"3", nil}, {"4", nil}, {".", nil}},
	{{"5", nil}, {"6", nil}, {"7", nil}, {"8", nil}, {"9", nil}},
	{{".", nil}, {"A", nil}, {"B", nil}, {"C", nil}, {".", nil}},
	{{".", nil}, {".", nil}, {"D", nil}, {".", nil}, {".", nil}},
}

func init() {
	for i := range map1 {
		for j := range map1[i] {
			map1[i][j].next = map[rune]*node{}
			map1[i][j].next['U'] = map1[max(i-1, 0)][j]
			map1[i][j].next['D'] = map1[min(i+1, 2)][j]
			map1[i][j].next['L'] = map1[i][max(j-1, 0)]
			map1[i][j].next['R'] = map1[i][min(j+1, 2)]
		}
	}
	for i := range map2 {
		for j := range map2[i] {
			map2[i][j].next = map[rune]*node{}
			map2[i][j].next['U'] = map2[max(i-1, abs(j-2))][j]
			map2[i][j].next['D'] = map2[min(i+1, 4-abs(j-2))][j]
			map2[i][j].next['L'] = map2[i][max(j-1, abs(i-2))]
			map2[i][j].next['R'] = map2[i][min(j+1, 4-abs(i-2))]
		}
	}
}

func main() {
	input, _ := os.ReadFile("input")
	code1, n1 := "", map1[1][1]
	code2, n2 := "", map2[2][0]
	for _, line := range strings.Split(string(input), "\n") {
		for _, c := range line {
			n1 = n1.next[c]
			n2 = n2.next[c]
		}
		code1 += n1.label
		code2 += n2.label
	}
	fmt.Println("part1:", code1)
	fmt.Println("part2:", code2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
