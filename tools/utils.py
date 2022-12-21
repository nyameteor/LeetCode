import re
import shutil
import subprocess
import os
from tempfile import mkstemp
from pathlib import Path


def gen_from_template(src: Path, dst: Path, pattern_map: dict):
    shutil.copy(src=src, dst=dst)
    for k, v in pattern_map.items():
        sed(pattern=k, replace=v, source=dst)


def sed(pattern, replace, source, dest=None, count=1):
    """Reads a source file and writes the destination file.

    In each line, replaces pattern with replace.

    Args:
        pattern (str): pattern to match (can be re.pattern)
        replace (str): replacement str
        source  (str): input filename
        count (int): number of occurrences to replace
        dest (str): destination filename, if not given, source will be over written.

    Ref: 
        https://stackoverflow.com/a/40843600
    """
    overwrite = False
    if not dest:
        _, dest = mkstemp()
        overwrite = True

    num_replaced = 0
    with open(source, 'r') as fin, open(dest, 'w') as fout:
        for line in fin:
            out = re.sub(pattern, replace, line)
            fout.write(out)

            if out != line:
                num_replaced += 1
            if num_replaced >= count:
                break
        fout.writelines(fin.readlines())

    if overwrite:
        shutil.move(dest, source)


def is_tracked(filepath):
    """
    Return if file is tracked by git.
    """
    if not os.path.exists(filepath):
        print("File not exist!")
        return False
    output = subprocess.run(
        ["git", "ls-files", f"{filepath}"],
        capture_output=True,
        text=True
    )
    return output.stdout != ''
