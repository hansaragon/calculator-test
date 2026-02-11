#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT_DIR="${ROOT_DIR}/dist"
OUT_FILE="${OUT_DIR}/calculator.exe"

mkdir -p "${OUT_DIR}"

printf 'Building self-contained Windows executable...\n'
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
  go build -trimpath -ldflags='-s -w' -o "${OUT_FILE}" "${ROOT_DIR}"

python3 - <<'PY'
from pathlib import Path
p = Path('dist/calculator.exe')
if not p.exists():
    raise SystemExit('Build failed: dist/calculator.exe not found')
print(f'Created: {p} ({p.stat().st_size} bytes)')
print(f'Header: {p.read_bytes()[:2]!r}')
PY
