package utils

import (
	"os"
	"path/filepath"
)

type FileUtils struct{}

func NewFileUtils() *FileUtils {
	return &FileUtils{}
}

// GetDirs 获取指定目录下的所有子目录
func (f *FileUtils) GetDirs(path string) []string {
	if path == "" {
		return nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs
}

// GetFiles 获取指定目录下的所有文件
func (f *FileUtils) GetFiles(path string) []string {
	if path == "" {
		return nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files
}

// PathExists 检查路径是否存在
func (f *FileUtils) PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir 检查路径是否为目录
func (f *FileUtils) IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func (f *FileUtils) IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func (f *FileUtils) JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}

func (f *FileUtils) GetFileExt(fileName string) string {
	return filepath.Ext(fileName)
}
