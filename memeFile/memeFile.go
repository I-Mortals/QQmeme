package memeFile

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//  MemeFile struct
type MemeFile struct {
	ctx context.Context
}

func NewMemeFile() *MemeFile {
	return &MemeFile{}
}

// 从顶层传入app的ctx。将ctx赋值到MemeFile的ctx中
// 因此，我们可以调用App运行时方法
func (m *MemeFile) SetContext(ctx context.Context) {
	m.ctx = ctx
}


func (m *MemeFile) Hello() string {
	fmt.Println("hello")
	return "hello"
}



var (
	// 存放图片链接的数据管道
	imageUrls []string
)


// 返回打开选择的文件路径
func (m *MemeFile) OpenFileDlg() []string {
	opts := runtime.OpenDialogOptions{}
	filePath, err := runtime.OpenMultipleFilesDialog(m.ctx, opts)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println(filePath)
	return filePath
}


// 请选择meme根目录
func (m *MemeFile) SelectRootDir() string {
	path, err := runtime.OpenDirectoryDialog(m.ctx, runtime.OpenDialogOptions{
		Title: "请选择meme根目录",
	})
	if err != nil {
		panic("系统出错，请重试")
	}

	return path
}

// 传入一个目录并返回改目录下的文件夹
func (m *MemeFile) GetDirs(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Println("读取目录失败:", err.Error())
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

// 传入一个目录并返回改目录下的图片
func (m *MemeFile) GetImages(path string) []string {
	entries, err := os.ReadDir(path)
	if err!= nil {
		log.Println("读取目录失败:", err.Error())
		return nil
	}

	var images []string
	for _, entry := range entries {
		if !entry.IsDir() {
			// 判断是否为图片文件
			fileName := entry.Name()
			if isImageFile(fileName) {
				images = append(images, entry.Name())
			}
		}
	}
	return images
}

// 判断文件是否为图片
func isImageFile(fileName string) bool {
	// 转换为小写以进行不区分大小写的比较
	fileName = strings.ToLower(fileName)
	// 检查常见图片扩展名
	return strings.HasSuffix(fileName, ".jpg") ||
		strings.HasSuffix(fileName, ".jpeg") ||
		strings.HasSuffix(fileName, ".png") ||
		strings.HasSuffix(fileName, ".gif") //||
		// strings.HasSuffix(fileName, ".bmp") ||
		// strings.HasSuffix(fileName, ".webp")
}


