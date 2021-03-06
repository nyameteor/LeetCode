#!/usr/bin/python3

# Thanks to https://github.com/Triple-Z/LeetCode/blob/master/utils/new_doc.py .

import shutil
from utils import sed

from pathlib import Path


def main():
    number = input("Input problem number:")
    title = input("Input problem title:")
    difficulty = input("Input problem difficulty(e:Easy, m:Medium, h:Hard):")
    difficulty = {
        'e': 'Easy',
        'm': 'Medium',
        'h': 'Hard',
    }.get(difficulty)

    cwd = Path(__file__).parents[0]
    root_dir = cwd / '..'

    topics_str = input('Input the problem topics (sep in comma): ')
    topics = topics_str.split(',')
    for i, topic in enumerate(topics):
        topics[i] = topic.strip()

    link = input('Input the problem link: ')
    if not (link.startswith('http://') or link.startswith('https://')):
        print('The problem link must be start with `http://` or `https://`')
        exit(1)
    filename = f"{number}. {title}"
    template_path = root_dir / 'tools' / 'template' / 'solution.md'
    doc_path = root_dir / 'docs' / f"{filename}.md"
    shutil.copy(template_path, doc_path)
    sed("<NUMBER>", number, doc_path)
    sed("<TITLE>", title, doc_path)
    sed("<DIFFICULTY>", difficulty, doc_path)
    sed("<TOPICS>", ", ".join(topics), doc_path)
    sed("<LINK>", link, doc_path)

    print(f"\nCreate doc at: {doc_path}")


if __name__ == "__main__":
    main()
