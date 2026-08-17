# Context 与 gRPC 调用边界

本目录记录 Context 在 gRPC 调用链中的关键用法。示例使用 `pb` 表示 protobuf 编译器生成的客户端和服务端代码，重点是 Context、metadata、截止时间和错误状态码，而不是重复展开 protobuf 生成物。

## 调用链

```text
HTTP handler context
  -> service context
    -> gRPC client context
      -> gRPC server handler context
        -> database context
```

每一层都必须继续传递收到的 Context。重新创建 `context.Background()` 会切断取消和截止时间传播。

## 客户端调用

```go
func getUser(ctx context.Context, client pb.UserServiceClient, id int64) (*pb.User, error) {
	callCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return client.GetUser(callCtx, &pb.GetUserRequest{Id: id})
}
```

客户端连接通常在进程启动时创建并复用，不应为每次 RPC 都建立新连接。每个 RPC 仍然应该有明确的截止时间。

## 服务端处理

```go
func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := queryUser(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, status.Error(codes.Canceled, "request canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
		return nil, status.Error(codes.Internal, "query user failed")
	}
	return user, nil
}
```

服务端 handler 必须把 Context 传给数据库、HTTP 客户端或下游 RPC。取消只会关闭 `Done`，不会强制终止不支持 Context 的函数。

## metadata

Context value 只在当前进程内有效。需要跨 gRPC 边界传递 request ID、租户 ID 或认证信息时，应使用 metadata，并在服务端进行校验。

```go
outgoing := metadata.NewOutgoingContext(ctx, metadata.Pairs("x-request-id", requestID))
resp, err := client.GetUser(outgoing, req)
```

服务端使用 `metadata.FromIncomingContext(ctx)` 读取 metadata。敏感信息必须配合 TLS 和认证机制，不能把 Context 当成安全边界。

## 错误判断

客户端使用 `status.Code(err)` 判断 gRPC 错误，不比较错误字符串：

```go
if status.Code(err) == codes.DeadlineExceeded {
	// 记录超时，并根据业务决定是否重试。
}
```

## 验证重点

- 父请求取消后，服务端 handler 是否退出；
- 下游调用是否继承同一个 Context；
- 子调用超时是否不会超过父请求截止时间；
- metadata 是否正确传递且没有泄露敏感信息；
- 客户端是否按状态码处理取消、超时和参数错误。
