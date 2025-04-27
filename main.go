package main

import (
	"FastMD5/gcore/app"
	"FastMD5/gcore/hash"
	"embed"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	args := os.Args
	app := app.NewApp()
	hasher := hash.NewHasher(args)

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "",
		Width:         550,
		Height:        250,
		Frameless:     true,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
			hasher,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
