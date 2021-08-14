import re
import sys
import shutil
import subprocess
import os
from tempfile import mkstemp


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


# Refer: https://stackoverflow.com/a/21286092
def insert_lines_after_matching(filepath, content, lines):
    """
    insert lines after matched specific line
    """
    if not os.path.exists(filepath):
        print("File not exist!")
        return False
    with open(filepath, "r")as in_file:
        buf = in_file.readlines()

    with open(filepath, "w")as out_file:
        for line in buf:
            if line == content:
                for l in lines:
                    line = line + l
            out_file.write(line)

    return True
