#!/usr/bin/python3

# Thanks to https://github.com/Triple-Z/LeetCode/blob/master/utils/new_doc.py .

import shutil
from utils import sed, gen_from_template
import re

from pathlib import Path

CUR_DIR: Path = Path(__file__).parents[0]
ROOT_DIR: Path = CUR_DIR / '..'
PROBLEMS_DIR: Path = ROOT_DIR / 'problems'

TEMPL_FILE: Path = ROOT_DIR / 'tools' / 'template' / 'solution.md'


def main():
    number = input("Number:")
    title = input("Title:")
    difficulty = input("Difficulty (e:Easy, m:Medium, h:Hard):")
    difficulty = {
        'e': 'Easy',
        'm': 'Medium',
        'h': 'Hard',
    }.get(difficulty)

    topic = input("Topics (sep in comma): ")
    topics = topic.split(',')
    for i, topic in enumerate(topics):
        topics[i] = topic.strip()
    topic = ", ".join(topics)

    link = input("Link:")
    if not (link.startswith("http://") or link.startswith("https://")):
        print("Link must be start with `http://` or `https://`")
        exit(1)

    name_pattern = re.compile(
        r'(?:http|https):\/\/leetcode.com\/problems\/(.*)\/',
        re.ASCII
    )
    problem_name = name_pattern.search(link).group(1)

    problem_dir = PROBLEMS_DIR / problem_name
    Path.mkdir(problem_dir, exist_ok=True)

    # Store metadata to doc
    doc_file = problem_dir / "README.md"
    gen_from_template(
        src=TEMPL_FILE,
        dst=doc_file,
        pattern_map={
            "<NUMBER>": number,
            "<TITLE>": title,
            "<DIFFICULTY>": difficulty,
            "<TOPICS>": topic,
            "<LINK>": link
        }
    )

    print(f"\nCreate doc at: {doc_file}")


if __name__ == "__main__":
    main()
