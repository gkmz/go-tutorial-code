package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// 使用标准库 slog 输出结构化日志，日志端点应放在管理网络中。
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("server started", "addr", ":8085")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })
	logger.Error("server stopped", "err", http.ListenAndServe(":8085", nil))
}
