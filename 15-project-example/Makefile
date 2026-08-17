# Makefile
.PHONY: build run test clean install

# 默认目标
all: build

# 构建
build:
	go build -o bin/calc ./cmd/calc

# 运行
# 使用方式: make run ARGS="5 + 3"
# 或者: make run ARGS="5 '*' 3"
run:
	go run cmd/calc/main.go $(ARGS)

# 运行示例
run-example:
	@echo "Addition example:"
	@go run cmd/calc/main.go 5 + 3
	@echo "\nSubtraction example:"
	@go run cmd/calc/main.go 10 - 2
	@echo "\nMultiplication example:"
	@go run cmd/calc/main.go 5 '*' 3
	@echo "\nDivision example:"
	@go run cmd/calc/main.go 10 / 2

# 测试
test:
	go test -v ./...

# 清理
clean:
	rm -rf bin/
	go clean

# 安装到 GOPATH
install:
	go install

# 交叉编译
build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/calc-linux ./cmd/calc
	GOOS=windows GOARCH=amd64 go build -o bin/calc.exe ./cmd/calc
	GOOS=darwin GOARCH=amd64 go build -o bin/calc-mac ./cmd/calc

# 格式化代码
fmt:
	go fmt ./...

# 代码检查
lint:
	golangci-lint run --timeout 5m
