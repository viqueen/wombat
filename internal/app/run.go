package app

import (
	"bytes"
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	wailsV2 "github.com/wailsapp/wails/v2"
	wailsV2Options "github.com/wailsapp/wails/v2/pkg/options"
	wailsV2Asset "github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var (
	appname = "Wombat"
	semver  = "0.0.0-dev"
)

// Run TODO: viqueen - consider moving all the wails related code to a separate package
func Run(assets embed.FS) int {
	appData, err := appDataLocation(appname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open add data directory: %v\n", err)
		return 1
	}
	defer crashlog(appData)

	err = wailsV2.Run(&wailsV2Options.App{
		Title: appname,
		AssetServer: &wailsV2Asset.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			&api{appData: appData},
		},
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "app: error running app: %v\n", err)
		return 1
	}

	return 0
}

func crashlog(appData string) {
	if r := recover(); r != nil {
		if _, err := os.Stat(appData); os.IsNotExist(err) {
			os.MkdirAll(appData, 0700)
		}
		var b bytes.Buffer
		b.WriteString(fmt.Sprintf("%+v\n\n", r))
		buf := make([]byte, 1<<20)
		s := runtime.Stack(buf, true)
		b.Write(buf[0:s])
		// TODO: viqueen - use non deprecated
		ioutil.WriteFile(filepath.Join(appData, "crash.log"), b.Bytes(), 0644)
	}
}
