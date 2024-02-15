#!/usr/bin/python3

# Thanks to https://github.com/Triple-Z/LeetCode/blob/master/utils/new_doc.py .

import re
from pathlib import Path

from utils import gen_from_template


ROOT_DIR: Path = Path(__file__).parent.parent
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
    link, problem_name = parse_link(link)

    problem_dir = PROBLEMS_DIR / problem_name
    Path.mkdir(problem_dir, exist_ok=True)

    # Store metadata to doc
    doc_file = (problem_dir / "README.md").resolve()
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


def parse_link(link: str):
    """
    Accept inputs:
        https://leetcode.com/problems/<problem-name>/
        https://leetcode.com/problems/<problem-name>/description/
    """
    if not (link.startswith("http://") or link.startswith("https://")):
        print("Link must be start with `http://` or `https://`")
        exit(1)

    link_pattern = re.compile(
        r'^(http|https)(://leetcode.com/problems/)(.+?)/(|.+?/)$',
        re.ASCII
    )
    match = link_pattern.search(link)
    if match is None:
        print("Failed to match link")
        exit(1)
    problem_name = match.group(3)
    link = link_pattern.sub(r"\g<1>\g<2>\g<3>/", link)

    return link, problem_name


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print('\nProcess terminated by keyboard interrupt.')
