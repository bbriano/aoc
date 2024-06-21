package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	var i int
	for i = 1; hash("bgvyzdsv", i)[:5] != "00000"; i++ {
	}
	fmt.Println("part1:", i)
	for i = 1; hash("bgvyzdsv", i)[:6] != "000000"; i++ {
	}
	fmt.Println("part2:", i)
}

func hash(s string, i int) string {
	data := []byte(s + strconv.Itoa(i))
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
