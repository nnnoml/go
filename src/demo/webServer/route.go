package webServer

import (
	"log"
	"net/http"
)

// web服务 路由
func route() {
	//web服务 中间件
	middleware()
	//首页
	http.HandleFunc("/", handler)
	// 表单界面
	http.HandleFunc("/login", loginHandle)

	// 表单界面
	http.HandleFunc("/upload", uploadHandle)

	// http.HandleFunc("/count", counter)

	// 启动服务
	// http.ListenAndServe("localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
