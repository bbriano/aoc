package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type machine struct {
	a, b, pc int
	inst     []inst
}

type inst struct {
	op     code
	r      register
	offset int
}

type code int

const (
	hlf code = iota
	tpl
	inc
	jmp
	jie
	jio
)

type register int

const (
	a register = iota
	b
)

func (m *machine) exec() {
	for m.pc < len(m.inst) {
		i := m.inst[m.pc]
		switch i.op {
		case hlf:
			m.hlf(i.r)
		case tpl:
			m.tpl(i.r)
		case inc:
			m.inc(i.r)
		case jmp:
			m.jmp(i.offset)
		case jie:
			m.jie(i.r, i.offset)
		case jio:
			m.jio(i.r, i.offset)
		}
		m.pc++
	}
}

func (m *machine) hlf(r register) {
	switch r {
	case a:
		m.a /= 2
	case b:
		m.b /= 2
	}
}

func (m *machine) tpl(r register) {
	switch r {
	case a:
		m.a *= 3
	case b:
		m.b *= 3
	}
}

func (m *machine) inc(r register) {
	switch r {
	case a:
		m.a++
	case b:
		m.b++
	}
}

func (m *machine) jmp(offset int) {
	m.pc += offset
	m.pc--
}

func (m *machine) jie(r register, offset int) {
	var x int
	switch r {
	case a:
		x = m.a
	case b:
		x = m.b
	}
	if x%2 == 0 {
		m.pc += offset
		m.pc--
	}
}

func (m *machine) jio(r register, offset int) {
	var x int
	switch r {
	case a:
		x = m.a
	case b:
		x = m.b
	}
	if x == 1 {
		m.pc += offset
		m.pc--
	}
}

func main() {
	var m machine
	input, _ := os.ReadFile("input")
	for _, line := range strings.Split(string(input), "\n") {
		f := strings.Fields(line)
		var i inst
		switch f[0] {
		case "hlf":
			i.op = hlf
			i.r = parsereg(f[1])
		case "tpl":
			i.op = tpl
			i.r = parsereg(f[1])
		case "inc":
			i.op = inc
			i.r = parsereg(f[1])
		case "jmp":
			i.op = jmp
			i.offset = atoi(f[1])
		case "jie":
			i.op = jie
			i.r = parsereg(f[1])
			i.offset = atoi(f[2])
		case "jio":
			i.op = jio
			i.r = parsereg(f[1])
			i.offset = atoi(f[2])
		}
		m.inst = append(m.inst, i)
	}
	m.exec()
	fmt.Println("part1:", m.b)
	m.a, m.b, m.pc = 1, 0, 0
	m.exec()
	fmt.Println("part2:", m.b)
}

func parsereg(s string) register {
	switch s[0] {
	case 'a':
		return a
	case 'b':
		return b
	}
	return register(-1)
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
