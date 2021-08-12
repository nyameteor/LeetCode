#!/usr/bin/env bash

# shellcheck disable=SC2155

set -o errexit
set -o xtrace

readonly CUR_DIR="$(cd -P -- "$(dirname -- "$0")" && pwd -P)"
readonly OUT_DIR="${CUR_DIR}/../out"
readonly CLASS_NAME="Solution"

main() {
  local src_file="${1}"

  javac -d "${OUT_DIR}" "${src_file}" && cd "${OUT_DIR}" && java "${CLASS_NAME}"
}

main "${1}"
