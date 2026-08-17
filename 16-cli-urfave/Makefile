# Makefile for CLI tool examples

.PHONY: help build run clean

help:
	@echo "CLI 工具示例 Makefile"
	@echo ""
	@echo "可用命令:"
	@echo "  make build     - 构建所有示例"
	@echo "  make run       - 运行所有示例"
	@echo "  make clean     - 清理构建文件"

# 构建所有示例
build:
	@echo "构建示例..."
	@cd 01-hello && go build -o hello main.go
	@cd 02-flags && go build -o flags main.go
	@cd 03-commands && go build -o commands main.go
	@cd 04-subcommands && go build -o subcommands main.go
	@cd 05-complete && go build -o complete main.go
	@echo "构建完成"

# 运行所有示例
run:
	@echo "=== 01-hello ==="
	@cd 01-hello && go run main.go
	@echo ""
	@echo "=== 02-flags ==="
	@cd 02-flags && go run main.go -n "老墨"
	@echo ""
	@echo "=== 03-commands ==="
	@cd 03-commands && go run main.go hello -n "老墨"
	@echo ""
	@echo "=== 04-subcommands ==="
	@cd 04-subcommands && go run main.go hello -n "老墨" weather
	@echo ""
	@echo "=== 05-complete ==="
	@cd 05-complete && go run main.go -v hello -n "老墨"

# 清理构建文件
clean:
	@echo "清理构建文件..."
	@rm -f 01-hello/hello
	@rm -f 02-flags/flags
	@rm -f 03-commands/commands
	@rm -f 04-subcommands/subcommands
	@rm -f 05-complete/complete
	@echo "清理完成"
