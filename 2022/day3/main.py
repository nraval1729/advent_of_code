def main():
    rucksacks = read_input("input.txt")
    print(solve_part_one(rucksacks))
    print(solve_part_two(rucksacks))


def solve_part_one(rucksacks):
    priority_sum = 0

    for rucksack in rucksacks:
        first_compartment, second_compartment = rucksack[:len(rucksack) // 2], rucksack[len(rucksack) // 2:]
        common = list(set(first_compartment) & set(second_compartment))[0]

        priority_sum += ord(common) - 38 if common.isupper() else ord(common) - 96

    return priority_sum


def solve_part_two(rucksacks):
    priority_sum = 0

    group_start_index = 0
    while True:
        if group_start_index + 3 > len(rucksacks):
            return priority_sum

        rucksack_one, rucksack_two, rucksack_three = rucksacks[group_start_index: group_start_index + 3]
        common = list(set(rucksack_one) & set(rucksack_two) & set(rucksack_three))[0]

        priority_sum += ord(common) - 38 if common.isupper() else ord(common) - 96

        group_start_index += 3


def read_input(filename):
    with open(filename) as fp:
        return [l.strip() for l in fp.readlines()]


if __name__ == "__main__":
    main()
