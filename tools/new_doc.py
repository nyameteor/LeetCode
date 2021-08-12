#!/usr/bin/python3

# Thanks to https://github.com/Triple-Z/LeetCode/blob/master/utils/new_doc.py .

import re
import shutil
from tempfile import mkstemp


def main():
    number = input("Input problem number:")
    title = input("Input problem title:")
    subtitle = input("Input problem subtitle:")
    difficulty = input("Input problem difficulty(e:Easy, m:Medium, h:Hard):")
    difficulty = {
        'e': 'Easy',
        'm': 'Medium',
        'h': 'Hard',
    }.get(difficulty)

    topics_str = input('Input the problem topics (sep in comma): ')
    topics = topics_str.split(',')
    for i, topic in enumerate(topics):
        topics[i] = topic.strip()
    for topic in topics:
        topic = f"`topic`"

    link = input('Input the problem link: ')
    if not (link.startswith('http://') or link.startswith('https://')):
        print('The problem link must be start with `http://` or `https://`')
        exit(1)
    filename = f"{number}. {title}"
    template_path = "./tools/template/solution.md"
    doc_path = f"./docs/{filename}.md"
    shutil.copy(template_path, doc_path)
    sed("<NUMBER>", number, doc_path)
    sed("<TITLE>", title, doc_path)
    sed("<SUBTITLE>", subtitle, doc_path)
    sed("<DIFFICULTY>", difficulty, doc_path)
    sed("<TOPICS>", ", ".join(topics), doc_path)
    sed("<LINK>", link, doc_path)

    print(f"\nCreate doc at: {doc_path}")


# Refer: https://stackoverflow.com/a/40843600
def sed(pattern, replace, source, dest=None, count=0):
    """Reads a source file and writes the destination file.

    In each line, replaces pattern with replace.

    Args:
        pattern (str): pattern to match (can be re.pattern)
        replace (str): replacement str
        source  (str): input filename
        count (int): number of occurrences to replace
        dest (str):   destination filename, if not given, source will be over written.        
    """

    fin = open(source, 'r')
    num_replaced = count

    if dest:
        fout = open(dest, 'w')
    else:
        fd, name = mkstemp()
        fout = open(name, 'w')

    for line in fin:
        out = re.sub(pattern, replace, line)
        fout.write(out)

        if out != line:
            num_replaced += 1
        if count and num_replaced > count:
            break
    try:
        fout.writelines(fin.readlines())
    except Exception as E:
        raise E

    fin.close()
    fout.close()

    if not dest:
        shutil.move(name, source)


if __name__ == "__main__":
    main()
