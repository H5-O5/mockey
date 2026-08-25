#!/usr/bin/env bash
# 2x2 evidence matrix on linux/amd64:
#   axis 1: mockey = upstream v1.3.0  vs  H5-O5 fork (this worktree)
#   axis 2: gcflags = "-N -l" (status quo)  vs  "-l" (MR !739 drops -N)
# Target: a SHORT embedded method, which is what -N's real job (preserving
# function SIZE) protects.
set -u
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
run() {
  local label="$1" mount="$2" flags="$3"
  echo "=== ${label} | gcflags=all=${flags} ==="
  docker run --rm --platform linux/amd64 \
    -v "$ROOT":/mockey -v "${mount}":/w -w /w \
    -e GOFLAGS=-mod=mod -e MOCKEY_CHECK_GCFLAGS=false \
    -e GOPROXY=https://goproxy.cn,direct \
    golang:1.23-bookworm \
    go test -run TestShortE2E -gcflags=all="${flags}" -count=1 . 2>&1 \
    | grep -E "too short to patch|return args not match|^ok|^--- (PASS|FAIL)|^PASS" | head -3
  echo ""
}
run "UPSTREAM v1.3.0" "$ROOT/.ab/upstream" "-N -l"
run "UPSTREAM v1.3.0" "$ROOT/.ab/upstream" "-l"
run "FORK (worktree)" "$ROOT/.ab/repro"    "-N -l"
run "FORK (worktree)" "$ROOT/.ab/repro"    "-l"
