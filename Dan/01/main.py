
SUM_TARGET = 2020


def read_input(file_name):
    file = open(file_name, "r")
    lines = file.read().splitlines()
    file.close()
    return lines


def find_entries(content, num):
    i = 0

    while i < len(content):
        entryA = int(content[i])
        j = 0
        while j < len(content):
            entryB = int(content[j])
            k = 0
            while k < len(content):
                entryC = int(content[k])
                sm = entryA + entryB + entryC
                if sm == SUM_TARGET:
                    return entryA, entryB, entryC
                k = k +1
            j = j + 1
        i = i + 1


def main():
    content = read_input("input.txt")
    entryA, entryB, entryC = find_entries(content, 3)
    results = entryA * entryB * entryC
    print("results: {}".format(results))


if __name__ == "__main__":
    main()
