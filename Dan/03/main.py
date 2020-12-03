
import sys
import os
sys.path.append(os.path.abspath('../utils'))
from fileUtils import read_input

import math

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

def is_tree(coordinate, map):
    x, y = coordinate
    hor_max = len(map[y])
    x = x % (hor_max)
    return map[y][x % (hor_max)] == "#"

def count_trees(results):
    return sum(results)

def get_slope_trees(map, slope):
    hor_slope, ver_slope = slope
    ver_max = len(map)
    path = get_traverse_path(hor_slope, ver_slope, ver_max)
    results = [is_tree(coordinate, map) for coordinate in path]
    treeCount = count_trees(results)
    return treeCount
    print("Tree count: {}".format(treeCount))

def main():
    content = read_input("input.txt")
    slopes = [
        (1, 1),
        (3, 1),
        (5, 1),
        (7, 1),
        (1, 2)
    ]
    trees = [get_slope_trees(content, slope) for slope in slopes]
    print(math.prod(trees))


if __name__ == "__main__":
    main()
