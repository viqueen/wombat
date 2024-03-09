package main

import (
	"embed"
	_ "embed"
	"os"
	"wombat/internal/app"
)

//go:embed frontend/public/build/bundle.js
var js string

//go:embed frontend/public/build/bundle.css
var css string

//go:embed frontend/public/build/extra.css
var extra string

//go:embed all:frontend/public/build
var assets embed.FS

func main() {
	os.Exit(app.Run(assets))
}
