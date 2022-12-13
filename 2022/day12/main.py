import math
from pprint import pprint
from collections import defaultdict


def main():
    elevation_graph, starts, end = read_input("input.txt")
    print(solve_part_one(elevation_graph, starts[0], end))
    print(solve_part_two(elevation_graph, starts, end))


def solve_part_one(elevation_graph, start, end):
    dist, prev = dijkstra(elevation_graph, start[0], start[1])

    return dist[end]


def solve_part_two(elevation_graph, starts, end):
    min_dist = math.inf

    for start in starts:
        dist, prev = dijkstra(elevation_graph, start[0], start[1])
        min_dist = min(min_dist, dist[end])

    return min_dist


def dijkstra(elevation_graph, source_ri, source_ci):
    dist, prev = {}, {}

    unvisited_queue = {}

    for vertex in [node for node in elevation_graph]:
        dist[vertex] = math.inf
        prev[vertex] = None
        unvisited_queue[vertex] = True

    # Set start node to have 0 distance
    dist[(source_ri, source_ci)] = 0

    while unvisited_queue:
        vertex_min_dist = find_min_dist_vertex(unvisited_queue, dist)
        unvisited_queue.pop(vertex_min_dist)

        for neighbor_ri, neighbor_ci, neighbor_distance_from_min_dist_vertex in elevation_graph[vertex_min_dist]:
            neighbor = (neighbor_ri, neighbor_ci)
            # Already seen this neighbor so ignore
            if neighbor not in unvisited_queue:
                continue

            new_distance = dist[vertex_min_dist] + neighbor_distance_from_min_dist_vertex
            if new_distance < dist[neighbor]:
                dist[neighbor] = new_distance
                prev[neighbor] = vertex_min_dist

    return dist, prev


def find_min_dist_vertex(unvisited, dist):
    min_dist, min_dist_vertex = math.inf, None

    for vertex in unvisited:
        if dist[vertex] <= min_dist:
            min_dist = dist[vertex]
            min_dist_vertex = vertex

    return min_dist_vertex


def read_input(filename):
    elevation_grid = []

    # defaultdict to prevent stupid if checks
    elevation_graph = defaultdict(list)
    start, end = None, None
    starts = []

    # First we construct a 2D list of lists to represent the elevation grid
    with open(filename) as fp:
        for line in fp.readlines():
            elevation_grid.append([c for c in line.strip()])

    # We will now use the 2D list to construct the elevation graph
    for ri, row in enumerate(elevation_grid):
        for ci, col in enumerate(row):
            if col == "E":
                end = (ri, ci)
            if col == "S":
                start = (ri, ci)
            if col == "a":
                starts.append((ri, ci))
            for i, j in find_valid_neighbors(ri, ci, elevation_grid):
                elevation_graph[(ri, ci)].append((i, j, 1))

    return elevation_graph, [start] + starts, end


def find_valid_neighbors(ri, ci, elevation_grid):
    valid_neighbors = []
    if is_in_grid(ri + 1, ci, elevation_grid) and is_gettable_from(ri + 1, ci, ri, ci, elevation_grid):
        valid_neighbors.append((ri + 1, ci))
    if is_in_grid(ri - 1, ci, elevation_grid) and is_gettable_from(ri - 1, ci, ri, ci, elevation_grid):
        valid_neighbors.append((ri - 1, ci))
    if is_in_grid(ri, ci + 1, elevation_grid) and is_gettable_from(ri, ci + 1, ri, ci, elevation_grid):
        valid_neighbors.append((ri, ci + 1))
    if is_in_grid(ri, ci - 1, elevation_grid) and is_gettable_from(ri, ci - 1, ri, ci, elevation_grid):
        valid_neighbors.append((ri, ci - 1))

    return valid_neighbors


def is_in_grid(ri, ci, elevation_grid):
    return 0 <= ri < len(elevation_grid) and 0 <= ci < len(elevation_grid[0])


def is_gettable_from(dest_ri, dest_ci, source_ri, source_ci, elevation_grid):
    destination, source = elevation_grid[dest_ri][dest_ci], elevation_grid[source_ri][source_ci]

    # Start (S) is "a"
    if source == "S":
        source = "a"
    if source == "E":
        source = "z"
    if destination == "E":
        destination = "z"
    if destination == "S":
        destination = "a"

    return \
        (ord(destination) - ord(source) == 1) or ord(destination) <= ord(source)


if __name__ == "__main__":
    main()
