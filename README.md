# calculator-test

Simple command-line calculator written in Go.

## Why no committed `.exe`?
Some code-hosting/review systems reject or flag binary files in pull requests. To avoid that, this repository now keeps only source code and a repeatable build script.

## Features
- Supports `+`, `-`, `*`, `/`
- Supports parentheses, whitespace, and unary signs (for example `-3 + (+2)`)
- Interactive REPL mode (run without arguments)
- One-shot mode (pass expression as command-line arguments)

## Run locally (Linux/macOS)
```bash
go run .
go run . "(2 + 3) * 4 / 5"
```

## Build a self-contained Windows `.exe`
```bash
./scripts/build-windows.sh
```

This generates:
- `dist/calculator.exe`

The build uses `CGO_ENABLED=0`, so the executable is self-contained and does not need cgo runtime DLLs.

## Example
```bash
./dist/calculator.exe "12 / (2 + 1)"
```
Output:
```text
4
```
