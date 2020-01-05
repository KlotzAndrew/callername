#!/bin/bash

set -ex

mkdir -p releases
GOOS=linux GOARCH=amd64 go build -o ./releases/linux-amd64 .
tar czfv ./releases/linux-amd64.tar.gz ./releases/linux-amd64

shasum -a 512 ./releases/* > ./releases/sha512sum.txt
cat ./releases/sha512sum.txt
