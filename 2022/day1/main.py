def main():
    calories = read_input("input.txt")

    print(solve_part_one(calories))
    print(solve_part_two(calories))


def solve_part_one(calories):
    return max([sum(elf_calories) for elf_calories in calories])


def solve_part_two(calories):
    sum_calories = [sum(elf_calories) for elf_calories in calories]

    return sum(sorted(sum_calories)[-3:])


def read_input(filename):
    calories = []
    with open(filename) as f:
        elf_calories = []
        for line in f.read().splitlines():
            if not line:
                calories.append(elf_calories)
                elf_calories = []
            else:
                elf_calories.append(int(line))

        calories.append(elf_calories)

    return calories


if __name__ == "__main__":
    main()
