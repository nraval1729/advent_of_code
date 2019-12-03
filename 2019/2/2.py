# Get input from https://adventofcode.com/2019/day/2/input
intcode_program = [1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 10, 23, 2, 13, 23, 27, 1, 5, 27, 31, 2, 6, 31, 35, 1, 6, 35,
                   39, 2, 39, 9, 43, 1, 5, 43, 47, 1, 13, 47, 51, 1, 10, 51, 55, 2, 55, 10, 59, 2, 10, 59, 63, 1, 9, 63, 67, 2, 67, 13, 71, 1, 71, 6,
                   75, 2, 6, 75, 79, 1, 5, 79, 83, 2, 83, 9, 87, 1, 6, 87, 91, 2, 91, 6, 95, 1, 95, 6, 99, 2, 99, 13, 103, 1, 6, 103, 107, 1, 2, 107,
                   111, 1, 111, 9, 0, 99, 2, 14, 0, 0]

# Restore gravity assist program to state before crash
NOUN_PART_ONE = 12
VERB_PART_ONE = 2

POSITION_STEP = 4

ADDITION_OPCODE = 1
MULTIPLICATION_OPCODE = 2
HALT_OPCODE = 99


def run_tests_part_one():
    assert run_intcode_program([1, 0, 0, 0, 99]) == [2, 0, 0, 0, 99]
    assert run_intcode_program([2, 3, 0, 3, 99]) == [2, 3, 0, 6, 99]
    assert run_intcode_program([2, 4, 4, 5, 99, 0]) == [2, 4, 4, 5, 99, 9801]
    assert run_intcode_program([1, 1, 1, 4, 99, 5, 6, 0, 99]) == [30, 1, 1, 4, 2, 5, 6, 0, 99]
    assert run_intcode_program([1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50]) == [3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50]

    print("Tests for part one passed!")


def run_intcode_program(intcode_program):
    for i in range(0, len(intcode_program), POSITION_STEP):
        opcode = intcode_program[i]

        if opcode == HALT_OPCODE:
            return intcode_program

        arg_1, arg_2, result_index = intcode_program[intcode_program[i + 1]], intcode_program[intcode_program[i + 2]], intcode_program[i + 3]
        if opcode == ADDITION_OPCODE:
            intcode_program[result_index] = arg_1 + arg_2
        elif opcode == MULTIPLICATION_OPCODE:
            intcode_program[result_index] = arg_1 * arg_2
        else:
            raise Exception("Unknown opcode {}. Expected opcodes {}".format(opcode, [ADDITION_OPCODE, MULTIPLICATION_OPCODE, HALT_OPCODE]))


def set_noun_and_verb(noun, verb, intcode_program):
    intcode_program[1] = noun
    intcode_program[2] = verb


def main():
    run_tests_part_one()

    intcode_program_copy = intcode_program[::]
    set_noun_and_verb(noun=NOUN_PART_ONE, verb=VERB_PART_ONE, intcode_program=intcode_program_copy)
    output = run_intcode_program(intcode_program_copy)
    print("Part one answer: {}".format(output[0]))

    for noun in range(0, 100):
        for verb in range(0, 100):
            intcode_program_copy = intcode_program[::]
            set_noun_and_verb(noun=noun, verb=verb, intcode_program=intcode_program_copy)
            if run_intcode_program(intcode_program_copy)[0] == 19690720:
                print("Part two answer: {}".format(100 * noun + verb))


if __name__ == "__main__":
    main()
