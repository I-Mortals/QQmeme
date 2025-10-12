//go:build darwin

package platform

import "mymeme/memeFile/platform/macos"

// NewClipboard 创建MacOS剪贴板实例
func NewClipboard() Clipboard {
	return macos.NewClipboard()
}
