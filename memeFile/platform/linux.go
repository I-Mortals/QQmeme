//go:build linux

package platform

import "mymeme/memeFile/platform/linux"

// NewClipboard 创建Linux剪贴板实例
func NewClipboard() Clipboard {
	return linux.NewClipboard()
}
