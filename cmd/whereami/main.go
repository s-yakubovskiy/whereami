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
// 1. Update documentation and how-to install guide. Also how-to setup IP Quality with api key
// 2. Add option to see history (--history 10 shows last ten locations in nice terminal output)

func main() {
	cmd.Execute()
}
