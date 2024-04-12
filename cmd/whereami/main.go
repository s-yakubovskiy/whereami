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
// 2. Adding experimental flag --gps to add posibility enrich whrmi location output with precise lat & long, altimeter and some other optional stuff

func main() {
	// fmt.Printf("%+v\n", config.Cfg)
	cmd.Execute()
}
