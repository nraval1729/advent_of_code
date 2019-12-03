# Get input from https://adventofcode.com/2019/day/1/input
module_masses = [113348, 85199, 56077, 108338, 88842, 86765, 127040, 88557, 87886, 110849, 95682, 131611, 79658, 139439, 62467, 82333, 79414, 116672,
                 118256,
                 127660, 59774, 67990, 81653, 143436, 101701, 90571, 131983, 70494, 86232, 137759, 87992, 107601, 141275, 100261, 133153, 136706,
                 84363, 114771,
                 115167, 64509, 97324, 70627, 93215, 60459, 145358, 102741, 85623, 145687, 106837, 146764, 91837, 80190, 114857, 133985, 55423, 60612,
                 63556,
                 139326, 73907, 71478, 95854, 87886, 91624, 85675, 141381, 122392, 73756, 131710, 99053, 135684, 114005, 95885, 75545, 55703, 80835,
                 106478, 74307,
                 113562, 134192, 117605, 138161, 132905, 117676, 125103, 147899, 107373, 142169, 72084, 68682, 115345, 63130, 143231, 72135, 91780,
                 122640, 74195,
                 84365, 97015, 81773, 74146]


def compute_fuel_requirement(module_mass):
    return module_mass // 3 - 2


def compute_total_fuel_requirement(module_mass):
    total_fuel, extra_fuel = 0, compute_fuel_requirement(module_mass=module_mass)
    while extra_fuel > 0:
        total_fuel += extra_fuel
        extra_fuel = compute_fuel_requirement(module_mass=extra_fuel)

    return total_fuel


def run_tests_for_part_one():
    assert compute_fuel_requirement(14) == 2, get_error_message_for_part_one(module_mass=14, expected=2)
    assert compute_fuel_requirement(1969) == 654, get_error_message_for_part_one(module_mass=1969, expected=654)
    assert compute_fuel_requirement(100756) == 33583, get_error_message_for_part_one(module_mass=100756, expected=33583)

    print("Part one tests passed!")


def run_tests_for_part_two():
    assert compute_total_fuel_requirement(14) == 2, get_error_message_for_part_two(module_mass=14, expected=2)
    assert compute_total_fuel_requirement(1969) == 966, get_error_message_for_part_two(module_mass=1969, expected=966)
    assert compute_total_fuel_requirement(100756) == 50346, get_error_message_for_part_two(module_mass=100756, expected=50346)

    print("Part two tests passed!")


def get_error_message_for_part_two(module_mass, expected):
    return "Incorrect fuel requirement for {}. Expected {} but computed: {}".format(module_mass, expected,
                                                                                    compute_total_fuel_requirement(module_mass))


def get_error_message_for_part_one(module_mass, expected):
    return "Incorrect fuel requirement for {}. Expected {} but computed: {}".format(module_mass, expected, compute_fuel_requirement(module_mass))


def main():
    run_tests_for_part_one()
    print("Part one: {}".format(sum([compute_fuel_requirement(module_mass=m) for m in module_masses])))

    run_tests_for_part_two()
    print("Part two: {}".format(sum([compute_total_fuel_requirement(module_mass=m) for m in module_masses])))


if __name__ == "__main__":
    main()
