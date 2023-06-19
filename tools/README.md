# Tools

Go to project folder.

```sh
cd LeetCode
```

To create a problem document file `./problems/<problem-name>/README.md` (requires to input problem metadata interactively):

```sh
make doc
```

To update the problem tables in `./README.md`:

```sh
make readme
```

To compile and run a problem solution:

```sh
./tools/run_code.sh "problems/<problem-name>/answer.cpp"
./tools/run_code.sh "problems/<problem-name>/answer.java"
```
