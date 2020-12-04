import os
import sys

sys.path.append(os.path.abspath('../utils'))
from itertools import groupby
import json

from fileUtils import read_input

VALID_FIELDS = [
    "byr",
    "iyr",
    "eyr",
    "hgt",
    "hcl",
    "ecl",
    "pid"
]

def get_passports(lines):
    grouped_lines = [" ".join(list(g)).replace(" ", ", ") for k, g in groupby(lines, key=bool) if k]
    passports = [{key_val_str.split(':')[0]: key_val_str.split(':')[1] for key_val_str in group.split(' ')} for group in grouped_lines ]
    return passports


def is_passport_valid(passport):
    return all(field in passport.keys() for field in VALID_FIELDS)


def count_valid_passports(results):
    return sum(results)

def main():
    content = read_input("input.txt")
    passports = get_passports(content)
    results = [is_passport_valid(passport) for passport in passports]
    valid_passports = count_valid_passports(results)
    print("Valid passports: {}".format(valid_passports))

    
if __name__ == "__main__":
    main()
