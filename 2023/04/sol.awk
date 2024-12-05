{
	sub(/.*:/, "")
	split($0, card, "|")
	split(card[1], winning)
	split(card[2], numbers)
	count = 0
	for (i in numbers) {
		for (j in winning) {
			if (numbers[i] == winning[j]) {
				count++
				break
			}
		}
	}
	points += int(2**(count-1))
	copy[NR]++
	for (i = 1; i <= count; i++) {
		copy[NR+i] += copy[NR]
	}
	scratchcards += copy[NR]
}
END {
	print "part1:", points
	print "part2:", scratchcards
}
