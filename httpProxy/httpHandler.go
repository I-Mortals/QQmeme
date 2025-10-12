package httpProxy

import (
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
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

	// 处理不同操作系统的路径格式
	var fileDir string

	if runtime.GOOS == "windows" {
		// Windows格式: /\c\path\to\file -> c:\path\to\file
		rootPath := filePath[0:3]
		fileDir = rootPath[1:2] + ":" + filePath[3:]
	} else if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		// Unix/Mac格式: /path/to/file -> /path/to/file
		fileDir = filePath
	}

	log.Println("ServePATH:", fileDir)
	f, err := os.OpenFile(fileDir, os.O_RDONLY, 0)
	if err != nil {
		log.Printf("打开文件失败: %s, 错误: %v", fileDir, err)
		w.WriteHeader(500)
		return
	}
	// 确保文件被关闭
	defer f.Close()

	w.Header().Set("Cache-Control", "max-age=86400") // 缓存一天
	bs, err := io.ReadAll(f)                         // Copy\ReadAll
	if err != nil {
		log.Printf("读取文件失败: %s, 错误: %v", fileDir, err)
		w.WriteHeader(500)
		return
	}
	w.Write(bs)
}
