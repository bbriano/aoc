
def main():
	with open("input") as f:
		map = [line.strip() for line in f]
	print("part1:", count(map, 3, 1))
	part2 = 1
	for dx, dy in [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]:
		part2 *= count(map, dx, dy)
	print("part2:", part2)

def count(map, dx, dy):
	out = 0
	x, y = 0, 0
	while y < len(map):
		out += at(map, x, y) == "#"
		x += dx
		y += dy
	return out

def at(map, x, y):
	n = len(map[0])
	return map[y][x%n]

if __name__ == "__main__":
	main()
