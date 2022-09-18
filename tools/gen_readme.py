#!/usr/bin/python3

import re
import shutil
import copy
import logging
import functools

from pathlib import Path
from urllib.parse import quote

from utils import sed


CUR_DIR: Path = Path(__file__).parents[0]
ROOT_DIR: Path = CUR_DIR / '..'
PROBLEMS_DIR: Path = ROOT_DIR / 'problems'

PROBLEM_TEMPL = {
    "number": 0,
    "title": "",
    "difficulty": "",
    "link": "",
    "topics": [],
    "problem_dir": Path(),
    "solution_files": {},
    "doc_file": "",
}


def main():
    logging.basicConfig(level=logging.DEBUG)
    problems = get_problems(problems_dir=PROBLEMS_DIR)
    content = gen_content(problems=problems, sort_key="number")

    template_file = CUR_DIR / 'template' / 'README.md'
    output_file = ROOT_DIR / 'README.md'
    shutil.copy(template_file, output_file)
    sed("<SOLUTION_LIST>", content, str(output_file))


def gen_content(problems: list, sort_key: str) -> str:
    headers = [
        "| #   | Title | Difficulty | Solution | Doc |",
        "| --- | ----- | ---------- | -------- | --- |"
    ]
    lines = [*headers]

    def compare(a, b):
        if a[sort_key] < b[sort_key]:
            return -1
        elif a[sort_key] > b[sort_key]:
            return 1
        else:
            return 0

    problems.sort(key=functools.cmp_to_key(compare))

    for problem in problems:
        number = problem["number"]
        title = f"[{problem['title']}]({problem['link']})"
        difficulty = problem["difficulty"]

        # special for github: use `problem_dir` instead of `doc_file`,
        # because github will render README.md in directory automatically.
        # but it may lead to ambiguity.
        doc = f"[📃]({problem['problem_dir'].relative_to(ROOT_DIR)})"

        solutions = []
        for key, value in problem["solution_files"].items():
            solutions.append(f"[{key}]({value.relative_to(ROOT_DIR)})")
        solution = ", ".join(solutions)

        lines.append(
            f"| {number} | {title} | {difficulty} | {solution} | {doc} |")

    return "\n".join(lines)


def get_problems(problems_dir: Path) -> list:
    """
    example:
    problems = [
        {
            "number": 1,
            "title": "Two Sum",
            "difficulty": "Easy",
            "topics": [
                "Array",
                "Hash Table"
            ],
            "link": "https://leetcode.com/problems/two-sum/",
            "problem_dir": Path("two-sum"),
            "solution_files": {
                "C++": Path("two-sum/answer.cpp")
            }
            "doc_file": Path("two-sum/README.md"),
        }
    ]
    """

    problems = []

    folders: list[Path] = []
    for child in PROBLEMS_DIR.iterdir():
        if child.is_dir():
            folders.append(child)

    for folder in folders:
        problem = get_problem(folder)
        if problem is not None:
            problems.append(problem)

    return problems


def get_problem(problem_dir: Path) -> dict:

    number = 0
    problem = copy.deepcopy(PROBLEM_TEMPL)

    problem["problem_dir"] = problem_dir

    # extract info from doc_file
    problem_info_pattern = re.compile(
        r'#\s(\d*).(.*)\s*\n\n'
        + r'-\sDifficulty:\s(Easy|Medium|Hard)\s*\n'
        + r'-\sTopics:(.*)\s*\n'
        + r'-\sLink:\s((?:http|https):\/\/.*)\s*\n', re.ASCII)

    doc_file = problem_dir / 'README.md'

    if not doc_file.exists():
        return problem

    number, title, difficulty, topics, link = None, None, None, None, None

    # read n lines from file
    n = 10
    with open(str(doc_file), 'r') as f:
        file_head = ''.join([next(f) for x in range(n)])
        matches = problem_info_pattern.search(file_head)
        if matches is None:
            logging.warning(
                f'Cannot get info from {doc_file}, file head = {file_head}.')
        else:
            number = matches.group(1)
            title = matches.group(2)
            difficulty = matches.group(3)
            topics = matches.group(4).strip().split(', ')
            link = matches.group(5)

    problem["number"] = int(number)
    problem["title"] = title
    problem["difficulty"] = difficulty
    problem["topics"] = topics
    problem["link"] = link

    # match files
    doc_name_by_suffix = {
        ".txt": "Text",
        ".md": "Markdown",
    }

    code_name_by_suffix = {
        ".cpp": "C++",
        ".java": "Java",
        ".py": "Python",
        ".go": "Go",
        ".rust": "Rust",
        ".js": "JS",
    }

    files: list[Path] = []
    for child in problem_dir.iterdir():
        if child.is_file():
            files.append(child)

    for file in files:
        suffix = file.suffix
        # match code files
        if suffix in code_name_by_suffix.keys():
            problem["solution_files"][code_name_by_suffix[suffix]] = file
        # match doc file (only keep the last one)
        if suffix in doc_name_by_suffix.keys():
            problem["doc_file"] = file

    return problem


if __name__ == '__main__':
    main()
