package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// pprof 只应暴露在受保护的管理网络中，不要直接公开到公网。
	log.Fatal(http.ListenAndServe("127.0.0.1:6060", nil))
}
