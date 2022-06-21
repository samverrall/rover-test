#!/usr/bin/env bash

set -e

golangci-lint --enable gosec,misspell run ./...

go test -v --cover --race --count=1 ./...