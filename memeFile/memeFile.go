package memeFile

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Windows API 常量定义
const (
	// 剪贴板格式常量
	CF_UNICODETEXT = 13 // Unicode文本格式
	CF_HDROP       = 15 // 文件拖放格式
	
	// 内存分配标志
	GMEM_MOVEABLE  = 0x0002 // 可移动内存
	GMEM_ZEROINIT  = 0x0040 // 初始化为零
	GHND           = GMEM_MOVEABLE | GMEM_ZEROINIT // 组合标志
)

// Windows API 函数声明
var (
	// DLL 句柄
	user32   = syscall.MustLoadDLL("user32.dll")
	kernel32 = syscall.MustLoadDLL("kernel32.dll")
	shell32  = syscall.MustLoadDLL("shell32.dll")

	// 剪贴板相关API
	procOpenClipboard    = user32.MustFindProc("OpenClipboard")
	procCloseClipboard   = user32.MustFindProc("CloseClipboard")
	procEmptyClipboard   = user32.MustFindProc("EmptyClipboard")
	procSetClipboardData = user32.MustFindProc("SetClipboardData")
	procGetClipboardData = user32.MustFindProc("GetClipboardData")
	
	// 内存管理API
	procGlobalAlloc = kernel32.MustFindProc("GlobalAlloc")
	procGlobalLock  = kernel32.MustFindProc("GlobalLock")
	procGlobalUnlock = kernel32.MustFindProc("GlobalUnlock")

	// 文件拖放API
	procDragQueryFile = shell32.MustFindProc("DragQueryFileW")
)

// DROPFILES 结构体定义 - Windows文件拖放结构
type DROPFILES struct {
	pFiles uint32 // 文件列表偏移量
	pt     POINT  // 拖放点坐标
	fNC    int32  // 非客户区标志
	fWide  int32  // 宽字符标志
}

// POINT 结构体定义 - 坐标点
type POINT struct {
	X, Y int32
}

// MemeFile 结构体 - 主要的文件操作和剪贴板管理结构
type MemeFile struct {
	ctx context.Context // Wails应用上下文
}

// NewMemeFile 创建新的MemeFile实例
func NewMemeFile() *MemeFile {
	return &MemeFile{}
}

// SetContext 设置Wails应用上下文
// 允许调用App运行时方法
func (m *MemeFile) SetContext(ctx context.Context) {
	m.ctx = ctx
}

// Hello 测试方法
func (m *MemeFile) Hello() string {
	log.Println("MemeFile Hello method called")
	return "hello"
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
	if path == "" {
		log.Println("错误: 路径为空")
		return nil
	}
	
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Printf("读取目录失败 %s: %v", path, err)
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

// GetImages 获取指定目录下的所有图片文件
func (m *MemeFile) GetImages(path string) []string {
	if path == "" {
		log.Println("错误: 路径为空")
		return nil
	}
	
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Printf("读取目录失败 %s: %v", path, err)
		return nil
	}

	var images []string
	for _, entry := range entries {
		if !entry.IsDir() && isImageFile(entry.Name()) {
			images = append(images, entry.Name())
		}
	}
	return images
}

// 支持的图片扩展名（预定义以提高性能）
var imageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}

// isImageFile 检查文件是否为支持的图片格式
func isImageFile(fileName string) bool {
	if fileName == "" {
		return false
	}
	
	// 转换为小写进行不区分大小写的比较
	fileName = strings.ToLower(fileName)
	
	// 使用预定义的扩展名列表
	for _, ext := range imageExtensions {
		if strings.HasSuffix(fileName, ext) {
			return true
		}
	}
	return false
}

// MemeInfo 结构体 - 存储meme信息
type MemeInfo struct {
	Name       string   `json:"name"`       // meme名称
	Code       string   `json:"code"`       // meme唯一标识
	ParentPath string   `json:"parentPath"` // 父目录路径
	Icon       string   `json:"icon"`       // 图标文件名
	Memes      []string `json:"memes"`      // 图片文件列表
}

// GenerateAllMemePath 扫描根目录，生成所有meme的信息
func (m *MemeFile) GenerateAllMemePath(rootPath string) []MemeInfo {
	if rootPath == "" {
		log.Println("错误: 根路径为空")
		return nil
	}
	
	// 检查根目录是否存在
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		log.Printf("根目录不存在: %s", rootPath)
		return nil
	}

	var loadedMemes []MemeInfo

	// 读取根目录下的所有条目
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

// Windows API 封装函数

// OpenClipboard 打开剪贴板
func OpenClipboard(hWnd uintptr) bool {
	ret, _, _ := procOpenClipboard.Call(hWnd)
	return ret != 0
}

// CloseClipboard 关闭剪贴板
func CloseClipboard() bool {
	ret, _, _ := procCloseClipboard.Call()
	return ret != 0
}

// EmptyClipboard 清空剪贴板
func EmptyClipboard() bool {
	ret, _, _ := procEmptyClipboard.Call()
	return ret != 0
}

// SetClipboardData 设置剪贴板数据
func SetClipboardData(uFormat uint, hMem uintptr) uintptr {
	ret, _, _ := procSetClipboardData.Call(uintptr(uFormat), hMem)
	return ret
}

// GlobalAlloc 分配全局内存
func GlobalAlloc(uFlags uint, dwBytes uintptr) uintptr {
	ret, _, _ := procGlobalAlloc.Call(uintptr(uFlags), dwBytes)
	return ret
}

// GlobalLock 锁定全局内存
func GlobalLock(hMem uintptr) uintptr {
	ret, _, _ := procGlobalLock.Call(hMem)
	return ret
}

// GlobalUnlock 解锁全局内存
// 注意：即使返回0也不一定是错误，特别是在内存块已经解锁的情况下
func GlobalUnlock(hMem uintptr) bool {
	ret, _, _ := procGlobalUnlock.Call(hMem)
	return ret != 0
}

// WriteFileToClipboard 将文件复制到剪贴板
func (m *MemeFile) WriteFileToClipboard(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("文件路径为空")
	}
	
	// 1. 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("文件不存在: %s", filePath)
	}

	// 2. 获取文件的完整路径
	fullPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("获取文件绝对路径失败: %v", err)
	}
	
	log.Printf("复制文件到剪贴板: %s", fullPath)

	// 3. 打开剪贴板
	if !OpenClipboard(0) {
		return fmt.Errorf("打开剪贴板失败")
	}
	defer CloseClipboard()

	// 4. 清空剪贴板
	if !EmptyClipboard() {
		return fmt.Errorf("清空剪贴板失败")
	}

	// 5. 将文件路径转换为UTF-16格式
	runes := []rune(fullPath)
	utf16Path := utf16.Encode(runes)
	// 添加双空终止符（Windows文件列表要求）
	utf16Path = append(utf16Path, 0) // 第一个空终止符
	utf16Path = append(utf16Path, 0) // 第二个空终止符（列表结束）

	// 6. 计算DROPFILES结构大小和总内存大小
	dropFilesSize := unsafe.Sizeof(DROPFILES{})
	pathSize := len(utf16Path) * 2
	totalSize := dropFilesSize + uintptr(pathSize)

	// 7. 分配全局内存
	hGlobal := GlobalAlloc(GHND, totalSize)
	if hGlobal == 0 {
		return fmt.Errorf("内存分配失败")
	}

	// 8. 锁定内存并获取指针
	ptr := GlobalLock(hGlobal)
	if ptr == 0 {
		return fmt.Errorf("内存锁定失败")
	}

	// 9. 设置DROPFILES结构
	// 注意：这里的unsafe.Pointer使用是安全的，因为ptr来自GlobalLock
	dropFiles := (*DROPFILES)(unsafe.Pointer(ptr))
	dropFiles.pFiles = uint32(dropFilesSize)
	dropFiles.pt.X = 0
	dropFiles.pt.Y = 0
	dropFiles.fNC = 0
	dropFiles.fWide = 1 // 使用Unicode

	// 10. 安全地复制文件路径到内存
	if pathSize > 0 && len(utf16Path) > 0 {
		// 注意：这里的unsafe.Pointer使用是安全的，因为ptr来自GlobalLock
		pathPtr := unsafe.Pointer(uintptr(ptr) + dropFilesSize)
		// 使用字节切片进行更安全的内存操作
		pathByteSlice := (*[1 << 21]byte)(pathPtr)[:pathSize:pathSize]
		pathUint16Slice := (*[1 << 20]uint16)(unsafe.Pointer(&pathByteSlice[0]))
		if len(utf16Path) <= len(pathUint16Slice) {
			copy(pathUint16Slice[:len(utf16Path)], utf16Path)
		}
	}

	// 11. 解锁内存
	_ = GlobalUnlock(hGlobal)

	// 12. 设置剪贴板数据
	if SetClipboardData(CF_HDROP, hGlobal) == 0 {
		return fmt.Errorf("设置剪贴板数据失败")
	}

	return nil
}

// WriteGIFToClipboard 专门用于复制GIF文件到剪贴板
func (m *MemeFile) WriteGIFToClipboard(gifPath string) error {
	if gifPath == "" {
		return fmt.Errorf("GIF文件路径为空")
	}
	
	// 检查文件扩展名
	ext := strings.ToLower(filepath.Ext(gifPath))
	if ext != ".gif" {
		return fmt.Errorf("文件不是GIF格式: %s", gifPath)
	}

	return m.WriteFileToClipboard(gifPath)
}

// ClipboardHasFiles 检查剪贴板中是否有文件
func ClipboardHasFiles() bool {
	if !OpenClipboard(0) {
		return false
	}
	defer CloseClipboard()

	hGlobal, _, _ := procGetClipboardData.Call(uintptr(CF_HDROP))
	return hGlobal != 0
}

// GetFilesFromClipboard 从剪贴板获取文件列表
func GetFilesFromClipboard() ([]string, error) {
	if !OpenClipboard(0) {
		return nil, fmt.Errorf("打开剪贴板失败")
	}
	defer CloseClipboard()

	hGlobal, _, _ := procGetClipboardData.Call(uintptr(CF_HDROP))
	if hGlobal == 0 {
		return nil, fmt.Errorf("剪贴板中没有文件数据")
	}

	ptr := GlobalLock(hGlobal)
	if ptr == 0 {
		return nil, fmt.Errorf("锁定内存失败")
	}
	defer GlobalUnlock(hGlobal)

	// 获取文件数量
	fileCount, _, _ := procDragQueryFile.Call(hGlobal, 0xFFFFFFFF, 0, 0)
	if fileCount == 0 {
		return []string{}, nil
	}

	var files []string
	for i := uintptr(0); i < fileCount; i++ {
		// 获取文件名长度
		length, _, _ := procDragQueryFile.Call(hGlobal, i, 0, 0)
		if length == 0 {
			continue
		}

		// 分配缓冲区并获取文件名
		buffer := make([]uint16, length+1)
		if len(buffer) > 0 {
			// 使用更安全的方式获取缓冲区指针
			// 注意：这里的unsafe.Pointer使用是安全的，因为buffer是Go分配的切片
			bufferPtr := &buffer[0]
			procDragQueryFile.Call(hGlobal, i, uintptr(unsafe.Pointer(bufferPtr)), length+1)
		}

		fileName := syscall.UTF16ToString(buffer)
		if fileName != "" {
			files = append(files, fileName)
		}
	}

	return files, nil
}
