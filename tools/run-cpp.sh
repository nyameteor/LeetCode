#!/usr/bin/env bash

# shellcheck disable=SC2155

set -o errexit
set -o xtrace

readonly CUR_DIR="$(cd -P -- "$(dirname -- "$0")" && pwd -P)"
readonly OUT_DIR="${CUR_DIR}/../out"
readonly OUT_FILE="solution.out"

main() {
  local src_file="${1}"

  g++ -std=c++14 "${src_file}" --output "${OUT_DIR}/${OUT_FILE}" && cd "${OUT_DIR}" && "${OUT_DIR}/${OUT_FILE}"
}

main "${1}"
