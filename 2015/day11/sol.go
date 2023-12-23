package main

import "fmt"

func main() {
	pass := "hepxcrrz"
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
	for i := 0; i < len(pass)-2; i++ {
		if pass[i]+1 == pass[i+1] && pass[i+1]+1 == pass[i+2] {
			return true
		}
	}
	return false
}

func noIOL(pass string) bool {
	for _, c := range pass {
		switch c {
		case 'i', 'o', 'l':
			return false
		}
	}
	return true
}

func aabb(pass string) bool {
	i := 0
	first := byte('a')
	for i < len(pass)-1 {
		if pass[i] == pass[i+1] {
			first = pass[i]
			break
		}
		i++
	}
	i += 2
	for i < len(pass)-1 {
		if pass[i] == pass[i+1] && pass[i] != first {
			return true
		}
		i++
	}
	return false
}

func inc(pass string) string {
	buf := []byte(pass)
	for i := len(pass) - 1; i >= 0; i-- {
		if pass[i] == 'z' {
			buf[i] = 'a'
		} else {
			buf[i]++
			break
		}
	}
	return string(buf)
}
