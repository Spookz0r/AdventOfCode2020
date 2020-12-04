import json
import os
import re
import sys
from itertools import groupby

# from fileUtils import read_input

sys.path.append(os.path.abspath('../utils'))


def read_input(file_name):
    file = open(file_name, "r")
    lines = file.read().splitlines()
    file.close()
    return lines


VALID_FIELDS = [
    "byr",
    "iyr",
    "eyr",
    "hgt",
    "hcl",
    "ecl",
    "pid"
]

VALID_ECL = [
    "amb",
    "blu",
    "brn",
    "gry",
    "grn",
    "hzl",
    "oth"
]

HCL_EXPR = re.compile("#[a-f0-9]{6}")

PID_EXPR = re.compile("[0-9]{9}")


def get_passports(lines):
    grouped_lines = [" ".join(list(g)).replace(" ", " ")
                     for k, g in groupby(lines, key=bool) if k]
    passports = [{key_val_str.split(':')[0]: key_val_str.split(
        ':')[1] for key_val_str in group.split(' ')} for group in grouped_lines]
    return passports


def validate_byr(val):
    return 1920 <= int(val) <= 2002


def validate_iyr(val):
    return 2010 <= int(val) <= 2020


def validate_eyr(val):
    return (len(val) == 4) and (2020 <= int(val) <= 2030)


def validate_hgt(val):
    metric = val[-2:]
    height = int(val[:-2])
    if (metric == "cm"):
        return 150 <= height <= 193
    else:
        return 59 <= height <= 76


def validate_hcl(val):
    return bool(HCL_EXPR.match(val))


def validate_ecl(val):
    return val in VALID_ECL


def validate_pid(val):
    return (bool(PID_EXPR.match(val)) and len(val) == 9)


def is_passport_valid(passport):
    passport_fields = passport.keys()
    return (
        all(field in passport_fields for field in VALID_FIELDS) and
        validate_byr(passport["byr"]) and
        validate_iyr(passport["iyr"]) and
        validate_eyr(passport["eyr"]) and
        validate_hgt(passport["hgt"]) and
        validate_hcl(passport["hcl"]) and
        validate_ecl(passport["ecl"]) and
        validate_pid(passport["pid"])
    )


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
