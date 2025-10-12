//go:build darwin

package macos

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework Foundation
#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>

void copyFileToClipboard(const char* filePath) {
    NSString *path = [NSString stringWithUTF8String:filePath];
    NSURL *fileURL = [NSURL fileURLWithPath:path];

    if (fileURL) {
        NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
        [pasteboard clearContents];
        [pasteboard writeObjects:@[fileURL]];
    }
}

int clipboardHasFiles() {
    NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
    NSArray *classes = @[[NSURL class]];
    NSDictionary *options = @{};

    BOOL canRead = [pasteboard canReadObjectForClasses:classes options:options];
    return canRead ? 1 : 0;
}

const char* getFilesFromClipboard() {
    NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
    NSArray *classes = @[[NSURL class]];
    NSDictionary *options = @{};

    NSArray *objects = [pasteboard readObjectsForClasses:classes options:options];
    if (objects && objects.count > 0) {
        NSMutableString *result = [NSMutableString string];
        for (NSURL *url in objects) {
            if (result.length > 0) {
                [result appendString:@"\n"];
            }
            [result appendString:url.path];
        }
        return [result UTF8String];
    }
    return "";
}
*/
import "C"

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// MacOSClipboard MacOS剪贴板实现
type MacOSClipboard struct{}

// NewClipboard 创建MacOS剪贴板实例
func NewClipboard() *MacOSClipboard {
	return &MacOSClipboard{}
}

// WriteFileToClipboard 将文件复制到剪贴板
func (m *MacOSClipboard) WriteFileToClipboard(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("文件路径为空")
	}

	// 规范化路径，确保使用正确的路径分隔符
	normalizedPath := filepath.Clean(filePath)

	if _, err := os.Stat(normalizedPath); os.IsNotExist(err) {
		return fmt.Errorf("文件不存在: %s", normalizedPath)
	}

	fullPath, err := filepath.Abs(normalizedPath)
	if err != nil {
		return fmt.Errorf("获取文件绝对路径失败: %v", err)
	}

	// 使用Cocoa框架复制文件到剪贴板
	C.copyFileToClipboard(C.CString(fullPath))

	return nil
}

// ClipboardHasFiles 检查剪贴板中是否有文件
func (m *MacOSClipboard) ClipboardHasFiles() bool {
	// 使用Cocoa框架检查剪贴板中是否有文件
	result := C.clipboardHasFiles()
	return result == 1
}

// GetFilesFromClipboard 从剪贴板获取文件列表
func (m *MacOSClipboard) GetFilesFromClipboard() ([]string, error) {
	cResult := C.getFilesFromClipboard()
	if cResult == nil {
		return []string{}, nil
	}

	result := C.GoString(cResult)
	if result == "" {
		return []string{}, nil
	}

	// 按行分割文件路径
	filePaths := strings.Split(result, "\n")
	var files []string
	for _, path := range filePaths {
		cleanPath := strings.TrimSpace(path)
		if cleanPath != "" {
			files = append(files, cleanPath)
		}
	}

	return files, nil
}
