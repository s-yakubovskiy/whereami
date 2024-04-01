package main

import (
	"github.com/s-yakubovskiy/whereami/cmd/cmd"
	"github.com/s-yakubovskiy/whereami/config"
)

var Version string

func init() {
	config.Init()
}

// TODO:
// 1. Test running whereami on clean vps with
// go install github.com/s-yakubovskiy/whereami/cmd/whereami@v1.0.2
// 2. Make default values for config (if it wasn't found). Create it under ~/.config/whereami/
// 3. Update documentation and how-to install guide. Also how-to setup IP Quality with api key

func main() {
	cmd.Execute()
}
