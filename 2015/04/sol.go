package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	var i int
	for !strings.HasPrefix(hash(i), "00000") {
		i++
	}
	fmt.Println("part1:", i)
	for !strings.HasPrefix(hash(i), "000000") {
		i++
	}
	fmt.Println("part2:", i)
}

func hash(i int) string {
	h := md5.New()
	fmt.Fprintf(h, "bgvyzdsv%d", i)
	return hex.EncodeToString(h.Sum(nil))
}
