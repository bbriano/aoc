package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	part1, part2 := 0, 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		part1 += 10*leftDigit(line) + rightDigit(line)
		part2 += 10*leftDigit2(line) + rightDigit2(line)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func leftDigit(s string) int {
	for _, c := range s {
		if '0' <= c && c <= '9' {
			return int(c) - '0'
		}
	}
	return 0
}

func rightDigit(s string) int {
	for i := range s {
		c := s[len(s)-1-i]
		if '0' <= c && c <= '9' {
			return int(c) - '0'
		}
	}
	return 0
}

func leftDigit2(s string) int {
	for i, c := range s {
		if '0' <= c && c <= '9' {
			return int(c) - '0'
		}
		for j, digit := range digits {
			if strings.HasPrefix(s[i:], digit) {
				return j + 1
			}
		}
	}
	return 0
}

func rightDigit2(s string) int {
	for i := range s {
		c := s[len(s)-1-i]
		if '0' <= c && c <= '9' {
			return int(c) - '0'
		}
		for j, digit := range digits {
			if strings.HasSuffix(s[:len(s)-i], digit) {
				return j + 1
			}
		}
	}
	return 0
}

var digits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}
