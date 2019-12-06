from collections import defaultdict


def are_digits_monotonically_increasing(number):
    return str(number) == ''.join(sorted(str(number)))


def are_at_least_two_adjacent_digits_alike(number):
    number_string = str(number)
    for i in range(1, len(number_string)):
        if number_string[i] == number_string[i - 1]:
            return True
    return False


# Here we can assume that number passes the monotonicity check for
# the purposes of this problem.
def is_step_two_rule_applied(number):
    digit_to_occurences = defaultdict(int)

    for digit in str(number):
        digit_to_occurences[digit] += 1

    return 2 in digit_to_occurences.values()


def count_possible_passwords(start, stop):
    step_1_count, step_2_count = 0, 0
    for n in range(start, stop + 1):
        if are_digits_monotonically_increasing(n) and are_at_least_two_adjacent_digits_alike(n):
            step_1_count += 1
        if are_digits_monotonically_increasing(n) and is_step_two_rule_applied(n):
            step_2_count += 1

    return step_1_count, step_2_count


def main():
    step_1_count, step_2_count = count_possible_passwords(382345, 843167)
    print("Step 1: {}, Step 2: {}".format(step_1_count, step_2_count))


if __name__ == "__main__":
    main()
