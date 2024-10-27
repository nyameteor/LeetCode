# Thanks to https://github.com/Triple-Z/LeetCode/blob/master/Makefile .

.DEFAULT_GOAL:=help

# set default shell
SHELL:=/bin/bash -o pipefail -o errexit

.PHONY: help
help:	## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: add-problem
add-problem:	## Add a new problem
	@python3 tools/add_problem.py
	
.PHONY: add
add: add-problem	## Alias for add-problem

.PHONY: doc
doc:	## Update documentation
	@python3 tools/update_doc.py