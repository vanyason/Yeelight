package main

import (
	"context"
	"embed"
	"yeelight/yeelight"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type App struct {
	ctx context.Context
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := &App{}
	bulb := &yeelight.YLightBulb{}

	err := wails.Run(&options.App{
		Title:            "Yeelight by Vanyason",
		Width:            768,
		Height:           1024,
		MinWidth:         480,
		MinHeight:        640,
		AssetServer:      &assetserver.Options{Assets: assets},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind:             []interface{}{app, bulb},
		OnStartup:        app.startup,
		OnShutdown: func(ctx context.Context) {
			bulb.Disconnect()
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
