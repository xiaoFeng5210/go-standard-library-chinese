package http_demo

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		saveFilePath := filepath.Join("assets", "uploads", "upload_test.jpg")
		var saveFile *os.File
		fmt.Printf("saveFilePath: %s\n", saveFilePath)
		saveFile, err = os.Create(saveFilePath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		_, err = io.Copy(saveFile, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("保存文件失败"))
			return
		}
		saveFile.Close()
		w.Write([]byte("文件上传成功"))
	})

	fmt.Println("Server is running on port 3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
