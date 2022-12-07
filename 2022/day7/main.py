def main():
    filesystem = read_input("input.txt")
    print(solve_part_one(filesystem))
    print(solve_part_two(filesystem))


def solve_part_one(root_node):
    all_nodes = [root_node]
    nodes_of_interest = []
    while all_nodes:
        curr_node = all_nodes.pop()
        if curr_node.type == "directory":
            if curr_node.calculate_size() <= 100000:
                nodes_of_interest.append(curr_node)
            all_nodes += [child for child in curr_node.children if child.type == "directory"]

    return sum([node.calculate_size() for node in nodes_of_interest])


def solve_part_two(root_node):
    used_space = root_node.calculate_size()
    unused_space = 70000000 - used_space

    space_to_free = 30000000 - unused_space

    # Set this shit to a high number
    smallest_deletable_directory_size = used_space

    all_nodes = [root_node]
    while all_nodes:
        curr_node = all_nodes.pop()

        if curr_node.type == "directory":
            curr_node_size = curr_node.calculate_size()
            all_nodes += [child for child in curr_node.children if child.type == "directory"]
            if space_to_free <= curr_node_size <= smallest_deletable_directory_size:
                smallest_deletable_directory_size = curr_node_size

    return smallest_deletable_directory_size


def read_input(filename):
    root_node = DirNode("/", "directory", None)
    curr_dir = root_node
    with open(filename) as fp:
        for line in fp.readlines():
            line = line.strip()

            if is_command(line):
                curr_dir = handle_command(line, curr_dir)
            else:
                curr_dir = handle_output(line, curr_dir)

    return root_node


def handle_command(line, curr_dir):
    command_name = line.split()[1]

    if command_name == "cd":
        new_dir = line.split()[2]
        if new_dir == "..":
            curr_dir = curr_dir.parent
        elif new_dir == "/":
            pass
        else:
            curr_dir = curr_dir.find_child(new_dir, "directory")
    if command_name == "ls":
        pass

    return curr_dir


def handle_output(line, curr_dir):
    if line.split()[0] == "dir":
        child_dir_name = line.split()[1]
        curr_dir.add_child(DirNode(name=child_dir_name, type="directory", parent=curr_dir))
    if line.split()[0].isdigit():
        filename = line.split()[1]
        curr_dir.add_child(DirNode(name=filename, type="file", parent=curr_dir, size=int(line.split()[0])))

    return curr_dir


def is_command(line):
    return line.startswith("$")


class DirNode(object):

    def __init__(self, name, type, parent, size=0):
        self.name = name
        self.type = type
        self.children = []
        self.parent = parent
        self.size = size

    def add_child(self, childNode):
        self.children.append(childNode)

    def calculate_size(self):
        if self.type == "file":
            return self.size
        else:
            return sum([child.calculate_size() for child in self.children])

    def find_child(self, name, type):
        for child in self.children:
            if child.get_name() == name and child.get_type() == type:
                return child

    def get_name(self):
        return self.name

    def get_type(self):
        return self.type

    def get_children_names(self):
        return ", ".join(child.get_name() for child in self.children)

    def __str__(self):
        return "{} - {} whose parent is {} and children are {}".format(
            self.type, self.name, self.parent.get_name() if self.parent else "None", self.get_children_names()
        )


if __name__ == "__main__":
    main()
