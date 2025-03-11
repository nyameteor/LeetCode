#!/usr/bin/python3

import logging
import re
from dataclasses import dataclass
from pathlib import Path

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger("update_doc")

CUR_DIR = Path(__file__).parent
ROOT_DIR = CUR_DIR.parent


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
        "directory": Path("two-sum"),
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
    directory: Path
    doc_file: Path
    solution_files: dict[str, Path]


def main():
    problems_dir = ROOT_DIR / "problems"
    problems = get_problems(problems_dir=problems_dir)
    problems = sorted(problems, key=lambda p: p.number)
    logger.info(f"Retrieved {len(problems)} problems.")

    templ_file = CUR_DIR / "template" / "README.md"
    logger.info(f"Reading template from: {templ_file.relative_to(ROOT_DIR)}")

    problem_table = gen_problem_table(problems=problems)
    readme_text = templ_file.read_text().format(problem_table=problem_table)
    readme_file = ROOT_DIR / "README.md"
    readme_file.write_text(readme_text)

    logger.info(f"Updated {readme_file.relative_to(ROOT_DIR)} successfully!")


def gen_problem_table(problems: list[Problem]) -> str:
    body = []
    for problem in problems:
        number = problem.number
        title = f"[{problem.title}]({problem.link})"
        difficulty = problem.difficulty

        # Special for GitHub: use `directory` instead of `doc_file`,
        # because GitHub will render README.md automatically,
        # but it may lead to ambiguity.
        doc = f"[📃]({problem.directory.relative_to(ROOT_DIR)})"

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


def get_problems(problems_dir: Path) -> list[Problem]:
    problems = []

    for child in problems_dir.iterdir():
        if child.is_dir():
            problem = ProblemExtractor(child).get_problem()
            if problem is not None:
                problems.append(problem)

    return problems


class ProblemExtractor:

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

    def __init__(self, problem_dir: Path):
        self.problem_dir = problem_dir
        self.doc_file = problem_dir / "README.md"

    def extract_metadata(self) -> tuple | None:
        try:
            with self.doc_file.open("r") as f:
                file_head = "".join(next(f) for _ in range(10))  # Read first 10 lines
        except Exception as e:
            logger.warning(f"Error reading {self.doc_file}: {e}")
            return None

        match = self.METADATA_PATTERN.search(file_head)
        if not match:
            logger.warning(f"Cannot extract metadata from {self.doc_file}, skipping.")
            return None

        number = int(match.group(1))
        title = match.group(2).strip()
        difficulty = match.group(3)
        topics = match.group(4).strip().split(", ")
        link = match.group(5)

        return number, title, difficulty, topics, link

    def find_solution_files(self) -> dict[str, Path]:
        solution_files = {
            self.CODE_SUFFIX_MAP[file.suffix]: file
            for file in self.problem_dir.iterdir()
            if file.is_file() and file.suffix in self.CODE_SUFFIX_MAP
        }

        # If no direct matches, try subdirectories
        # E.g., cpp/*.cpp, java/*.java.
        if not solution_files:
            solution_files = {
                self.CODE_SUFFIX_MAP[f".{subdir.name}"]: subdir
                for subdir in self.problem_dir.iterdir()
                if subdir.is_dir() and f".{subdir.name}" in self.CODE_SUFFIX_MAP
            }

        return solution_files

    def get_problem(self) -> Problem | None:
        metadata = self.extract_metadata()
        if not metadata:
            return None

        number, title, difficulty, topics, link = metadata
        solution_files = self.find_solution_files()

        return Problem(
            number=number,
            title=title,
            difficulty=difficulty,
            link=link,
            topics=topics,
            directory=self.problem_dir,
            doc_file=self.doc_file,
            solution_files=solution_files,
        )


if __name__ == "__main__":
    main()
