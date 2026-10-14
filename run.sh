#!/usr/bin/env bash

set -o errexit \
    -o nounset

cur_dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
exec "${cur_dir}/tools/run_code.sh" "$@"
