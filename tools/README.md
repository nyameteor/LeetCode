# Tools

First, go to the project directory:

```sh
cd <path/to/this-repository>
```

## Add a New Problem

To add a new problem:

```sh
make add-problem
# or
make add
```

This will prompt you to input problem metadata interactively, and creates a document at `./problems/<problem-name>/README.md`.

After that, you can create a solution for the problem. The solution file can be named anything you like, for example: `./problems/<problem-name>/<solution-name>.cpp`

## Update Documentation

To update the table of contents of problems in the documentation:

```sh
make doc
```

This updates `./README.md` with the latest problem entries.

## Compile and Run a Solution

To compile and run a specific solution:

```sh
./tools/run_code.sh "problems/<problem-name>/<solution-name>.cpp"
./tools/run_code.sh "problems/<problem-name>/<solution-name>.java"
```
