
import sys
import os
sys.path.append(os.path.abspath('../utils'))
from fileUtils import read_input

def get_traverse_path(hor, ver, ver_max):
    x, y = (0,0)
    coordinates = [(x, y)]
    while y < ver_max - 1:
        for i in range(hor):
            x = x + 1
        for j in range(ver):
            y = y + 1
        coordinates.append((x, y))
    return coordinates

def isTree(coordinate, map):
    x, y = coordinate
    hor_max = len(map[y])
    x = x % (hor_max)
    return map[y][x % (hor_max)] == "#"

def countTrees(results):
    return sum(results)

def main():
    content = read_input("input.txt")
    hor_slope = 3
    ver_slope = 1
    ver_max = len(content)
    path = get_traverse_path(hor_slope, ver_slope, ver_max)
    results = [isTree(coordinate, content) for coordinate in path]
    treeCount = countTrees(results)
    print("Tree count: {}".format(treeCount))

if __name__ == "__main__":
    main()