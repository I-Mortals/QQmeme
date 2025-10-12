//go:build windows

package platform

import "mymeme/memeFile/platform/windows"

// NewClipboard 创建Windows剪贴板实例
func NewClipboard() Clipboard {
	return windows.NewClipboard()
}
