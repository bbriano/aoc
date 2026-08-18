
def main():
	with open("input") as f:
		numbers = [int(line) for line in f]
	print("part1:", twosum(numbers, 2020))
	print("part2:", threesum(numbers, 2020))

def twosum(numbers, target):
	seen = set()
	for x in numbers:
		y = target - x
		if y in seen:
			return x * y
		seen.add(x)

def threesum(numbers, target):
	for i, x in enumerate(numbers):
		yz = twosum(numbers[i+1:], target-x)
		if yz is not None:
			return x * yz

if __name__ == "__main__":
	main()
