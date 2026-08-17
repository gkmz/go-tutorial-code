# 第 1.9 章：错误处理

## 本章内容

本章介绍 Go 语言的错误处理机制，包括：

- error 接口和错误创建
- 基本错误处理模式
- 错误包装和检查（Go 1.13+）
- 自定义错误类型
- Panic 和 Recover
- 错误处理最佳实践

## 代码文件

| 文件 | 说明 |
|------|------|
| `01-basic-error.go` | 基础错误处理 |
| `02-error-wrapping.go` | 错误包装（%w vs %v）|
| `03-custom-error.go` | 自定义错误类型 |
| `04-panic-recover.go` | Panic 和 Recover |
| `05-pitfall-defer-error.go` | 坑 1：defer 中的错误 |
| `06-pitfall-goroutine-panic.go` | 坑 3：goroutine 中的 panic |
| `07-retry-mechanism.go` | 实战：重试机制 |
| `08-multi-error.go` | 实战：错误聚合 |
| `09-complete-example.go` | 完整示例 |

## 运行示例

```bash
# 运行基础示例
go run 01-basic-error.go

# 运行错误包装示例
go run 02-error-wrapping.go

# 运行自定义错误示例
go run 03-custom-error.go

# 运行 panic/recover 示例
go run 04-panic-recover.go

# 运行踩坑示例
go run 05-pitfall-defer-error.go
go run 06-pitfall-goroutine-panic.go

# 运行实战示例
go run 07-retry-mechanism.go
go run 08-multi-error.go
go run 09-complete-example.go
```

## 核心知识点

### 1. error 接口

```go
type error interface {
    Error() string
}
```

### 2. 创建错误

```go
// 方式 1：errors.New
err := errors.New("something went wrong")

// 方式 2：fmt.Errorf
err := fmt.Errorf("failed to open file: %s", filename)

// 方式 3：自定义错误类型
type MyError struct {
    Code int
    Msg  string
}

func (e MyError) Error() string {
    return fmt.Sprintf("error %d: %s", e.Code, e.Msg)
}
```

### 3. 错误包装（Go 1.13+）

```go
// 使用 %w 包装错误
return fmt.Errorf("failed to read config: %w", err)

// 检查错误
if errors.Is(err, os.ErrNotExist) {
    // 处理文件不存在
}

// 提取错误类型
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("Failed path:", pathErr.Path)
}
```

### 4. Panic 和 Recover

```go
// Panic
panic("something went wrong")

// Recover（只在 defer 中有效）
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

### 5. 卫语句风格

```go
func processFile(filename string) error {
    // 先处理错误
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    // 正常逻辑保持在最左侧
    // ...
    return nil
}
```

## 学习目标

完成本章学习后，你应该能够：

- 已完成 理解 Go 的错误处理理念
- 已完成 正确创建和返回错误
- 已完成 使用 %w 包装错误并检查
- 已完成 定义自定义错误类型
- 已完成 正确使用 panic 和 recover
- 已完成 避免常见的错误处理陷阱
- 已完成 实现重试和错误聚合

## 常见问题

### Q1: 什么时候使用 panic，什么时候用 error？

A: 判断标准：
- 可以恢复的错误 → 返回 error
- 程序无法继续执行 → panic
- 库代码 → 返回 error，让调用方决定
- 程序初始化失败 → 可以 panic

### Q2: %w 和 %v 有什么区别？

A: 
- `%w` 包装错误，保留原始错误，可以用 `errors.Is` 和 `errors.As` 检查
- `%v` 只是格式化错误消息，丢失原始错误信息

### Q3: 为什么 defer file.Close() 的错误也要检查？

A: Close 可能失败（如磁盘满、网络断开），忽略错误可能导致数据丢失。

### Q4: goroutine 中的 panic 能被外部 recover 吗？

A: 不能。必须在 goroutine 内部使用 defer recover。

## 练习题

1. 编写文件操作函数，正确处理 defer Close 的错误
2. 实现重试机制，支持指数退避
3. 定义 ValidationError 类型，聚合多个字段错误
4. 使用 panic/recover 实现类似 try-catch 的错误处理
5. 实现错误包装函数，添加时间戳和调用栈
6. 创建 HTTP 请求函数，根据状态码返回不同错误

---
