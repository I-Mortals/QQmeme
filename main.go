package main

import (
	"embed"

	"mymeme/httpProxy"
	"mymeme/memeFile"

	"context"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 创建应用结构的实例
	app := NewApp()
	memeFile := memeFile.NewMemeFile()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "QQmeme",
		Width:  1024,
		Height: 768,
		MinWidth:  800,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Handler: httpProxy.NewFileLoader(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		// OnStartup: app.startup,
		OnStartup: func(ctx context.Context){
            app.SetContext(ctx)
            memeFile.SetContext(ctx)
        },
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,memeFile,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
