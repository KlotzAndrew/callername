#!/bin/bash

set -ex

mkdir -p releases
GOOS=linux GOARCH=amd64 go build -o ./releases/linux-amd64 .
tar czfv ./releases/linux-amd64.tar.gz ./releases/linux-amd64

shasum -a 256 ./releases/* > ./releases/sha256sums.txt
cat ./releases/sha256sums.txt
