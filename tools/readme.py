#!/usr/bin/python3

# Thanks to https://github.com/Triple-Z/LeetCode/blob/master/utils/update_readme.py .

import re
import argparse
import logging
import shutil

from pathlib import Path
from urllib.parse import quote
from utils import insert_lines_after_matching


def main():
    """
    Update README.md from codes and docs.

    example_problems = {
        "1": {
            "title": "[Two Sum](https://leetcode.com/problems/...)",
            "solution": [
                "[C++](cpp/1.cpp)",
                "[Java](java/1.java)",
            ]
            "difficulty": "Easy",
            "doc": "[📃](doc/1.md)"
        }
    }

    """

    logging.basicConfig(level=logging.DEBUG)

    parser = argparse.ArgumentParser(
        description='Generate README.md')
    parser.add_argument('-o', '--output-file', dest='output_file_name',
                        type=str, default='README-generated.md')

    args = parser.parse_args()
    output_file_name = args.output_file_name

    cwd = Path(__file__).parents[0]
    root_dir = cwd / '..'
    template_file = cwd / 'template' / 'README.md'
    output_file = root_dir / output_file_name

    code_infos = {
        "cpp": {
            "pattern": re.compile(r'code\/cpp\/(\d+)\.cpp', re.ASCII),
            "filelist": Path(root_dir).glob('code/cpp/*.cpp'),
            "name": "C++"
        },
        "go": {
            "pattern": re.compile(r'code\/go\/(\d+)\.go', re.ASCII),
            "filelist": Path(root_dir).glob('code/go/*.go'),
            "name": "Go"
        },
        "java": {
            "pattern": re.compile(r'code\/java\/(\d+)\.java', re.ASCII),
            "filelist": Path(root_dir).glob('code/java/*.java'),
            "name": "Java"
        },
        "js": {
            "pattern": re.compile(r'code\/js\/(\d+)\.js', re.ASCII),
            "filelist": Path(root_dir).glob('code/js/*.js'),
            "name": "JS"
        },
        "py3": {
            "pattern": re.compile(r'code\/py3\/(\d+)\.py', re.ASCII),
            "filelist": Path(root_dir).glob('code/py3/*.py'),
            "name": "Python"
        },
        "rust": {
            "pattern": re.compile(r'code\/rust\/(\d+)\.rs', re.ASCII),
            "filelist": Path(root_dir).glob('code/rust/*.rs'),
            "name": "Rust"
        }
    }

    doc_pattern = re.compile(
        r'docs\/(\d+)\. ([\w\d\s\'\"\(\)\-\,\.]+)\.md', re.ASCII)

    problem_info_pattern = re.compile(
        r'-\sSubtitle:\s(.*)\s*\n-\sDifficulty:\s(Easy|Medium|Hard)\s*\n-\sTopics:(.*)\s*\n-\sLink:\s((?:http|https):\/\/.*)\s*\n', re.ASCII)

    doc_filelist = Path(root_dir).glob('docs/*.md')

    problems = {}
    lines = []

    for key, code_info in code_infos.items():
        code_filelist = code_info["filelist"]
        for code_file in code_filelist:
            code_res = find_match(str(code_file), code_info["pattern"])
            if(code_res) == None:
                continue

            number = code_res.group(1)
            logging.debug(f"process {number} {key} file")

            number = f"{number}"
            if number not in problems:
                problems[number] = {
                    "title": "",
                    "solution": [],
                    "difficulty": "",
                    "doc": ""
                }
            rel_code_file = code_file.relative_to(root_dir)
            code_name = code_info["name"]
            code_url = quote(str(rel_code_file))
            problems[number]["solution"].append(
                f"[{code_name}]({code_url})")

    for doc_file in doc_filelist:
        doc_res = find_match(str(doc_file), doc_pattern)
        if doc_res == None:
            continue

        number = doc_res.group(1)
        title = doc_res.group(2)
        logging.debug(f"processing {number} doc file, {title}")

        # open the doc to get problem info
        rel_doc_file = doc_file.relative_to(root_dir)
        doc_url = quote(str(rel_doc_file))
        subtitle, difficulty, topic, problem_link = None, None, None, None
        with open(str(rel_doc_file), 'r') as file:
            # read the first 10 lines for each doc file
            file_head = [next(file) for x in range(10)]
            file_head = ''.join(file_head)
            file_res = problem_info_pattern.search(file_head)
            if file_res is None:
                logging.warning(
                    'Cannot get problem info in {} , ignored'.format(doc_file))
                logging.debug(file_head)
            else:
                subtitle = file_res.group(1)
                difficulty = file_res.group(2)
                topic = file_res.group(3)
                problem_link = file_res.group(4)

        if number in problems:
            problems[number]["title"] = f"[{title}]({problem_link})"
            problems[number]["difficulty"] = f"{difficulty}"
            problems[number]["doc"] = f"[📃]({doc_url})"
        else:
            logging.warning("No code file found, ignored...")

    for number, problem in sorted(problems.items()):
        title = problem["title"]
        solution = ",".join(problem["solution"])
        difficulty = problem["difficulty"]
        doc = problem["doc"]
        line = f"| {number} | {title} | {solution} | {difficulty} | {doc} |\n"
        lines.append(line)

    shutil.copy(template_file, output_file)
    matching_line = "| --- | ----- | -------- | ---------- | ---- |\n"
    insert_lines_after_matching(output_file, matching_line, lines)


def find_match(filename: str, pattern):
    """
    return match result from pattern
    """
    matched = pattern.search(filename)
    if(matched is None):
        logging.warning(f"The filename {filename} is invalid, ignored.")
        return None
    else:
        return matched


if __name__ == "__main__":
    main()
