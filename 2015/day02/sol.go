package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input")
	re := regexp.MustCompile(`(.+)x(.+)x(.+)`)
	paper, ribbon := 0, 0
	for _, match := range re.FindAllStringSubmatch(string(input), -1) {
		l, _ := strconv.Atoi(match[1])
		w, _ := strconv.Atoi(match[2])
		h, _ := strconv.Atoi(match[3])
		a, b, c := l*w, w*h, h*l
		paper += 2*(a+b+c) + min(a, b, c)
		ribbon += 2*min(l+w, w+h, h+l) + l*w*h
	}
	fmt.Println("part1:", paper)
	fmt.Println("part2:", ribbon)
}
