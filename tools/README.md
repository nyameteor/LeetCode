# Tools

Go to project folder.

```shell
cd LeetCode
```

Input problem metadata interactively, it will create a new folder `./problems/*` and store metadata in `./problems/*/README.md`.

```shell
make doc
```

Update the table items in `./README.md`.

```shell
make readme
```

Compile and run source code.

```shell
./tools/run_code.sh problems/*/answer.cpp
./tools/run_code.sh problems/*/answer.java
```
