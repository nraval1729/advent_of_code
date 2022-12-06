def main():
    crates, moves = read_input("input.txt")
    # print(solve_part_one(crates, moves))
    print(solve_part_two(crates, moves))


def solve_part_one(crates, moves):
    for move in moves:
        num_crates_to_move, move_from_stack, move_to_stack = move
        for num in range(num_crates_to_move):
            crates[move_to_stack - 1].append(crates[move_from_stack - 1].pop())

    message = ""

    for crate in crates:
        message += crate[-1]

    return message


def solve_part_two(crates, moves):
    for move in moves:
        num_crates_to_move, move_from_stack, move_to_stack = move

        moved_crates = []
        for num in range(num_crates_to_move):
            moved_crates.append(crates[move_from_stack - 1].pop())

        for crate_to_move in reversed(moved_crates):
            crates[move_to_stack - 1].append(crate_to_move)

    message = ""

    for crate in crates:
        message += crate[-1]

    return message


def read_input(filename):
    crates = read_crates(filename)
    moves = read_moves(filename)

    return crates, moves


def read_crates(filename):
    # return [
    #     ["Z", "N"],
    #     ["M", "C", "D"],
    #     ["P"],
    # ]

    return [
        ["G", "T", "R", "W"],
        ["G", "C", "H", "P", "M", "S", "V", "W"],
        ["C", "L", "T", "S", "G", "M"],
        ["J", "H", "D", "M", "W", "R", "F"],
        ["P", "Q", "L", "H", "S", "W", "F", "J"],
        ["P", "J", "D", "N", "F", "M", "S"],
        ["Z", "B", "D", "F", "G", "C", "S", "J"],
        ["R", "T", "B"],
        ["H", "N", "W", "L", "C"],
    ]


def read_moves(filename):
    moves = []

    with open(filename) as fp:
        for line in fp.readlines():
            if not line.startswith("move"):
                continue

            move = [int(ch) for ch in line.strip().split(" ") if ch.isdigit()]
            moves.append(move)

    return moves


if __name__ == "__main__":
    main()
