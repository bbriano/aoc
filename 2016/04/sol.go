package main

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	part1, part2 := 0, 0
	re := regexp.MustCompile(`([a-z-]+)-([0-9]+)\[([a-z]+)\]`)
	for _, match := range re.FindAllStringSubmatch(string(input), -1) {
		name := match[1]
		sector, _ := strconv.Atoi(match[2])
		sum := match[3]
		if checksum(strings.Replace(name, "-", "", -1)) == sum {
			part1 += sector
		}
		if strings.Contains(decrypt(name, sector), "north") {
			part2 = sector
		}
	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func checksum(s string) string {
	var rune2count [26]int
	for _, r := range s {
		rune2count[r-'a']++
	}
	count2runes := make(map[int][]rune)
	for i, n := range rune2count {
		count2runes[n] = append(count2runes[n], rune(i)+'a')
	}
	keys := slices.Collect(maps.Keys(count2runes))
	slices.Sort(keys)
	slices.Reverse(keys)
	out := ""
	for _, n := range keys {
		out = out + string(count2runes[n])
		if len(out) >= 5 {
			break
		}
	}
	return out[:5]
}

func decrypt(s string, rotate int) string {
	out := []rune(s)
	for i := range out {
		if out[i] == '-' {
			continue
		}
		idx := out[i] - 'a' + rune(rotate)
		out[i] = idx%26 + 'a'
	}
	return string(out)
}
