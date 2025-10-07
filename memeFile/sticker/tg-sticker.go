package sticker

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	rlottie "github.com/yazmeyaa/go-rlottie"
)

type TelegramSticker struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	IsAnimated   bool   `json:"is_animated"`
	IsVideo      bool   `json:"is_video"`
}

type TelegramAPIResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Name        string            `json:"name"`
		Title       string            `json:"title"`
		Description string            `json:"description"`
		IsAnimated  bool              `json:"is_animated"`
		IsVideo     bool              `json:"is_video"`
		Stickers    []TelegramSticker `json:"stickers"`
	} `json:"result"`
	Description string `json:"description"`
}

type TelegramFileResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		FileID   string `json:"file_id"`
		FilePath string `json:"file_path"`
	} `json:"result"`
}

type DownloadProgress struct {
	Current      int     `json:"current"`
	Total        int     `json:"total"`
	Status       string  `json:"status"`
	Percentage   float64 `json:"percentage"`
	SuccessCount int     `json:"successCount"`
	FailedCount  int     `json:"failedCount"`
}

type TelegramDownloader struct {
	ctx        context.Context
	botToken   string
	proxyURL   string
	client     *http.Client
	ffmpegPath string
}

func NewTelegramDownloader(ctx context.Context, botToken, proxyURL string, needProxy bool) *TelegramDownloader {
	var transport *http.Transport

	if needProxy && proxyURL != "" {
		parsedProxyURL, err := url.Parse(proxyURL)
		if err == nil {
			transport = &http.Transport{
				Proxy: http.ProxyURL(parsedProxyURL),
			}
		}
	}

	if transport == nil {
		// 不设置代理
		transport = &http.Transport{}
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	return &TelegramDownloader{
		ctx:        ctx,
		botToken:   botToken,
		proxyURL:   proxyURL,
		client:     client,
		ffmpegPath: "ffmpeg",
	}
}

func (td *TelegramDownloader) DownloadTgStickerSet(stickerSetName string, savePath string, progressCallback func(DownloadProgress)) error {
	// 获取贴纸集
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/getStickerSet?name=%s", td.botToken, stickerSetName)
	resp, err := td.client.Get(apiURL)
	if err != nil {
		return fmt.Errorf("获取贴纸集合失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	var apiResp TelegramAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return fmt.Errorf("解析响应失败: %v", err)
	}

	if !apiResp.OK {
		return fmt.Errorf("API 错误: %s", apiResp.Description)
	}

	// 保存到 rootPath/sticker集合名字 下面
	saveDir := filepath.Join(savePath, stickerSetName)
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	stickers := apiResp.Result.Stickers

	total := len(stickers)

	// 更新进度
	progressCallback(DownloadProgress{
		Current:      0,
		Total:        total,
		Status:       fmt.Sprintf("开始下载 %s (%d 个贴纸)", apiResp.Result.Title, total),
		Percentage:   0.0,
		SuccessCount: 0,
		FailedCount:  0,
	})

	successCount := 0
	failedCount := 0

	for i, sticker := range stickers {
		// 开始下载
		progressCallback(DownloadProgress{
			Current:      i + 1,
			Total:        total,
			Status:       fmt.Sprintf("下载贴纸 %d/%d", i+1, total),
			Percentage:   float64(i+1) / float64(total) * 100,
			SuccessCount: successCount,
			FailedCount:  failedCount,
		})

		if err := td.downloadSticker(sticker, saveDir, i+1, total); err != nil {
			log.Printf("下载贴纸失败 %s: %v", sticker.FileID, err)
			failedCount++
		} else {
			successCount++
		}

		// 下载完成
		progressCallback(DownloadProgress{
			Current:      i + 1,
			Total:        total,
			Status:       fmt.Sprintf("下载完成 %d/%d (成功: %d, 失败: %d)", i+1, total, successCount, failedCount),
			Percentage:   float64(i+1) / float64(total) * 100,
			SuccessCount: successCount,
			FailedCount:  failedCount,
		})
	}

	log.Printf("下载完成: 成功 %d 个，失败 %d 个", successCount, failedCount)

	stickerData := map[string]interface{}{
		"title": apiResp.Result.Title,
		"name":  apiResp.Result.Name,
		"icon":  nil,
		"url":   fmt.Sprintf("https://t.me/addstickers/%s", stickerSetName),
	}

	stickerDataPath := filepath.Join(saveDir, "sticker.json")
	data, _ := json.MarshalIndent(stickerData, "", "  ")
	os.WriteFile(stickerDataPath, data, 0644)

	// 完成
	progressCallback(DownloadProgress{
		Current:      total,
		Total:        total,
		Status:       fmt.Sprintf("下载完成: %s (成功: %d, 失败: %d)", apiResp.Result.Title, successCount, failedCount),
		Percentage:   100.0,
		SuccessCount: successCount,
		FailedCount:  failedCount,
	})

	return nil
}

func (td *TelegramDownloader) downloadSticker(sticker TelegramSticker, saveDir string, current, total int) error {
	log.Printf("开始下载第 %d/%d 张贴纸，ID: %s", current, total, sticker.FileID)

	fileURL := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", td.botToken, sticker.FileID)
	resp, err := td.client.Get(fileURL)
	if err != nil {
		return fmt.Errorf("获取文件信息失败: %v", err)
	}
	defer resp.Body.Close()

	var fileResp TelegramFileResponse
	if err := json.NewDecoder(resp.Body).Decode(&fileResp); err != nil {
		return fmt.Errorf("解析文件信息失败: %v", err)
	}

	if !fileResp.OK {
		return fmt.Errorf("获取文件信息失败")
	}

	downloadURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", td.botToken, fileResp.Result.FilePath)
	log.Printf("下载链接: %s", downloadURL)
	fileResp2, err := td.client.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("下载文件失败: %v", err)
	}
	defer fileResp2.Body.Close()

	fileData, err := io.ReadAll(fileResp2.Body)
	if err != nil {
		return fmt.Errorf("读取文件内容失败: %v", err)
	}

	fileName := sticker.FileUniqueID

	log.Printf("文件类型: %s, 尺寸: %dx%d",
		func() string {
			if sticker.IsVideo {
				return "视频贴纸"
			} else if sticker.IsAnimated {
				return "TGS动画贴纸"
			} else {
				return "静态贴纸"
			}
		}(), sticker.Width, sticker.Height)

	if sticker.IsVideo {
		fileName += ".gif"
		filePath := filepath.Join(saveDir, fileName)
		return td.convertWebmToGif(fileData, filePath)
	} else if sticker.IsAnimated {
		fileName += ".gif"
		filePath := filepath.Join(saveDir, fileName)
		width := sticker.Width / 2
		height := sticker.Height / 2
		if width == 0 {
			width = 256
		}
		if height == 0 {
			height = 256
		}
		return td.convertTGSToGif(fileData, filePath, width, height)
	} else {
		fileName += ".png"
		filePath := filepath.Join(saveDir, fileName)
		return td.convertWebpToPng(fileData, filePath)
	}
}

func (td *TelegramDownloader) convertWebpToPng(data []byte, outputPath string) error {
	cmd := exec.Command(td.ffmpegPath,
		"-i", "pipe:0",
		"-vf", "scale=512:-1:flags=lanczos",
		"-c:v", "png",
		"-pix_fmt", "rgba",
		"-y", outputPath)

	cmd.Stdin = bytes.NewReader(data)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("WebP 转 PNG 失败: %v, 错误: %s", err, stderr.String())
	}

	return nil
}

func (td *TelegramDownloader) convertWebmToGif(data []byte, outputPath string) error {
	cmd := exec.Command(td.ffmpegPath,
		"-vcodec", "libvpx-vp9",
		"-i", "pipe:0",
		"-vf", "split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse",
		"-loop", "0",
		"-y", outputPath)

	cmd.Stdin = bytes.NewReader(data)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Printf("WebM 转 GIF 失败: %v, 错误: %s", err, stderr.String())
	}

	return nil
}

func (td *TelegramDownloader) convertTGSToGif(data []byte, outputPath string, width, height int) error {
	// 解压 TGS 文件
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("无法创建 GZIP 读取器: %w", err)
	}
	defer gz.Close()

	jsonData, err := io.ReadAll(gz)
	if err != nil {
		return fmt.Errorf("无法读取解压后的 JSON 数据: %w", err)
	}

	// 清除rlottie缓存
	rlottie.LottieConfigureModelCacheSize(0) // 禁用缓存

	// 使用 rlottie 加载动画，使用唯一key避免缓存
	uniqueKey := fmt.Sprintf("tgs_%s", outputPath)
	anim := rlottie.LottieAnimationFromData(string(jsonData), uniqueKey, "")
	if anim == nil {
		return fmt.Errorf("无法从 JSON 数据加载 Lottie 动画")
	}
	defer rlottie.LottieAnimationDestroy(anim)

	duration := rlottie.LottieAnimationGetDuration(anim)
	frameRate := rlottie.LottieAnimationGetFramerate(anim)
	originalTotalFrames := int(rlottie.LottieAnimationGetTotalframe(anim))

	if originalTotalFrames == 0 {
		return fmt.Errorf("错误: 动画不包含任何帧")
	}

	// 计算帧数
	var totalFrames int
	var delay int

	if frameRate > 0 {
		// 如果原始帧率很高，适当降低但保持流畅
		if frameRate >= 30 {
			totalFrames = 24 // 高帧率动画用24帧
		} else if frameRate >= 15 {
			totalFrames = 20 // 中等帧率用20帧
		} else {
			totalFrames = 16 // 低帧率用16帧
		}
	} else {
		// 如果无法获取原始帧率，根据时长判断
		if duration <= 2.0 {
			totalFrames = 24
		} else if duration <= 5.0 {
			totalFrames = 20
		} else {
			totalFrames = 16
		}
	}

	// 计算延迟
	if duration > 0 {
		// 根据动画总时长和帧数计算每帧延迟
		delayPerFrame := duration / float64(totalFrames) // 每帧的秒数
		delay = int(delayPerFrame * 100)                 // 转换为百分之一秒

		// 限制延迟范围，避免过快或过慢
		if delay < 2 {
			delay = 2 // 最小2/100秒 = 20ms
		} else if delay > 20 {
			delay = 20 // 最大20/100秒 = 200ms
		}
	} else {
		delay = 4 // 默认4/100秒 = 40ms
	}

	log.Printf("TGS 动画信息: 时长 %.2f 秒, 原始帧率 %.2f, 原始帧数 %d, 输出帧数 %d, 延迟 %d/100秒",
		duration, frameRate, originalTotalFrames, totalFrames, delay)

	// 预分配所有切片
	outGif := &gif.GIF{
		Image:    make([]*image.Paletted, 0, totalFrames),
		Delay:    make([]int, 0, totalFrames),
		Disposal: make([]byte, totalFrames),
	}

	widthUint := uint(width)
	heightUint := uint(height)

	for i := 0; i < totalFrames; i++ {
		progress := float64(i) / float64(totalFrames-1) // 0 到 1 的进度
		originalFrameNum := uint(progress * float64(originalTotalFrames-1))

		buffer := make([]byte, widthUint*heightUint*4)
		rgbaBuffer := make([]byte, widthUint*heightUint*4)

		rlottie.LottieAnimationRender(anim, originalFrameNum, buffer, widthUint, heightUint, widthUint*4)

		for j := 0; j < len(buffer); j += 4 {
			// 交换 R 和 B 通道：BGRA -> RGBA
			rgbaBuffer[j] = buffer[j+2]   // R = B
			rgbaBuffer[j+1] = buffer[j+1] // G = G
			rgbaBuffer[j+2] = buffer[j]   // B = R
			rgbaBuffer[j+3] = buffer[j+3] // A = A
		}

		rgbaFrame := &image.RGBA{
			Pix:    rgbaBuffer,
			Stride: int(widthUint * 4),
			Rect:   image.Rect(0, 0, int(widthUint), int(heightUint)),
		}

		// 将 RGBA 帧转换为支持透明的 Paletted 帧
		palettedFrame := td.rgbaToPaletted(rgbaFrame)

		outGif.Image = append(outGif.Image, palettedFrame)
		outGif.Delay = append(outGif.Delay, delay)
		outGif.Disposal[i] = gif.DisposalBackground

		if (i+1)%10 == 0 || i == totalFrames-1 {
			log.Printf("TGS 转换进度: %d / %d 帧...", i+1, totalFrames)
		}
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("无法创建输出文件 '%s': %w", outputPath, err)
	}
	defer f.Close()

	if err := gif.EncodeAll(f, outGif); err != nil {
		return fmt.Errorf("编码 GIF 失败: %w", err)
	}

	log.Printf("TGS 转 GIF 成功: %s", outputPath)
	return nil
}

// 预定义调色板
var globalPalette = func() color.Palette {
	p := color.Palette{color.Transparent}
	p = append(p, palette.Plan9[:255]...)
	return p
}()

// rgbaToPaletted 将 RGBA 图像转换为 Paletted 图像，保留透明度并使用抖动算法保证颜色质量
func (td *TelegramDownloader) rgbaToPaletted(rgba *image.RGBA) *image.Paletted {
	bounds := rgba.Bounds()
	paletted := image.NewPaletted(bounds, globalPalette)

	// 使用 Floyd-Steinberg 抖动算法进行绘制，以保证颜色质量
	draw.FloydSteinberg.Draw(paletted, bounds, rgba, image.Point{})

	// 优化透明区域处理：直接操作像素数据而不是逐个像素访问
	pix := rgba.Pix
	palettedPix := paletted.Pix
	stride := rgba.Stride

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			idx := y*stride + x*4
			if pix[idx+3] < 128 { // Alpha < 128
				palettedIdx := y*paletted.Stride + x
				palettedPix[palettedIdx] = 0 // 设置为透明色索引
			}
		}
	}
	return paletted
}
