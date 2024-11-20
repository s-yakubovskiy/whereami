package cmd

import (
	"fmt"

	"github.com/mbndr/figlet4go"
)

const appName = "whrmi"

func introduce() {
	ascii := figlet4go.NewAsciiRender()

	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorMagenta,
		figlet4go.ColorCyan,
	}

	renderStr, _ := ascii.RenderOpts(appName, options)
	fmt.Print(renderStr)
	fmt.Println("    ... getting your location data ...")
}
