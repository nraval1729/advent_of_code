def main():
    risk_level = parse_input("input.txt")

    print("Solution for part 1: ", solve_part_one(risk_level))
    print("Solution for part 2: ", solve_part_two(risk_level))


def solve_part_one(risk_level):
    dp = []
    for _ in range(len(risk_level)):
        dp.append([0] * len(risk_level[0]))

    # Fill first row
    dp[0][0] = risk_level[0][0]
    for i in range(1, len(dp[0])):
        dp[0][i] = risk_level[0][i] + dp[0][i - 1]

    # Fill first column
    for i in range(1, len(dp)):
        dp[i][0] = risk_level[i][0] + dp[i-1][0]

    # Minimum risk level to get to cell (i,j) is the risk_level at (i, j) + the min of the risk levels above and to its left
    for i in range(1, len(dp)):
        for j in range(1, len(dp[0])):
            dp[i][j] = risk_level[i][j] + min(dp[i-1][j], dp[i][j-1])

    return dp[len(dp)-1][len(dp[0])-1] - dp[0][0]


def solve_part_two(risk_level):
    from pprint import pprint

    full_risk_level = []
    for _ in range(len(risk_level)*5):
        full_risk_level.append([0] * len(risk_level[0] * 5))

    # Fill initial tile aka the "original" risk_level into this full risk_level
    # Basically fill the "smaller cave" into this "larger cave"
    for i in range(len(risk_level)):
        for j in range(len(risk_level[0])):
            full_risk_level[i][j] = risk_level[i][j]

    # Fill first "column" of tiles
    for i in range(len(risk_level), len(risk_level)*5):
        for j in range(len(risk_level[0])):
            full_risk_level[i][j] = (full_risk_level[i-len(risk_level)][j] % 9) + 1

    # Fill every remaining tile
    for i in range(len(risk_level)*5):
        for j in range(len(risk_level[0]), len(risk_level[0])*5):
            full_risk_level[i][j] = (full_risk_level[i][j-len(risk_level[0])] % 9) + 1
    pprint([i for i in ])
    return solve_part_one(full_risk_level)


def parse_input(filename):
    risk_level = []
    with open(filename, "rb") as f:
        for line in f.read().splitlines():
            risk_level.append([int(v) for v in line])

    return risk_level


if __name__ == "__main__":
    main()
