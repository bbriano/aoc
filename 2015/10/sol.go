package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "1113122113"
	for range 40 {
		s = looksay(s)
	}
	fmt.Println("part1:", len(s))
	for range 10 {
		s = looksay(s)
	}
	fmt.Println("part2:", len(s))
}

func looksay(s string) string {
	var sb strings.Builder
	for len(s) > 0 {
		n := 1
		for n < len(s) && s[n] == s[n-1] {
			n++
		}
		sb.WriteString(strconv.Itoa(n))
		sb.WriteByte(s[0])
		s = s[n:]
	}
	return sb.String()
}
