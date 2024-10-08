#!/usr/bin/python3

import functools
import logging
import re
from dataclasses import dataclass
from pathlib import Path


CUR_DIR: Path = Path(__file__).parents[0]
ROOT_DIR: Path = CUR_DIR / ".."
PROBLEMS_DIR: Path = ROOT_DIR / "problems"


@dataclass(frozen=True)
class Problem(object):
    """
    Example problem:
    {
        "number": 1,
        "title": "Two Sum",
        "difficulty": "Easy",
        "topics": [
            "Array",
            "Hash Table"
        ],
        "link": "https://leetcode.com/problems/two-sum/",
        "folder": Path("two-sum"),
        "solution_files": {
            "C++": Path("two-sum/answer.cpp")
        }
        "doc_file": Path("two-sum/README.md"),
    }
    """

    number: int
    title: str
    difficulty: str
    link: str
    topics: list[str]
    folder: Path
    doc_file: Path
    solution_files: dict


def main():
    logging.basicConfig(level=logging.DEBUG)

    problems = get_problems(problems_dir=PROBLEMS_DIR)
    templ_file = CUR_DIR / "template" / "README.md"
    readme_text = templ_file.read_text().format(
        problem_table=gen_problem_table(
            problems=sort_problems(problems=problems, sort_key="number")
        )
    )
    readme_file = ROOT_DIR / "README.md"
    readme_file.write_text(readme_text)

    logging.info("Updated README.md successfully!")


def gen_problem_table(problems: list[Problem]) -> str:
    body = []
    for problem in problems:
        number = problem.number
        title = f"[{problem.title}]({problem.link})"
        difficulty = problem.difficulty

        # Special for GitHub: use `folder` instead of `doc_file`,
        # because GitHub will render README.md automatically,
        # but it may lead to ambiguity.
        doc = f"[📃]({problem.folder.relative_to(ROOT_DIR)})"

        solutions = []
        for key, value in problem.solution_files.items():
            solutions.append(f"[{key}]({value.relative_to(ROOT_DIR)})")
        solution = ", ".join(solutions)

        body.append(f"| {number} | {title} | {difficulty} | {solution} | {doc} |")

    head = [
        "| #   | Title | Difficulty | Solution | Doc |",
        "| --- | ----- | ---------- | -------- | --- |",
    ]
    table = "\n".join([*head, *body])

    return table


def sort_problems(
    problems: list[Problem], sort_key: str, reverse=False
) -> list[Problem]:
    def compare(a, b):
        if getattr(a, sort_key) < getattr(b, sort_key):
            return -1
        elif getattr(a, sort_key) > getattr(b, sort_key):
            return 1
        else:
            return 0

    return sorted(problems, key=functools.cmp_to_key(compare), reverse=reverse)


def get_problems(problems_dir: Path) -> list[Problem]:
    problems = []

    folders: list[Path] = []
    for child in problems_dir.iterdir():
        if child.is_dir():
            folders.append(child)

    for folder in folders:
        problem = get_problem(folder)
        if problem is not None:
            problems.append(problem)

    return problems


def get_problem(problem_dir: Path) -> Problem | None:
    # Extract metadata from doc
    doc_file = problem_dir / "README.md"

    pattern = re.compile(
        r"#\s(\d*).(.*)\s*\n\n"
        + r"-\sDifficulty:\s(Easy|Medium|Hard)\s*\n"
        + r"-\sTopics:(.*)\s*\n"
        + r"-\sLink:\s((?:http|https)://.*)\s*\n",
        re.ASCII,
    )

    number, title, difficulty, topics, link = None, None, None, None, None
    with open(doc_file, "r") as f:
        file_head = "".join([next(f) for x in range(10)])
        matches = pattern.search(file_head)
        if matches is None:
            logging.warning(f"Cannot get metadata from {doc_file}, Skipped.")
            return None
        number = int(matches.group(1))
        title = matches.group(2).strip()
        difficulty = matches.group(3)
        topics = matches.group(4).strip().split(", ")
        link = matches.group(5)

    # Match solution files
    # Try to match code file by suffix.
    # E.g., answer.cpp, solution.java.
    files: list[Path] = []
    dirs: list[Path] = []
    for child in sorted(problem_dir.iterdir()):
        if child.is_file():
            files.append(child)
        elif child.is_dir():
            dirs.append(child)

    code_name_by_suffix = {
        ".c": "C",
        ".cpp": "C++",
        ".java": "Java",
        ".py": "Python",
        ".go": "Go",
        ".rust": "Rust",
        ".js": "JS",
        ".rkt": "Racket",
    }

    solution_files = {}
    for file in files:
        suffix = file.suffix
        if suffix in code_name_by_suffix.keys():
            solution_files[code_name_by_suffix[suffix]] = file

    # If file suffix matches none, try to match sub folder.
    # E.g., cpp/*.cpp, java/*.java.
    if len(solution_files) == 0:
        for dir in dirs:
            name = f".{dir.name}"
            if name in code_name_by_suffix.keys():
                solution_files[code_name_by_suffix[name]] = dir

    return Problem(
        number=number,
        title=title,
        difficulty=difficulty,
        link=link,
        topics=topics,
        folder=problem_dir,
        doc_file=doc_file,
        solution_files=solution_files,
    )


if __name__ == "__main__":
    main()
