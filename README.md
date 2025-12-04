# Sonata

[![Go Reference](https://pkg.go.dev/badge/github.com/sonata-labs/sonata.svg)](https://pkg.go.dev/github.com/sonata-labs/sonata)
[![Go Report Card](https://goreportcard.com/badge/github.com/sonata-labs/sonata?style=flat&v=1)](https://goreportcard.com/report/github.com/sonata-labs/sonata?style=flat&v=1)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The protocol for open audio distribution.

## Single Node Development Setup

```bash
# install sonata cli
go install github.com/sonata-labs/sonata/cmd/sonata@latest
go install github.com/air-verse/air@latest

# initialize a local node to a local dir
sonata init --home=./tmp/sonata-dev

# start dev node using local dir
air -- run --home ./tmp/sonata-dev
```

## Code Generation

```bash 
# install buf
go install github.com/bufbuild/buf/cmd/buf@latest

# generate code
buf generate

# tidy up
go mod tidy
```

To clean up the generated code, run:
```bash
rm -rf gen
```

## VSCode/Cursor Setup

VSCode/Cursor is the recommended code editor for this project. The extensions below are recommended:
- [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [Buf](https://marketplace.visualstudio.com/items?itemName=bufbuild.vscode-buf)
- [YAML](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml)
- [TOML](https://marketplace.visualstudio.com/items?itemName=tamasfe.even-better-toml)
- [Github Actions](https://marketplace.visualstudio.com/items?itemName=github.vscode-github-actions)

## Obsidian Setup

Obsidian is the recommended tool for managing documentation and notes. Open this project as a vault and navigate to the `docs` folder to view the documentation.