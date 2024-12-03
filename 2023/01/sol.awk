{
	r = reverse(s = $1)
	sum += digit(s) digit(r)
	sub(/one|two|three|four|five|six|seven|eight|nine/, ":&", s)
	sub(/:one/, 1, s)
	sub(/:two/, 2, s)
	sub(/:three/, 3, s)
	sub(/:four/, 4, s)
	sub(/:five/, 5, s)
	sub(/:six/, 6, s)
	sub(/:seven/, 7, s)
	sub(/:eight/, 8, s)
	sub(/:nine/, 9, s)
	sub(/eno|owt|eerht|ruof|evif|xis|neves|thgie|enin/, ":&", r)
	sub(/:eno/, 1, r)
	sub(/:owt/, 2, r)
	sub(/:eerht/, 3, r)
	sub(/:ruof/, 4, r)
	sub(/:evif/, 5, r)
	sub(/:xis/, 6, r)
	sub(/:neves/, 7, r)
	sub(/:thgie/, 8, r)
	sub(/:enin/, 9, r)
	sum2 += digit(s) digit(r)
}
END {
	print "part1:", sum
	print "part2:", sum2
}

function digit(s) {
	return substr(s, match(s, /[1-9]/), 1)
}

function reverse(s,   r) {
	for (i = 1; i <= length(s); i++) {
		r = substr(s, i, 1) r
	}
	return r
}
