def main():
    datastream = read_input("input.txt")
    print(solve_part_one(datastream))
    print(solve_part_two(datastream))


def solve_part_one(datastream):
    return solve(datastream, 4)


def solve_part_two(datastream):
    return solve(datastream, 14)


def solve(datastream, marker_len):
    for i in range(0, len(datastream) - marker_len - 1):
        if len(set(datastream[i:i + marker_len])) == len(datastream[i:i + marker_len]):
            return i + marker_len
    return 0


def read_input(filename):
    with open(filename) as fp:
        return fp.readlines()[0]


if __name__ == "__main__":
    main()
