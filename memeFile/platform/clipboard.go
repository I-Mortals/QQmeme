package platform

// Clipboard 定义跨平台剪贴板操作接口
type Clipboard interface {
	// WriteFileToClipboard 将文件路径写入剪贴板
	WriteFileToClipboard(filePath string) error

	// ClipboardHasFiles 检查剪贴板中是否有文件
	ClipboardHasFiles() bool

	// GetFilesFromClipboard 从剪贴板获取文件列表
	GetFilesFromClipboard() ([]string, error)
}
