def main():
    head_movements = read_input("input.txt")
    print(solve_part_one(head_movements, 2, debug=False))
    print(solve_part_two(head_movements, 10, debug=False))


def solve_part_one(head_movements, rope_len, debug):
    rope = [(0, 0) for _ in range(rope_len)]

    tail_visited_pos = {(0, 0): True}

    for head_movement in head_movements:
        if debug:
            print("Move: {}".format(head_movement))
        do_string_movement(head_movement, rope, tail_visited_pos, debug)

    return len(tail_visited_pos)


def do_string_movement(head_movement, rope, tail_visited_pos, debug=False):
    head_direction, head_distance = head_movement

    for j in range(1, int(head_distance) + 1):
        rope[0] = do_head_movement(head_direction, rope[0][0], rope[0][1])
        for i in range(len(rope) - 1):
            if head_touches_tail(rope[i][0], rope[i][1], rope[i + 1][0], rope[i + 1][1]):
                continue
            rope[i + 1] = do_tail_movement(rope[i][0], rope[i][1], rope[i + 1][0], rope[i + 1][1])

        tail_visited_pos[rope[-1]] = True
        if debug:
            print("Rope after {} {}: ".format(head_direction, j), rope)


def do_head_movement(head_direction, hx, hy):
    if head_direction == "R":
        return hx + 1, hy
    elif head_direction == "L":
        return hx - 1, hy
    elif head_direction == "U":
        return hx, hy + 1
    else:
        return hx, hy - 1


def head_touches_tail(hx, hy, tx, ty):
    return abs(hx - tx) <= 1 and abs(hy - ty) <= 1


def do_tail_movement(hx, hy, tx, ty):
    if hx == tx:
        if hy > ty:
            return tx, ty + 1
        else:
            return tx, ty - 1
    if hy == ty:
        if hx > tx:
            return tx + 1, ty
        else:
            return tx - 1, ty

    # Not in same row or column
    if hx > tx and hy > ty:
        return tx + 1, ty + 1  # top right
    if hx > tx and hy < ty:
        return tx + 1, ty - 1  # bottom right
    if hx < tx and hy < ty:
        return tx - 1, ty - 1  # bottom left
    if hx < tx and hy > ty:
        return tx - 1, ty + 1  # top left


def solve_part_two(head_movements, rope_len, debug):
    return solve_part_one(head_movements, rope_len, debug)


def read_input(filename):
    input = []
    with open(filename) as fp:
        for line in fp.readlines():
            line = line.strip()
            (direction, distance) = line.split()
            input.append((direction, distance))

    return input


if __name__ == "__main__":
    main()
