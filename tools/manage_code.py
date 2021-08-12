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
    if language == "java":
        code_file = f"./{language}/src/{number}.{py_suffix}"
    else:
        code_file = f"./{language}/{number}.{language}"

    if manage_option == "add":
        if os.path.exists(code_file):
            print("Code file exist!")
            sys.exit(1)
        else:
            open(code_file, "w")
        print(f"\nCreate {language} file at {code_file} successfully.")
    elif manage_option == "commit":
        if not os.path.exists(code_file):
            print("Code file not exist!")
            sys.exit(1)
        # Check if file is tracked
        output = subprocess.run(
            ["git", "ls-files", f"{code_file}"], capture_output=True, text=True)
        if output.stdout == '':
            action = "AC"
        else:
            action = "Update"
        subprocess.run(
            ["git", "add", f"{code_file}"], capture_output=True, text=True)
        commit_out = subprocess.run(
            ["git", "commit", "-S", "-m", f"{action} {number} with {language}"], capture_output=True, text=True)
        print(commit_out.stdout)
        print(f"\nCommit {language} file at {code_file} successfully.")


if __name__ == "__main__":
    main()
