package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/post", pkg.handlePost) // 使用 pkg.HandlePost 函数

	log.Println("Starting HTTP server on :5000...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
