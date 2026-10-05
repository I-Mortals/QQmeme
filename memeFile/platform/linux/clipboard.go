//go:build linux

package linux

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// LinuxClipboard Linux 剪贴板实现
//
// Linux 桌面环境没有统一的文件剪贴板 API，这里通过系统工具实现：
//   - Wayland: wl-clipboard (wl-copy / wl-paste)
//   - X11:     xclip
//
// 文件以 text/uri-list 的形式写入剪贴板，与 Windows 的 CF_HDROP、
// macOS 的 NSURL 行为对应。
type LinuxClipboard struct{}

// NewClipboard 创建Linux剪贴板实例
func NewClipboard() *LinuxClipboard {
	return &LinuxClipboard{}
}

type clipboardBackend struct {
	copyCmd  string
	pasteCmd string
}

// detectBackend 根据当前会话与可用工具选择剪贴板后端
func detectBackend() (clipboardBackend, bool) {
	has := func(name string) bool {
		_, err := exec.LookPath(name)
		return err == nil
	}

	if has("wl-copy") && has("wl-paste") {
		return clipboardBackend{copyCmd: "wl-copy", pasteCmd: "wl-paste"}, true
	}
	if has("xclip") {
		return clipboardBackend{copyCmd: "xclip", pasteCmd: "xclip"}, true
	}
	return clipboardBackend{}, false
}

func isWayland(cmd string) bool {
	return filepath.Base(cmd) == "wl-copy" || filepath.Base(cmd) == "wl-paste"
}

// WriteFileToClipboard 将文件以 text/uri-list 形式写入剪贴板
func (l *LinuxClipboard) WriteFileToClipboard(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("文件路径为空")
	}

	normalizedPath := filepath.Clean(filePath)
	if _, err := os.Stat(normalizedPath); err != nil {
		return fmt.Errorf("文件不存在: %s", normalizedPath)
	}

	absPath, err := filepath.Abs(normalizedPath)
	if err != nil {
		return fmt.Errorf("获取文件绝对路径失败: %v", err)
	}

	uri := (&url.URL{Scheme: "file", Path: absPath}).String() + "\r\n"

	backend, ok := detectBackend()
	if !ok {
		return fmt.Errorf("未找到剪贴板工具，请安装 wl-clipboard (Wayland) 或 xclip (X11)")
	}

	var cmd *exec.Cmd
	if isWayland(backend.copyCmd) {
		cmd = exec.Command(backend.copyCmd, "--type", "text/uri-list")
	} else {
		cmd = exec.Command(backend.copyCmd, "-selection", "clipboard", "-t", "text/uri-list", "-i")
	}
	cmd.Stdin = strings.NewReader(uri)

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("写入剪贴板失败: %v: %s", err, strings.TrimSpace(string(out)))
	}

	return nil
}

// ClipboardHasFiles 检查剪贴板中是否存在文件列表
func (l *LinuxClipboard) ClipboardHasFiles() bool {
	backend, ok := detectBackend()
	if !ok {
		return false
	}

	out, err := readClipboard(backend.pasteCmd, true)
	if err != nil {
		return false
	}
	return strings.Contains(out, "text/uri-list")
}

// GetFilesFromClipboard 从剪贴板获取文件路径列表
func (l *LinuxClipboard) GetFilesFromClipboard() ([]string, error) {
	backend, ok := detectBackend()
	if !ok {
		return []string{}, nil
	}

	out, err := readClipboard(backend.pasteCmd, false)
	if err != nil {
		return []string{}, nil
	}

	var files []string
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		u, err := url.Parse(line)
		if err != nil || u.Scheme != "file" {
			continue
		}
		files = append(files, u.Path)
	}

	return files, nil
}

// readClipboard 读取剪贴板内容。listTypes 为 true 时返回可用的目标类型列表
func readClipboard(pasteCmd string, listTypes bool) (string, error) {
	var cmd *exec.Cmd
	if isWayland(pasteCmd) {
		if listTypes {
			cmd = exec.Command(pasteCmd, "--list-types")
		} else {
			cmd = exec.Command(pasteCmd, "--type", "text/uri-list")
		}
	} else {
		if listTypes {
			cmd = exec.Command(pasteCmd, "-selection", "clipboard", "-t", "TARGETS", "-o")
		} else {
			cmd = exec.Command(pasteCmd, "-selection", "clipboard", "-t", "text/uri-list", "-o")
		}
	}

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
