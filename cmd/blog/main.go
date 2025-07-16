package main

import (
	"github.com/limyel/mywheels/internal/blog/handler"
	"log"
	"net/http"
)

func main() {
	// 创建一个多路复用器
	mux := http.NewServeMux()

	// 创建文件服务器，提供静态文件服务。http.Dir 指定静态文件的根目录
	files := http.FileServer(http.Dir("./ui/static/"))
	// 将 /static/ 前缀的请求交给文件服务器处理
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", handler.HomePage)

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
