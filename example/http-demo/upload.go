package http_demo

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 限制上传大小 10GB
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	// 多文件处理
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		safePath := filepath.Join("uploads", filepath.Base(fileHeader.Filename))

		if strings.Contains(safePath, "..") {
			http.Error(w, "Invalid file path", http.StatusBadRequest)
			return
		}

		dst, err := os.Create(safePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
