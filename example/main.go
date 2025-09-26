package main

import (
	http_demo "example/http-demo"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	// 包裹下
	http.Handle("/", http_demo.TimeCountMiddleware(handler))
	http.ListenAndServe(":3333", nil)
}
