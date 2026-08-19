# Makefile for Go Modules examples

.PHONY: help tidy deps clean run-all test

help:
	@echo "Go Modules 示例 Makefile"
	@echo ""
	@echo "可用命令:"
	@echo "  make tidy      - 整理所有示例的依赖"
	@echo "  make deps      - 下载所有示例的依赖"
	@echo "  make clean     - 清理缓存"
	@echo "  make run-all   - 运行所有示例"
	@echo "  make test      - 验证练习和 workspace 示例"

# 整理依赖
tidy:
	@echo "整理依赖..."
	@cd 01-basic && go mod tidy
	@cd 02-dependencies && go mod tidy
	@cd 03-replace && go mod tidy
	@cd 04-versions && go mod tidy
	@cd 05-private && go mod tidy
	@echo "完成"

# 下载依赖
deps:
	@echo "下载依赖..."
	@cd 01-basic && go mod download
	@cd 02-dependencies && go mod download
	@cd 03-replace && go mod download
	@cd 04-versions && go mod download
	@cd 05-private && go mod download
	@echo "完成"

# 清理缓存
clean:
	@echo "清理缓存..."
	@go clean -modcache
	@echo "完成"

# 运行所有示例
run-all:
	@echo "=== 01-basic ==="
	@cd 01-basic && go run main.go
	@echo ""
	@echo "=== 02-dependencies ==="
	@cd 02-dependencies && go run main.go
	@echo ""
	@echo "=== 03-replace ==="
	@cd 03-replace && go run main.go
	@echo ""
	@echo "=== 04-versions ==="
	@cd 04-versions && go run main.go
	@echo ""
	@echo "=== 05-private ==="
	@cd 05-private && go run main.go
	@echo ""
	@echo "=== 06-workspace ==="
	@cd 06-workspace && go run ./app

test:
	@cd exercises && go test ./...
	@cd 06-workspace && go test ./app/... ./greeter/...
