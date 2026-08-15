#!/usr/bin/env bash

set -euo pipefail

# 先格式化全部示例，再验证每个独立模块和单文件示例。
repo_dir=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
cd "$repo_dir"

gofmt -w $(find . -name '*.go' -type f -not -path './.git/*')

for mod in $(find . -name go.mod -not -path './.git/*' -print); do
	mod_dir=$(dirname "$mod")
	printf 'testing module: %s\n' "$mod_dir"
	(cd "$mod_dir" && go test ./...)
done

for file in $(find . -name '*.go' -not -name '*_test.go' -not -path './.git/*' \
	-not -path './12-testing/*' -not -path './13-modules/*' -not -path './15-project-example/*' -not -path './16-cli-urfave/*' \
	-not -path './18-advanced/*' -not -path './19-web/*' -print); do
	printf 'running example: %s\n' "$file"
	go run "$file" >/dev/null
done

printf 'verification passed\n'
