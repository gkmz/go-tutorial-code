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

while IFS= read -r -d '' file; do
	# 只有 main 包才能作为单文件程序运行，库包交给上面的 go test 验证。
	if ! grep -qE '^package[[:space:]]+main([[:space:]]|$)' "$file"; then
		continue
	fi
	printf 'running example: %s\n' "$file"
	go run "$file" >/dev/null
done < <(find . -name '*.go' -not -name '*_test.go' -not -path './.git/*' \
	-not -path '*/exercises/*' \
	-not -path './12-testing/*' -not -path './13-modules/*' -not -path './15-project-example/*' -not -path './16-cli-urfave/*' \
	-not -path './18-advanced/*' -not -path './19-web/*' -print0)

printf 'verification passed\n'
