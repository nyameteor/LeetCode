#!/usr/bin/env bash

set -ex

readonly OUT_DIR="./out"
readonly CLASS_NAME="Solution"

main() {
  local java_file="${1}"

  javac -d "${OUT_DIR}" "${java_file}" && cd "${OUT_DIR}" && java "${CLASS_NAME}"
}

main "${1}"
