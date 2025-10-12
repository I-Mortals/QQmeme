package memeFile

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"mymeme/memeFile/platform"
	"mymeme/memeFile/sticker"
	"mymeme/memeFile/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// MemeFile 结构体 - 主要的文件操作和剪贴板管理结构
type MemeFile struct {
	ctx        context.Context // Wails应用上下文
	fileUtils  *utils.FileUtils
	imageUtils *utils.ImageUtils
	downloader *sticker.TelegramDownloader
	clipboard  platform.Clipboard // 跨平台剪贴板实例
}

// NewMemeFile 创建新的MemeFile实例
func NewMemeFile() *MemeFile {
	return &MemeFile{
		fileUtils:  utils.NewFileUtils(),
		imageUtils: utils.NewImageUtils(),
		clipboard:  platform.NewClipboard(),
	}
}

// SetContext 设置Wails应用上下文
// 允许调用App运行时方法
func (m *MemeFile) SetContext(ctx context.Context) {
	m.ctx = ctx
}

// OpenFileDlg 打开文件选择对话框，返回选择的文件路径列表
func (m *MemeFile) OpenFileDlg() []string {
	if m.ctx == nil {
		log.Println("错误: 上下文未设置")
		return nil
	}

	opts := runtime.OpenDialogOptions{}
	filePaths, err := runtime.OpenMultipleFilesDialog(m.ctx, opts)
	if err != nil {
		log.Printf("打开文件对话框失败: %v", err)
		return nil
	}

	log.Printf("选择的文件: %v", filePaths)
	return filePaths
}

// SelectRootDir 打开目录选择对话框，选择meme根目录
func (m *MemeFile) SelectRootDir() string {
	if m.ctx == nil {
		log.Println("错误: 上下文未设置")
		return ""
	}

	path, err := runtime.OpenDirectoryDialog(m.ctx, runtime.OpenDialogOptions{
		Title: "请选择meme根目录",
	})
	if err != nil {
		log.Printf("选择目录失败: %v", err)
		return ""
	}

	return path
}

// GetDirs 获取指定目录下的所有子目录
func (m *MemeFile) GetDirs(path string) []string {
	return m.fileUtils.GetDirs(path)
}

// GetImages 获取指定目录下的所有图片文件
func (m *MemeFile) GetImages(path string) []string {
	return m.imageUtils.GetImages(path)
}

// MemeInfo 结构体 - 存储meme信息
type MemeInfo struct {
	Name       string   `json:"name"`       // meme名称
	Code       string   `json:"code"`       // meme唯一标识
	ParentPath string   `json:"parentPath"` // 父目录路径
	Icon       string   `json:"icon"`       // 图标文件名
	Memes      []string `json:"memes"`      // 图片文件列表
}

// GenerateAllMemePath 生成所有meme的信息
func (m *MemeFile) GenerateAllMemePath(rootPath string) []MemeInfo {
	if rootPath == "" {
		log.Println("错误: 根路径为空")
		return nil
	}

	if !m.fileUtils.PathExists(rootPath) {
		log.Printf("根目录不存在: %s", rootPath)
		return nil
	}

	var loadedMemes []MemeInfo

	entries, err := os.ReadDir(rootPath)
	if err != nil {
		log.Printf("读取根目录失败 %s: %v", rootPath, err)
		return nil
	}

	// 遍历根目录下的每个条目
	for _, entry := range entries {
		// 只处理子文件夹
		if !entry.IsDir() {
			continue
		}

		dirName := entry.Name()
		dirPath := filepath.Join(rootPath, dirName)

		// 获取子文件夹内的所有图片文件名
		imageNames := m.GetImages(dirPath)
		if len(imageNames) == 0 {
			log.Printf("跳过无图片的文件夹: %s", dirName)
			continue
		}

		// 创建meme信息
		memeInfo := MemeInfo{
			Name:       dirName,       // 文件夹名作为名称
			Code:       dirName,       // 文件夹名作为唯一标识
			ParentPath: dirPath,       // 父文件夹路径
			Icon:       imageNames[0], // 使用第一张图片作为图标
			Memes:      imageNames,    // 图片文件列表
		}

		loadedMemes = append(loadedMemes, memeInfo)
	}

	log.Printf("成功加载 %d 个meme目录", len(loadedMemes))
	return loadedMemes
}

func (m *MemeFile) WriteFileToClipboard(filePath string) error {
	log.Printf("复制文件到剪贴板: %s", filePath)
	return m.clipboard.WriteFileToClipboard(filePath)
}

func (m *MemeFile) ClipboardHasFiles() bool {
	return m.clipboard.ClipboardHasFiles()
}

func (m *MemeFile) GetFilesFromClipboard() ([]string, error) {
	return m.clipboard.GetFilesFromClipboard()
}

func (m *MemeFile) RenameFilesInOrder(tabName string, folderPath string, fileNames []string) error {
	if tabName == "" {
		return fmt.Errorf("tab名称不能为空")
	}
	if folderPath == "" {
		return fmt.Errorf("文件夹路径不能为空")
	}
	if len(fileNames) == 0 {
		return fmt.Errorf("文件列表不能为空")
	}

	if !m.fileUtils.PathExists(folderPath) {
		return fmt.Errorf("文件夹不存在: %s", folderPath)
	}

	// 创建临时文件名映射，避免重命名冲突
	tempFileNames := make(map[string]string)
	finalFileNames := make(map[string]string)

	// 将所有文件重命名为临时名称
	for i, fileName := range fileNames {
		if fileName == "" {
			continue
		}

		oldPath := filepath.Join(folderPath, fileName)

		if !m.fileUtils.PathExists(oldPath) {
			continue
		}

		ext := m.fileUtils.GetFileExt(fileName)

		// 生成临时文件名
		tempFileName := fmt.Sprintf("temp_%d_%s%s", i, tabName, ext)
		tempPath := filepath.Join(folderPath, tempFileName)

		// 生成最终文件名
		finalFileName := fmt.Sprintf("%s_%02d%s", tabName, i+1, ext)

		// 重命名为临时文件名
		if err := os.Rename(oldPath, tempPath); err != nil {
			return fmt.Errorf("重命名文件失败 %s -> %s: %v", oldPath, tempPath, err)
		}

		tempFileNames[tempFileName] = fileName
		finalFileNames[tempFileName] = finalFileName
	}

	// 将临时文件名重命名为最终文件名
	for tempFileName, finalFileName := range finalFileNames {
		tempPath := filepath.Join(folderPath, tempFileName)
		finalPath := filepath.Join(folderPath, finalFileName)

		if err := os.Rename(tempPath, finalPath); err != nil {
			// 尝试回滚已重命名的文件
			if originalName, exists := tempFileNames[tempFileName]; exists {
				originalPath := filepath.Join(folderPath, originalName)
				os.Rename(tempPath, originalPath)
			}

			return fmt.Errorf("重命名文件失败 %s -> %s: %v", tempFileName, finalFileName, err)
		}
	}

	return nil
}

func (m *MemeFile) RenameFile(folderPath string, oldFileName string, newFileName string) error {
	if folderPath == "" {
		return fmt.Errorf("文件夹路径不能为空")
	}
	if oldFileName == "" {
		return fmt.Errorf("原文件名不能为空")
	}
	if newFileName == "" {
		return fmt.Errorf("新文件名不能为空")
	}

	if !m.fileUtils.PathExists(folderPath) {
		return fmt.Errorf("文件夹不存在: %s", folderPath)
	}

	oldPath := filepath.Join(folderPath, oldFileName)
	newPath := filepath.Join(folderPath, newFileName)

	if !m.fileUtils.PathExists(oldPath) {
		return fmt.Errorf("原文件不存在: %s", oldFileName)
	}

	if m.fileUtils.PathExists(newPath) {
		return fmt.Errorf("文件名已存在: %s", newFileName)
	}

	if err := os.Rename(oldPath, newPath); err != nil {
		return fmt.Errorf("重命名文件失败 %s -> %s: %v", oldFileName, newFileName, err)
	}

	return nil
}

func (m *MemeFile) RenameFoldersInOrder(rootPath string, folderNames []string) error {
	if rootPath == "" {
		return fmt.Errorf("根路径不能为空")
	}
	if len(folderNames) == 0 {
		return fmt.Errorf("文件夹列表不能为空")
	}

	if !m.fileUtils.PathExists(rootPath) {
		return fmt.Errorf("根目录不存在: %s", rootPath)
	}

	// 创建临时文件夹名映射，避免重命名冲突
	tempFolderNames := make(map[string]string)
	finalFolderNames := make(map[string]string)

	for i, folderName := range folderNames {
		if folderName == "" {
			continue
		}

		oldPath := filepath.Join(rootPath, folderName)

		// 检查原文件夹是否存在
		if !m.fileUtils.PathExists(oldPath) {
			continue
		}

		// 生成临时文件夹名
		tempFolderName := fmt.Sprintf("temp_%d_%s", i, folderName)
		tempPath := filepath.Join(rootPath, tempFolderName)

		// 生成最终文件夹名 - 如果已经是 序号_ 开头，则去掉序号部分
		cleanFolderName := m.cleanFolderName(folderName)
		finalFolderName := fmt.Sprintf("%02d_%s", i+1, cleanFolderName)

		// 重命名为临时文件夹名
		if err := os.Rename(oldPath, tempPath); err != nil {
			return fmt.Errorf("重命名文件夹失败 %s -> %s: %v", oldPath, tempPath, err)
		}

		tempFolderNames[tempFolderName] = folderName
		finalFolderNames[tempFolderName] = finalFolderName
	}

	// 将临时文件夹名重命名为最终文件夹名
	for tempFolderName, finalFolderName := range finalFolderNames {
		tempPath := filepath.Join(rootPath, tempFolderName)
		finalPath := filepath.Join(rootPath, finalFolderName)

		if err := os.Rename(tempPath, finalPath); err != nil {
			// 如果失败，回滚已重命名的文件夹
			if originalName, exists := tempFolderNames[tempFolderName]; exists {
				originalPath := filepath.Join(rootPath, originalName)
				os.Rename(tempPath, originalPath)
			}

			return fmt.Errorf("重命名文件夹失败 %s -> %s: %v", tempFolderName, finalFolderName, err)
		}
	}

	return nil
}

func (m *MemeFile) cleanFolderName(folderName string) string {
	if strings.Contains(folderName, "_") {
		parts := strings.SplitN(folderName, "_", 2)
		if len(parts) == 2 && len(parts[0]) > 0 {
			// 检查第一部分是否是数字
			isNumber := true
			for _, r := range parts[0] {
				if r < '0' || r > '9' {
					isNumber = false
					break
				}
			}
			if isNumber {
				return parts[1]
			}
		}
	}
	return folderName
}

func (m *MemeFile) DownloadTgStickerSet(stickerSetName string, savePath string, botToken string, proxyURL string, needProxy bool) error {
	m.downloader = sticker.NewTelegramDownloader(m.ctx, botToken, proxyURL, needProxy)

	progressCallback := func(progress sticker.DownloadProgress) {
		if m.ctx != nil {
			progressData := map[string]interface{}{
				"stickerSetName": stickerSetName,
				"progress":       progress,
			}
			runtime.EventsEmit(m.ctx, "telegram-download-progress", progressData)
		}
	}

	return m.downloader.DownloadTgStickerSet(stickerSetName, savePath, progressCallback)
}

// DeleteTgStickerSet 删除Telegram贴纸集合文件夹
func (m *MemeFile) DeleteTgStickerSet(rootPath string, stickerSetName string) error {
	// 构建贴纸集合的文件夹路径
	stickerSetPath := filepath.Join(rootPath, stickerSetName)

	// 检查文件夹是否存在
	if _, err := os.Stat(stickerSetPath); os.IsNotExist(err) {
		return fmt.Errorf("贴纸集合文件夹不存在: %s", stickerSetName)
	}

	// 删除整个文件夹
	if err := os.RemoveAll(stickerSetPath); err != nil {
		return fmt.Errorf("删除贴纸集合失败: %v", err)
	}

	return nil
}

// DeleteMemeFile 删除单个表情包文件
func (m *MemeFile) DeleteMemeFile(rootPath string, folderCode string, fileName string) error {
	// 构建文件的完整路径
	filePath := filepath.Join(rootPath, folderCode, fileName)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("文件不存在: %s", fileName)
	}

	// 删除文件
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("删除文件失败: %v", err)
	}

	return nil
}
