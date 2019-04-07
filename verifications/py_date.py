#!/usr/bin/env python3
'''
Display the result of applying the provided 'fmt' with the provided 'date'
'''


import sys
from datetime import datetime


def main() -> None:
    if len(sys.argv) != 3:
        print("Usage: py-date.py format date", file=sys.stderr)
        sys.exit(1)

    fmt = sys.argv[1]
    date = datetime.strptime(sys.argv[2], "%c")

    print(date.strftime(fmt), end='')


if __name__ == '__main__':
    main()
