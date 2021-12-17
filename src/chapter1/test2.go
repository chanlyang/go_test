package main

import (
	"net/http"
)

//实现一个简单的http服务器
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
