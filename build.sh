#!/usr/bin/env bash

set -euo pipefail

go mod tidy
go build -o ./bin/server ./cmd/server
