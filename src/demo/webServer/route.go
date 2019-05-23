package webServer

import (
	"log"
	"net/http"
)

// web服务 路由
func route() {
	//web服务 中间件
	middleware()

	//web服务 路由配置
	http.HandleFunc("/", handler)
	// http.HandleFunc("/count", counter)

	// 启动服务
	// http.ListenAndServe("localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
