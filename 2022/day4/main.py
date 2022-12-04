def main():
    sections = read_input("input.txt")
    print(solve_part_one(sections))
    print(solve_part_two(sections))


def solve_part_one(sections):
    num_overlaps = 0

    for section in sections:
        one_start, one_end, two_start, two_end = section

        overlapping_conditions = []

        # First section is only one element long and is same as second's start
        overlapping_conditions.append(one_start == two_start and one_end == two_start)

        # First section is only one element long and is same as second's end
        overlapping_conditions.append(one_start == two_end and one_end == two_end)

        # Second section is only one element long and is same as first's start
        overlapping_conditions.append(two_start == one_start and two_end == one_start)

        # Second section is only one element long and is same as first's end
        overlapping_conditions.append(two_start == one_end and two_end == one_end)

        # First section is fully contained within second
        overlapping_conditions.append(one_start >= two_start and one_end <= two_end)

        # Second section is fully contained within first
        overlapping_conditions.append(two_start >= one_start and two_end <= one_end)

        num_overlaps = num_overlaps + 1 if any(overlapping_conditions) else num_overlaps

    return num_overlaps


def solve_part_two(sections):
    num_overlaps = 0

    for section in sections:
        one_start, one_end, two_start, two_end = section
        first_section = list(range(one_start, one_end+1))
        second_section = list(range(two_start, two_end+1))

        if set(first_section) & set(second_section):
            num_overlaps += 1

    return num_overlaps

def read_input(filename):
    sections = []
    with open(filename) as fp:
        for pair in fp.readlines():
            section_one, section_two = pair.strip().split(",")
            section_one_start = int(section_one.split("-")[0])
            section_one_end = int(section_one.split("-")[1])

            section_two_start = int(section_two.split("-")[0])
            section_two_end = int(section_two.split("-")[1])

            sections.append((section_one_start, section_one_end, section_two_start, section_two_end))

    return sections


if __name__ == "__main__":
    main()
