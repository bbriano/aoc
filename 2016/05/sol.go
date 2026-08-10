package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	part1 := ""
	part2 := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	for i, count := 0, 0; count < 8; i++ {
		h := hash(fmt.Sprintf("ffykfhsq%d", i))
		if strings.HasPrefix(h, "00000") {
			if len(part1) < 8 {
				part1 += string(h[5])
			}
			idx, val := h[5]-'0', h[6]
			if 0 <= idx && idx <= 7 && part2[idx] == 0 {
				part2[idx] = val
				count++
			}
		}
	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", string(part2))
}

func hash(s string) string {
	h := md5.New()
	fmt.Fprintf(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
