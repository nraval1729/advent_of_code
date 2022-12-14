from functools import cmp_to_key
from pprint import pprint
import itertools


def main():
    packet_pairs = read_input("input.txt")
    print(solve_part_one(packet_pairs))
    print(solve_part_two(packet_pairs))


def solve_part_one(packet_pairs):
    indices_sum = 0
    for idx, packet_pairs in enumerate(packet_pairs):
        left, right = packet_pairs
        ret = recurse(left, right)
        if ret:
            indices_sum += idx + 1

    return indices_sum


def solve_part_two(packet_pairs):
    divider_packet_two = [[2]]
    divider_packet_six = [[6]]

    all_packets = [divider_packet_two, divider_packet_six]

    for packet_pair in packet_pairs:
        all_packets.append(packet_pair[0])
        all_packets.append(packet_pair[1])

    # We just need to sort this fucking list I think
    sorted_packets = sorted(all_packets, key=cmp_to_key(sort_comparator))
    return (sorted_packets.index(divider_packet_six) + 1) * (sorted_packets.index(divider_packet_two) + 1)

    # for packet_permutation in itertools.permutations(all_packets):
    #     print(packet_permutation)
    #     print("******************************************************************")
    #     if not all([recurse(packet_permutation[i], packet_permutation[i+1]) for i in range(0, len(packet_permutation)-1, 2)]):
    #         continue
    #
    #     # We've found a permutation with all packets sorted! Let's find the fucking decoder key
    #     return (packet_permutation.index(divider_packet_six)+1) * (packet_permutation.index(divider_packet_two)+1)


def sort_comparator(left, right):
    ret = recurse(left, right)
    if ret is None:
        return 0
    if ret:
        return -1
    return 1


def recurse(left, right):
    # If both are ints, we compare them directly
    if isinstance(left, int) and isinstance(right, int):
        if left < right:
            return True
        elif left > right:
            return False
        else:
            return None  # Both are equal so we'll send back None to signify continuation

    # If both are lists, we will compare each of their elements individually
    if isinstance(left, list) and isinstance(right, list):
        for left_i, right_i in itertools.zip_longest(left, right, fillvalue="-1"):
            # Left list ran out of items so packets are in order
            if left_i == "-1":
                return True
            # Right list ran out of items so packets are NOT in order
            if right_i == "-1":
                return False
            inorder = recurse(left_i, right_i)
            if inorder is None:
                continue
            return inorder

    # Finally, if one is an int and the other a list, we convert the int to a list and compare them
    if isinstance(left, int) and isinstance(right, list):
        inorder = recurse([left], right)
        if inorder is not None:
            return inorder

    if isinstance(left, list) and isinstance(right, int):
        inorder = recurse(left, [right])
        if inorder is not None:
            return inorder


def read_input(filename):
    import json

    packets = []

    with open(filename) as fp:
        for line in fp.readlines():
            line = line.strip()
            if not line:
                continue
            packets.append(json.loads(line))

    return [(packets[i], packets[i + 1]) for i in range(0, len(packets) - 1, 2)]


if __name__ == "__main__":
    main()
