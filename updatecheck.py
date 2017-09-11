#!/usr/bin/python
import json
import sys

if __name__ == "__main__":
    with open('results.json') as data_file:
        data = json.load(data_file)
    if len(sys.argv) == 1:
        print(data)
        sys.exit(0)
    if len(sys.argv) == 2:
        print(data[sys.argv[1]])
    if len(sys.argv) == 3:
        print(data[sys.argv[1]][sys.argv[2]])
    if len(sys.argv) == 4:
        print(data[sys.argv[1]][sys.argv[2]][sys.argv[3]])
