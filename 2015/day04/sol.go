package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	i := 1
	for hash("bgvyzdsv", i)[:5] != "00000" {
		i++
	}
	fmt.Println("part1:", i)

	i = 1
	for hash("bgvyzdsv", i)[:6] != "000000" {
		i++
	}
	fmt.Println("part2:", i)
}

func hash(s string, i int) string {
	h := md5.New()
	fmt.Fprintf(h, "%s%d", s, i)
	sum := h.Sum([]byte{})
	return hex.EncodeToString(sum)
}
