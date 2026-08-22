from collections import defaultdict

def main():
	containedby = defaultdict(dict)
	contains = defaultdict(dict)
	with open("input") as f:
		for line in f:
			if "no other bags" in line:
				continue
			parent, children = line.split(" bags contain ")
			for entry in children.removesuffix(".\n").split(", "):
				count, child = entry.split(" ", 1)
				count = int(count)
				child = child.removesuffix("s").removesuffix(" bag")
				containedby[child][parent] = count
				contains[parent][child] = count

	seen = set()
	tosee = {"shiny gold"}
	while len(tosee) != 0:
		node = tosee.pop()
		for parent in containedby[node]:
			if parent not in seen:
				tosee.add(parent)
		seen.add(node)
	print("part1:", len(seen)-1)

	def count(node):
		out = 1
		for child, n in contains[node].items():
			out += n * count(child)
		return out
	print("part2:", count("shiny gold")-1)

if __name__ == "__main__":
	main()
