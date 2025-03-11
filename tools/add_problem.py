#!/usr/bin/python3

import re
from pathlib import Path

CUR_DIR = Path(__file__).parent
ROOT_DIR = CUR_DIR.parent


def prompt_input(prompt, transform=lambda x: x):
    """
    Prompt for input and apply transformation. If transformation fails, ask again.
    """
    while True:
        value = input(prompt).strip()
        try:
            result = transform(value)
            if result is not None:
                return result
        except Exception as e:
            print(f"Error: {e}")
        print("Invalid input. Please try again.")


def main():
    print("\nAdd a new LeetCode problem. Press Ctrl+C to exit.\n")

    number = prompt_input("Number: ", lambda x: int(x))

    title = prompt_input("Title: ")

    difficulty = prompt_input(
        "Difficulty (e: Easy, m: Medium, h: Hard): ",
        lambda x: ({"e": "Easy", "m": "Medium", "h": "Hard"}[x.lower()]),
    )

    topics = prompt_input(
        "Topics (comma-separated): ",
        lambda x: ", ".join([topic.strip() for topic in x.split(",")]) if x else None,
    )

    def parse_link(link: str):
        link_pattern = re.compile(
            r"^(http|https)://leetcode\.com/problems/(.+?)/", re.ASCII
        )
        match = link_pattern.match(link)
        if match:
            return link, match.group(2)  # Extract normalized problem name
        raise ValueError("Invalid LeetCode problem link.")

    link, problem_name = prompt_input("Link: ", parse_link)

    problems_dir = ROOT_DIR / "problems"
    problem_dir = problems_dir / problem_name
    problem_dir.mkdir(exist_ok=True)

    templ_file = CUR_DIR / "template" / "solution.md"
    doc_file = (problem_dir / "README.md").resolve()
    doc_text = templ_file.read_text().format(
        number=number, title=title, difficulty=difficulty, topics=topics, link=link
    )
    doc_file.write_text(doc_text)

    print(f"\nCreated problem documentation at: {doc_file.relative_to(ROOT_DIR)}")


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nProcess terminated by keyboard interrupt.")
