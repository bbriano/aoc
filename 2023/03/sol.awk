{ line[NR] = $0 }
END {
	for (i = 1; i <= NR; i++) {
		j = 1
		while (match(substr(line[i], j), "[0-9]+")) {
			start = j + RSTART - 1
			n = substr(line[i], start, RLENGTH)
			surround = ""
			for (row = i-1; row <= i+1; row++) {
				for (col = start-1; col <= start+RLENGTH; col++) {
					c = substr(line[row], col, 1)
					surround = surround c
					if (c == "*") {
						star[row, col] = star[row, col] " " n
					}
				}
			}
			if (surround ~ /[^0-9.]/) {
				sum += n
			}
			j = start + RLENGTH
		}
	}
	for (i in star) {
		split(star[i], f)
		if (length(f) == 2) {
			sum2 += f[1] * f[2]
		}
	}
	print "part1:", sum
	print "part2:", sum2
}
