#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
go run ./hack/tools/geps/main.go
