#!/usr/bin/env bash

# shellcheck disable=SC2155

set -o errexit \
    -o nounset

readonly CUR_DIR=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
readonly OUT_DIR="${CUR_DIR}/../out"

run_c() {
    local src_file="${1}"
    local out_file="solution.out"

    gcc -g \
        "${src_file}" \
        --output "${OUT_DIR}/${out_file}" &&
        "${OUT_DIR}/${out_file}"
}

run_cpp() {
    local src_file="${1}"
    local out_file="solution.out"

    g++ -g \
        -std=c++17 "${src_file}" \
        --output "${OUT_DIR}/${out_file}" &&
        "${OUT_DIR}/${out_file}"
}

run_java() {
    local src_file="${1}"
    local out_file="Solution"

    javac -d "${OUT_DIR}" "${src_file}" &&
        cd "${OUT_DIR}" &&
        java -ea "${out_file}"
}

run_go() {
    local src_file="${1}"

    go run "${src_file}"
}

run_rust() {
    local src_file="${1}"
    local out_file="solution.out"

    rustc \
        --edition 2024 \
        -o "${OUT_DIR}/${out_file}" \
        "${src_file}" &&
        "${OUT_DIR}/${out_file}"
}

main() {
    local src_file="${1}"

    mkdir -p "${OUT_DIR}"

    local filename=$(basename -- "${src_file}")
    local extension="${filename##*.}"

    case "${extension}" in
    c)
        run_c "${src_file}"
        ;;
    cpp)
        run_cpp "${src_file}"
        ;;
    java)
        run_java "${src_file}"
        ;;
    go)
        run_go "${src_file}"
        ;;
    rs)
        run_rust "${src_file}"
        ;;
    *)
        echo "Not a valid source file."
        ;;
    esac
}

main "${1}"
