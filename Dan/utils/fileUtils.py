def read_input(file_name):
    file = open(file_name, "r")
    lines = file.read().splitlines()
    file.close()
    return lines
