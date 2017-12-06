#!/usr/bin/env python
"""This script generates personal info for mocking purpose"""

import uuid
import argparse
from mimesis import Personal

def parse_args():
    """Configure parser and parse command line arguments"""
    parser = argparse.ArgumentParser(description='Generate mock data.')
    parser.add_argument('file', type=str, help='file to generate')
    parser.add_argument('idfrom', type=int, help='id start value')
    parser.add_argument('idto', type=int, help='id end value (excluding)')
    parser.add_argument('locale', metavar='locale', type=str, choices=['en', 'de', 'ru', 'ja'],
                        help='locale for names')
    return parser.parse_args()

def generate(filename, idfrom, idto, locale):
    """Generate and write data to the file"""
    person = Personal(locale)
    with open(filename, 'wt') as file:
        for i in range(idfrom, idto + 1):
            print(i, uuid.uuid4().hex, person.name(), person.surname(), person.email(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  person.age(), person.age(), person.age(), person.age(), person.age(),
                  sep='|', file=file)

if __name__ == '__main__':
    vals = parse_args()
    generate(vals.file, vals.idfrom, vals.idto, vals.locale)
