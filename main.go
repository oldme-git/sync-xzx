package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"syscall"
)

var (
	//go:embed all:frontend/dist
	assets embed.FS
	// dll地址
	xzxApi = syscall.NewLazyDLL("AIO_API.dll")
	// 主机地址
	ip = "172.20.2.220"
	// 端口号
	port = 8500
	// 系统代码
	sysCode = 21
	// 站点号
	terminalNo = 10
)

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "sync-xzx",
		Width:  540,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 100},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
