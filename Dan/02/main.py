import sys
import os
sys.path.append(os.path.abspath('../utils'))
from fileUtils import read_input

def extract_info(line):
    info = line.split(" ")
    info[1] = info[1].strip(":")
    return {
        "password": info[2],
        "letter": info[1].strip(":"),
        "occurence": [int(num) for num in info[0].split("-")]
    }

def is_password_valid(info):
    letter_count = info["password"].count(info["letter"])
    start, end = info["occurence"]
    return start <= letter_count <= end

def count_valid_passwords(results):
    return sum(results)

def main():
    content = read_input("input.txt")
    password_info_list = [extract_info(line) for line in content]
    results = [is_password_valid(info) for info in password_info_list]
    num_valid_passwords = count_valid_passwords(results)
    print("Valid passwords: {}".format(num_valid_passwords))
    

if __name__ == "__main__":
    main()