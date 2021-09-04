# Tools

Helpful tools.

The following examples are all run in the project root folder.

Display help for all targets described in Makefile.

```shell
make
# or
make help
```

Create new doc(solution note) file based on template. Input problem information interactively, it will generate doc file in `docs` folder.

```shell
make doc
```

Update project root `README.md` file based on template. It will use regex to match solutions and docs file, and generate links to them.

```shell
make readme
```

Compile and run from solution source file.

```shell
./tools/run-code.sh code/java/1.java
./tools/run-code.sh code/cpp/1.cpp
...
```
