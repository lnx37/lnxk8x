#!/bin/bash

set -e
set -o pipefail
set -u
set -x

cd "$(dirname "$0")"

date

cp -a ../lnxk8s/env.sh .

python sh2go.py

# go build .
go build -ldflags="-s -w" .

date
