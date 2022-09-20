import re
import sys
import shutil
import subprocess
import os
from tempfile import mkstemp
from pathlib import Path


def gen_from_template(src: Path, dst: Path, pattern_map: dict):
    shutil.copy(src=src, dst=dst)
    for k, v in pattern_map.items():
        sed(pattern=k, replace=v, source=dst)


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


def is_tracked(filepath):
    """
    Return if file is tracked by git.
    """
    if not os.path.exists(filepath):
        print("File not exist!")
        return False
    # Check if file is tracked
    output = subprocess.run(
        ["git", "ls-files", f"{filepath}"], capture_output=True, text=True)
    if output.stdout == '':
        return False
    else:
        return True
