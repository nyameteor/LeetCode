#!/usr/bin/env bash

# shellcheck disable=SC2155

set -o errexit \
    -o nounset \
    -o xtrace

readonly CUR_DIR=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
readonly OUT_DIR="${CUR_DIR}/../out"
readonly JAVA_OUT="Solution"
readonly CPP_OUT="solution.out"

run_cpp() {
    local src_file="${1}"

    g++ -g -std=c++17 "${src_file}" --output "${OUT_DIR}/${CPP_OUT}" &&
        "${OUT_DIR}/${CPP_OUT}"
}

run_java() {
    local src_file="${1}"

    javac -d "${OUT_DIR}" "${src_file}" &&
        cd "${OUT_DIR}" &&
        java -ea "${JAVA_OUT}"
}

main() {
    local src_file="${1}"

    # Refer: https://stackoverflow.com/a/965069
    local filename=$(basename -- "${src_file}")
    local extension="${filename##*.}"

    case "${extension}" in
    cpp)
        run_cpp "${src_file}"
        ;;
    java)
        run_java "${src_file}"
        ;;
    *)
        echo "Not a valid source file."
        ;;
    esac
}

main "${1}"
