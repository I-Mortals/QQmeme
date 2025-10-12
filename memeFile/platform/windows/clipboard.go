//go:build windows

package windows

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// Windows API 函数声明
var (
	// DLL 句柄
	User32   = syscall.MustLoadDLL("user32.dll")
	Kernel32 = syscall.MustLoadDLL("kernel32.dll")
	Shell32  = syscall.MustLoadDLL("shell32.dll")

	// 剪贴板相关API
	ProcOpenClipboard    = User32.MustFindProc("OpenClipboard")
	ProcCloseClipboard   = User32.MustFindProc("CloseClipboard")
	ProcEmptyClipboard   = User32.MustFindProc("EmptyClipboard")
	ProcSetClipboardData = User32.MustFindProc("SetClipboardData")
	ProcGetClipboardData = User32.MustFindProc("GetClipboardData")

	// 内存管理API
	ProcGlobalAlloc  = Kernel32.MustFindProc("GlobalAlloc")
	ProcGlobalLock   = Kernel32.MustFindProc("GlobalLock")
	ProcGlobalUnlock = Kernel32.MustFindProc("GlobalUnlock")

	// 文件拖放API
	ProcDragQueryFile = Shell32.MustFindProc("DragQueryFileW")
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

const (
	// 剪贴板格式常量
	CF_UNICODETEXT = 13 // Unicode文本格式
	CF_HDROP       = 15 // 文件拖放格式

	// 内存分配标志
	GMEM_MOVEABLE = 0x0002                        // 可移动内存
	GMEM_ZEROINIT = 0x0040                        // 初始化为零
	GHND          = GMEM_MOVEABLE | GMEM_ZEROINIT // 组合标志
)

// Windows API 封装函数

// OpenClipboard 打开剪贴板
func OpenClipboard(hWnd uintptr) bool {
	ret, _, _ := ProcOpenClipboard.Call(hWnd)
	return ret != 0
}

// CloseClipboard 关闭剪贴板
func CloseClipboard() bool {
	ret, _, _ := ProcCloseClipboard.Call()
	return ret != 0
}

// EmptyClipboard 清空剪贴板
func EmptyClipboard() bool {
	ret, _, _ := ProcEmptyClipboard.Call()
	return ret != 0
}

func SetClipboardData(uFormat uint, hMem uintptr) uintptr {
	ret, _, _ := ProcSetClipboardData.Call(uintptr(uFormat), hMem)
	return ret
}

func GlobalAlloc(uFlags uint, dwBytes uintptr) uintptr {
	ret, _, _ := ProcGlobalAlloc.Call(uintptr(uFlags), dwBytes)
	return ret
}

func GlobalLock(hMem uintptr) uintptr {
	ret, _, _ := ProcGlobalLock.Call(hMem)
	return ret
}

func GlobalUnlock(hMem uintptr) bool {
	ret, _, _ := ProcGlobalUnlock.Call(hMem)
	return ret != 0
}

func WriteFileToClipboard(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("文件路径为空")
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("文件不存在: %s", filePath)
	}

	fullPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("获取文件绝对路径失败: %v", err)
	}

	if !OpenClipboard(0) {
		return fmt.Errorf("打开剪贴板失败")
	}
	defer CloseClipboard()

	if !EmptyClipboard() {
		return fmt.Errorf("清空剪贴板失败")
	}

	runes := []rune(fullPath)
	utf16Path := utf16.Encode(runes)

	utf16Path = append(utf16Path, 0)
	utf16Path = append(utf16Path, 0)

	dropFilesSize := unsafe.Sizeof(DROPFILES{})
	pathSize := len(utf16Path) * 2
	totalSize := dropFilesSize + uintptr(pathSize)

	hGlobal := GlobalAlloc(GHND, totalSize)
	if hGlobal == 0 {
		return fmt.Errorf("内存分配失败")
	}

	ptr := GlobalLock(hGlobal)
	if ptr == 0 {
		return fmt.Errorf("内存锁定失败")
	}

	dropFiles := (*DROPFILES)(unsafe.Pointer(ptr))
	dropFiles.pFiles = uint32(dropFilesSize)
	dropFiles.pt.X = 0
	dropFiles.pt.Y = 0
	dropFiles.fNC = 0
	dropFiles.fWide = 1 // 使用Unicode

	if pathSize > 0 && len(utf16Path) > 0 {
		pathPtr := unsafe.Pointer(uintptr(ptr) + dropFilesSize)
		pathByteSlice := (*[1 << 21]byte)(pathPtr)[:pathSize:pathSize]
		pathUint16Slice := (*[1 << 20]uint16)(unsafe.Pointer(&pathByteSlice[0]))
		if len(utf16Path) <= len(pathUint16Slice) {
			copy(pathUint16Slice[:len(utf16Path)], utf16Path)
		}
	}

	_ = GlobalUnlock(hGlobal)

	if SetClipboardData(CF_HDROP, hGlobal) == 0 {
		return fmt.Errorf("设置剪贴板数据失败")
	}

	return nil
}

// ClipboardHasFiles 检查剪贴板中是否有文件
func ClipboardHasFiles() bool {
	if !OpenClipboard(0) {
		return false
	}
	defer CloseClipboard()

	hGlobal, _, _ := ProcGetClipboardData.Call(uintptr(CF_HDROP))
	return hGlobal != 0
}

// GetFilesFromClipboard 从剪贴板获取文件列表
func GetFilesFromClipboard() ([]string, error) {
	if !OpenClipboard(0) {
		return nil, fmt.Errorf("打开剪贴板失败")
	}
	defer CloseClipboard()

	hGlobal, _, _ := ProcGetClipboardData.Call(uintptr(CF_HDROP))
	if hGlobal == 0 {
		return nil, fmt.Errorf("剪贴板中没有文件数据")
	}

	ptr := GlobalLock(hGlobal)
	if ptr == 0 {
		return nil, fmt.Errorf("锁定内存失败")
	}
	defer GlobalUnlock(hGlobal)

	// 获取文件数量
	fileCount, _, _ := ProcDragQueryFile.Call(hGlobal, 0xFFFFFFFF, 0, 0)
	if fileCount == 0 {
		return []string{}, nil
	}

	var files []string
	for i := uintptr(0); i < fileCount; i++ {
		// 获取文件名长度
		length, _, _ := ProcDragQueryFile.Call(hGlobal, i, 0, 0)
		if length == 0 {
			continue
		}

		// 分配缓冲区并获取文件名
		buffer := make([]uint16, length+1)
		if len(buffer) > 0 {
			// 使用更安全的方式获取缓冲区指针
			// 注意：这里的unsafe.Pointer使用是安全的，因为buffer是Go分配的切片
			bufferPtr := &buffer[0]
			ProcDragQueryFile.Call(hGlobal, i, uintptr(unsafe.Pointer(bufferPtr)), length+1)
		}

		fileName := syscall.UTF16ToString(buffer)
		if fileName != "" {
			files = append(files, fileName)
		}
	}

	return files, nil
}
