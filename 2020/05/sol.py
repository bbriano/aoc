
def main():
	with open("input") as f:
		codes = [line.strip() for line in f]
	m = len(codes)
	n = len(codes[0])
	maxid = 0
	available = [True] * m*n
	for code in codes:
		row, col = coord(code)
		id = row*8 + col
		maxid = max(maxid, id)
		available[id] = False
	print("part1:", maxid)
	print("part2:", available.index(False) + available[available.index(False):].index(True))

def coord(code):
	row, col = 0, 0
	for i in range(7):
		row += (code[6-i] == "B") * 1<<i
	for i in range(3):
		col += (code[9-i] == "R") * 1<<i
	return row, col

if __name__ == "__main__":
	main()
