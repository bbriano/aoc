package main

import (
	"fmt"
	"strings"
)

func main() {
	pass := "hepxcrrq"
	for !valid(pass) {
		pass = inc(pass)
	}
	fmt.Println("part1:", pass)
	pass = inc(pass)
	for !valid(pass) {
		pass = inc(pass)
	}
	fmt.Println("part2:", pass)
}

func valid(pass string) bool {
	return abc(pass) && noIOL(pass) && aabb(pass)
}

func abc(pass string) bool {
	for i := 1; i < len(pass)-1; i++ {
		if pass[i-1]+1 == pass[i] && pass[i] == pass[i+1]-1 {
			return true
		}
	}
	return false
}

func noIOL(pass string) bool {
	return strings.IndexAny(pass, "iol") == -1
}

func aabb(pass string) bool {
	pairs := 0
	for i := 0; i < len(pass)-1; i++ {
		if pass[i] == pass[i+1] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}

func inc(pass string) string {
	b := []byte(pass)
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 'z' {
			b[i]++
			break
		}
		b[i] = 'a'
	}
	return string(b)
}
