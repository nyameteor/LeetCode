# Tools

## Add a New Problem

To add a new problem:

```sh
python3 tools/add_problem.py
```

This will prompt you to enter metadata (name, title, link, etc.) and creates a note:

```
problems/<problem-name>/README.md
```

You can then add your solution, for example:

```
problems/<problem-name>/<solution-name>.cpp
```

## Update Documentation

To regenerate the problems table in `README.md`:

```sh
python3 tools/update_doc.py
```

This scans all problems and updates the table.

## Run a Solution

To compile and run a solution (currently supports C, C++, and Java)):

```sh
tools/run_code.sh "problems/<problem-name>/<solution-name>.cpp"
```
