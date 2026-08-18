from collections import Counter

def main():
	part1 = part2 = 0
	with open("input") as f:
		lines = [line.strip() for line in f]
	for line in lines:
		[prefix, password] = line.split(": ")
		[range, letter] = prefix.split(" ")
		[a, b] = map(int, range.split("-"))
		part1 += a <= Counter(password)[letter] <= b
		part2 += Counter(password[a-1]+password[b-1])[letter] == 1
	print("part1:", part1)
	print("part2:", part2)

if __name__ == "__main__":
	main()
