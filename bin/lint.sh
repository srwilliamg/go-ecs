#!/bin/bash
# Get project root directory

ROOT_DIR="$(
    cd -- "$(dirname "$0")" >/dev/null 2>&1
    cd ..
    pwd -P
)"

which golangci-lint
if [ "$?" == "0" ]; then
    golangci-lint run

    exit $?
fi

# Validate Docker is installed

which docker >/dev/null

if [ "$?" != "0" ]; then
    echo " - [ ERROR ] Docker is NOT installed! Install it following the instructions in: https://docs.docker.com/desktop/mac/install/"

    exit 1
fi

# Run lint

docker run --rm -v "${ROOT_DIR}:/app:ro" --workdir="/app" golangci/golangci-lint:v2.2.2 golangci-lint run
