def main():
    grid = read_input("input.txt")
    print(solve_part_one(grid))
    print(solve_part_two(grid))


def solve_part_one(grid):
    num_rows, num_cols = len(grid), len(grid[0])
    # Since all trees on the edges are visible, we'll init the counter with all these trees
    # Count them by adding the 2 vertical edges and 2 horizontal edges and subtracting the double counted trees
    num_visible_trees = num_rows * 2 + num_cols * 2 - 4

    for row_index in range(1, num_rows - 1):
        for col_index in range(1, num_cols - 1):
            if not is_visible(row_index, col_index, grid):
                continue

            # print("{} at ({}, {}) is visible".format(grid[row_index][col_index], row_index, col_index))
            num_visible_trees += 1

    return num_visible_trees


def solve_part_two(grid):
    num_rows, num_cols = len(grid), len(grid[0])
    max_scenic_score = 0

    for row_index in range(1, num_rows - 1):
        for col_index in range(1, num_cols - 1):
            this_tree_height = grid[row_index][col_index]
            num_trees_left, num_trees_right, num_trees_top, num_trees_bottom = 0, 0, 0, 0

            # How many trees can you see from here on the LEFT?
            for j in range(col_index - 1, -1, -1):
                if grid[row_index][j] < this_tree_height:
                    num_trees_left += 1
                if grid[row_index][j] >= this_tree_height:
                    num_trees_left += 1
                    break

            # How many trees can you see from here on the RIGHT?
            for j in range(col_index + 1, len(grid[0])):
                if grid[row_index][j] < this_tree_height:
                    num_trees_right += 1
                if grid[row_index][j] >= this_tree_height:
                    num_trees_right += 1
                    break

            # How many trees can you see from here on the TOP?
            for i in range(row_index - 1, -1, -1):
                if grid[i][col_index] < this_tree_height:
                    num_trees_top += 1
                if grid[i][col_index] >= this_tree_height:
                    num_trees_top += 1
                    break

            # How many trees can you see from here on the BOTTOM?
            for i in range(row_index + 1, len(grid)):
                if grid[i][col_index] < this_tree_height:
                    num_trees_bottom += 1
                if grid[i][col_index] >= this_tree_height:
                    num_trees_bottom += 1
                    break

            # print("({}, {}) -> {}: Left: {}, Right: {}, Top: {}, Bottom: {}".format(
            #     row_index, col_index, this_tree_height, num_trees_left, num_trees_right, num_trees_top, num_trees_bottom))

            max_scenic_score = max(max_scenic_score, num_trees_left * num_trees_right * num_trees_top * num_trees_bottom)

    return max_scenic_score


# A tree is "visible" if it is visible from at least one side (left, right, top, bottom) of it.
def is_visible(row_index, col_index, grid):
    this_tree_height = grid[row_index][col_index]

    return any([
        # Is every tree to the left of this tree shorter than it?
        all(tree_height < this_tree_height for tree_height in grid[row_index][:col_index]),

        # Is every tree to the right of this tree shorter than it?
        all(tree_height < this_tree_height for tree_height in grid[row_index][col_index + 1:]),

        # Is every tree to the top of this tree shorter than it?
        all(tree_height < this_tree_height for tree_height in [grid[i][col_index] for i in range(row_index)]),

        # Is every tree to the bottom of this tree shorter than it?
        all(tree_height < this_tree_height for tree_height in [grid[i][col_index] for i in range(row_index + 1, len(grid[0]))]),
    ])


def read_input(filename):
    grid = []

    with open(filename) as fp:
        for line in fp.readlines():
            grid.append([int(tree) for tree in line.strip()])

    return grid


if __name__ == "__main__":
    main()
