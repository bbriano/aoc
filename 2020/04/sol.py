import re

def main():
	passports = parse()
	count1 = 0
	count2 = 0
	for p in passports:
		count1 += valid1(p)
		count2 += valid2(p)
	print("part1:", count1)
	print("part2:", count2)

def parse():
	with open("input") as f:
		paragraphs = f.read().split("\n\n")
	passports = []
	for par in paragraphs:
		fields = {}
		for field in re.split(r"[ \n]", par):
			key, val = field.split(":")
			fields[key] = val
		passports.append(fields)
	return passports

def valid1(passport):
	for key in ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]:
		if key not in passport:
			return False
	return True

def valid2(passport):
	if not valid1(passport):
		return False
	if not (1920 <= int(passport["byr"]) <= 2002):
		return False
	if not (2010 <= int(passport["iyr"]) <= 2020):
		return False
	if not (2020 <= int(passport["eyr"]) <= 2030):
		return False
	if passport["hgt"].endswith("cm"):
		if not (150 <= int(passport["hgt"].removesuffix("cm")) <= 193):
			return False
	elif passport["hgt"].endswith("in"):
		if not (59 <= int(passport["hgt"].removesuffix("in")) <= 76):
			return False
	else:
		return False
	if not re.fullmatch(r"#[0-9a-z]{6}", passport["hcl"]):
		return False
	if passport["ecl"] not in {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}:
		return False
	if not re.fullmatch(r"[0-9]{9}", passport["pid"]):
		return False
	return True

if __name__ == "__main__":
	main()
