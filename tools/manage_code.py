#!/usr/bin/python3
import subprocess
import argparse
import os
import sys


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "manage", help="add: create code file; commit: git-commit code file")
    args = parser.parse_args()
    manage_option = args.manage

    number = input("Input problem number:")
    language = input(
        "Input problem language(c:C++, g:Go, j:Java, n:Node.js, p:Python3, r:Rust):")
    language = {
        'c': 'cpp',
        'g': 'go',
        'j': 'java',
        'n': 'js',
        'p': 'py3',
        'r': 'rs'
    }.get(language)
    if language == "py3":
        py_suffix = language.replace("3", "")
        code_file = f"./{language}/{number}.{py_suffix}"
    else:
        code_file = f"./{language}/{number}.{language}"

    if manage_option == "add":
        if os.path.exists(code_file):
            print("Code file exist!")
            sys.exit(1)
        else:
            open(code_file, "w")
        print(f"\nCreate {language} file at {code_file} successfully.")


if __name__ == "__main__":
    main()
