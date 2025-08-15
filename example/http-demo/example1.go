package http_demo

import (
	"fmt"
	"net/http"
)

func Example1Server() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
	fmt.Println("Server is running on port 3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
