# Thanks to https://github.com/Triple-Z/LeetCode/blob/master/Makefile .

.DEFAULT_GOAL:=help

# set default shell
SHELL=/bin/bash -o pipefail -o errexit

.PHONY: help
help:	## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: doc
doc:	## Create a new problem folder and doc file (interactive mode)
	@python3 tools/gen_doc.py

.PHONY: readme
readme:	## Update README.md
	@python3 tools/gen_readme.py
	@echo "Updated README.md successfully!"