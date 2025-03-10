#!/usr/bin/python3

import functools
import logging
import re
from dataclasses import dataclass
from pathlib import Path


CUR_DIR: Path = Path(__file__).parents[0]
ROOT_DIR: Path = CUR_DIR.parent
PROBLEMS_DIR: Path = ROOT_DIR / "problems"

METADATA_PATTERN = re.compile(
    r"#\s(\d*).(.*)\s*\n\n"
    + r"-\sDifficulty:\s(Easy|Medium|Hard)\s*\n"
    + r"-\sTopics:(.*)\s*\n"
    + r"-\sLink:\s((?:http|https)://.*)\s*\n",
    re.ASCII,
)

CODE_SUFFIX_MAP = {
    ".c": "C",
    ".cpp": "C++",
    ".java": "Java",
    ".py": "Python",
    ".go": "Go",
    ".rust": "Rust",
    ".js": "JS",
    ".rkt": "Racket",
}


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
    solution_files: dict[str, Path]


def main():
    logging.basicConfig(level=logging.DEBUG)

    problems = get_problems(problems_dir=PROBLEMS_DIR)
    logging.info(f"Retrieved {len(problems)} problems.")

    templ_file = CUR_DIR / "template" / "README.md"
    logging.info(f"Reading template from: {templ_file.relative_to(ROOT_DIR)}")

    readme_text = templ_file.read_text().format(
        problem_table=gen_problem_table(
            problems=sort_problems(problems=problems, sort_key="number")
        )
    )
    readme_file = ROOT_DIR / "README.md"
    readme_file.write_text(readme_text)

    logging.info(f"Updated {readme_file.relative_to(ROOT_DIR)} successfully!")


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

    def extract_metadata(doc_file: Path) -> tuple | None:
        try:
            with doc_file.open("r") as f:
                file_head = "".join(next(f) for _ in range(10))  # Read first 10 lines
        except Exception as e:
            logging.warning(f"Error reading {doc_file}: {e}")
            return None

        match = METADATA_PATTERN.search(file_head)
        if not match:
            logging.warning(f"Cannot extract metadata from {doc_file}, skipping.")
            return None

        number = int(match.group(1))
        title = match.group(2).strip()
        difficulty = match.group(3)
        topics = match.group(4).strip().split(", ")
        link = match.group(5)

        return number, title, difficulty, topics, link

    def find_solution_files(problem_dir: Path) -> dict[str, Path]:
        solution_files = {
            CODE_SUFFIX_MAP[file.suffix]: file
            for file in problem_dir.iterdir()
            if file.is_file() and file.suffix in CODE_SUFFIX_MAP
        }

        # If no direct matches, try subdirectories
        # E.g., cpp/*.cpp, java/*.java.
        if not solution_files:
            solution_files = {
                CODE_SUFFIX_MAP[f".{subdir.name}"]: subdir
                for subdir in problem_dir.iterdir()
                if subdir.is_dir() and f".{subdir.name}" in CODE_SUFFIX_MAP
            }

        return solution_files

    doc_file = problem_dir / "README.md"
    metadata = extract_metadata(doc_file)
    if not metadata:
        return None

    number, title, difficulty, topics, link = metadata
    solution_files = find_solution_files(problem_dir)

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
