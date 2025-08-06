package httpProxy

import (
	"io"
	"log"
	"net/http"
	"os"
)

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("ServeHTTP:", r.URL)

	filePath := r.URL.Path
	rootPath := filePath[0:3]
	fileDir := rootPath[1:2] + ":" + filePath[3:]

	log.Println("ServePATH:",fileDir)
	f, err := os.OpenFile(fileDir, os.O_RDONLY, 0)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	// 确保文件被关闭
	defer f.Close()

	w.Header().Set("Cache-Control", "max-age=86400") // 缓存一天
	bs, err := io.ReadAll(f) // Copy\ReadAll
	w.Write(bs)
}
