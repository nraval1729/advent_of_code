enemy_move_to_my_move_win = {
    "C": "X",  # I choose rock when enemy chooses scissors
    "A": "Y",  # I choose paper when enemy chooses rock
    "B": "Z",  # I choose scissors when enemy chooses paper
}
enemy_move_to_my_move_draw = {
    "A": "X",
    "B": "Y",
    "C": "Z",
}
enemy_move_to_my_move_loss = {
    "C": "Y",  # I choose paper when enemy chooses scissors
    "A": "Z",  # I choose scissors when enemy chooses rock
    "B": "X",  # I choose rock when enemy chooses paper
}

shape_scores = {
    "X": 1,
    "Y": 2,
    "Z": 3,
}


def main():
    strategy = read_input("input.txt")

    print(solve_part_one(strategy))
    print(solve_part_two(strategy))


def solve_part_one(strategy):
    score = 0

    for enemy_move, my_move in strategy:
        if my_move == enemy_move_to_my_move_win[enemy_move]:  # I won
            score += shape_scores[my_move] + 6
        elif my_move == enemy_move_to_my_move_draw[enemy_move]:  # I drew
            score += shape_scores[my_move] + 3
        else:  # I lost
            score += shape_scores[my_move]

    return score


def solve_part_two(strategy):
    score = 0

    for enemy_move, outcome in strategy:
        if outcome == "X":  # I need to lose
            score += shape_scores[enemy_move_to_my_move_loss[enemy_move]]
        elif outcome == "Y":  # I need to draw
            score += shape_scores[enemy_move_to_my_move_draw[enemy_move]] + 3
        else:
            score += shape_scores[enemy_move_to_my_move_win[enemy_move]] + 6

    return score


def read_input(filename):
    strategy = []
    with open(filename) as file_pointer:
        for line in file_pointer.readlines():
            strategy.append([move for move in line.split()])

    return strategy


if __name__ == "__main__":
    main()
