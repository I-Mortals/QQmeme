package utils

import (
	"os"
	"strings"
)

type ImageUtils struct{}

func NewImageUtils() *ImageUtils {
	return &ImageUtils{}
}

var imageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}

// IsImageFile 检查文件是否为支持的图片格式
func (i *ImageUtils) IsImageFile(fileName string) bool {
	if fileName == "" {
		return false
	}

	fileName = strings.ToLower(fileName)

	for _, ext := range imageExtensions {
		if strings.HasSuffix(fileName, ext) {
			return true
		}
	}
	return false
}

// GetImages 获取指定目录下的所有图片文件
func (i *ImageUtils) GetImages(path string) []string {
	if path == "" {
		return nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var images []string
	for _, entry := range entries {
		if !entry.IsDir() && i.IsImageFile(entry.Name()) {
			images = append(images, entry.Name())
		}
	}
	return images
}
