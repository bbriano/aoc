
def main():
	groups = []
	with open("input") as f:
		for par in f.read().split("\n\n"):
			people = []
			for line in par.split("\n"):
				people.append(line)
			groups.append(people)
	part1, part2 = 0, 0
	for people in groups:
		u = set()
		i = set("abcdefghijklmnopqrstuvwxyz")
		for questions in people:
			u = u.union(set(questions))
			i = i.intersection(set(questions))
		part1 += len(u)
		part2 += len(i)
	print("part1:", part1)
	print("part2:", part2)

if __name__ == "__main__":
	main()
