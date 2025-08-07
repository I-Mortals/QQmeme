package memeFile

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	if err != nil {
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
		strings.HasSuffix(fileName, ".gif") ||
        strings.HasSuffix(fileName, ".bmp") ||
        strings.HasSuffix(fileName, ".webp")
}

type MemeInfo struct {
	Name      string
	Code      string
	ParentPath string
	Icon      string
	Memes     []string
}

// 生成所有meme的路径
func (m *MemeFile) GenerateAllMemePath(rootPath string) []MemeInfo {
	var loadedMemes []MemeInfo

	// 读取根目录下的所有条目
	entries, err := os.ReadDir(rootPath)
	if err != nil {
		log.Println("读取根目录失败:", err.Error())
		return nil
	}

	// 遍历根目录下的每个条目
	for _, entry := range entries {
		// 只处理子文件夹
		if entry.IsDir() {
			// 子文件夹名称
			dirName := entry.Name()
			// 子文件夹完整路径
			dirPath := filepath.Join(rootPath, dirName)

			// 获取子文件夹内的所有图片文件名
			imageNames := m.GetImages(dirPath)
			if imageNames == nil {
				continue // 跳过无图片的文件夹
			}

			// 生成图片完整路径列表
			var memeImages []string
			for _, imgName := range imageNames {
				memeImages = append(memeImages, imgName)
			}

            // 检查是否有图标
            if len(memeImages) <= 0 {
                continue // 跳过无图片的文件夹
            }

			memeInfo := MemeInfo{
				Name:      dirName,    // 文件夹名作为名称
				Code:      dirName,    // 文件夹名作为唯一标识（可根据需要修改）
				ParentPath: dirPath,   // 父文件夹路径
				Icon:  memeImages[0],  // 图标
				Memes: memeImages, // 图片完整路径列表
			}

			loadedMemes = append(loadedMemes, memeInfo)
		}
	}

	return loadedMemes
}
