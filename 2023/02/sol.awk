BEGIN {
	bag["red"] = 12
	bag["green"] = 13
	bag["blue"] = 14
}
{
	possible = 1
	max["red"] = max["green"] = max["blue"] = 0
	game = substr($0, index($0, ":")+1)
	split(game, round, ";")
	for (i in round) {
		split(round[i], cube, ",")
		for (j in cube) {
			split(cube[j], f)
			if (f[1] > bag[f[2]]) {
				possible = 0
			}
			if (f[1] > max[f[2]]) {
				max[f[2]] = f[1]
			}
		}
	}
	sum += possible * $2
	sum2 += max["red"] * max["green"] * max["blue"]
}
END {
	print "part1:", sum
	print "part2:", sum2
}
